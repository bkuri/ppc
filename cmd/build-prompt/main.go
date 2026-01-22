package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Rules struct {
	ExclusiveGroups []string `yaml:"exclusive_groups"`
}

type Frontmatter struct {
	ID       string   `yaml:"id"`
	Desc     string   `yaml:"desc"`
	Priority int      `yaml:"priority"`
	Tags     []string `yaml:"tags"`
	Requires []string `yaml:"requires"`
}

type Module struct {
	Path     string
	Layer    int
	Front    Frontmatter
	Body     string
	FromReq  bool
	Selected bool
}

var layerOrder = []string{"base", "modes", "traits", "policies", "contracts"}

func layerIndexFromPath(p string) int {
	parts := strings.Split(filepath.ToSlash(p), "/")
	for i, s := range layerOrder {
		for _, part := range parts {
			if part == s {
				return i
			}
		}
	}
	return 0
}

func dief(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(2)
}

func parseFrontmatter(raw []byte) (Frontmatter, string, bool, error) {
	s := string(raw)
	if !strings.HasPrefix(s, "---\n") && !strings.HasPrefix(s, "---\r\n") {
		return Frontmatter{}, strings.TrimRight(s, "\n"), false, nil
	}

	idx := strings.Index(s[4:], "\n---\n")
	delimLen := len("\n---\n")
	if idx == -1 {
		idx = strings.Index(s[4:], "\r\n---\r\n")
		delimLen = len("\r\n---\r\n")
	}
	if idx == -1 {
		return Frontmatter{}, "", false, fmt.Errorf("frontmatter start found but missing closing ---")
	}

	yml := s[4 : 4+idx]
	body := s[4+idx+delimLen:]
	body = strings.TrimLeft(body, "\r\n")
	body = strings.TrimRight(body, "\n")

	var fm Frontmatter
	if err := yaml.Unmarshal([]byte(yml), &fm); err != nil {
		return Frontmatter{}, "", false, fmt.Errorf("invalid YAML frontmatter: %w", err)
	}
	return fm, body, true, nil
}

func loadRules(promptsDir string) Rules {
	p := filepath.Join(promptsDir, "rules.yml")
	b, err := os.ReadFile(p)
	if err != nil {
		dief("missing rules file: %s", p)
	}
	var r Rules
	if err := yaml.Unmarshal(b, &r); err != nil {
		dief("invalid rules.yml: %v", err)
	}
	return r
}

func listMarkdownFiles(root string) []string {
	var out []string
	_ = filepath.WalkDir(root, func(p string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
			out = append(out, p)
		}
		return nil
	})
	sort.Strings(out)
	return out
}

func loadModules(promptsDir string) map[string]*Module {
	files := listMarkdownFiles(promptsDir)
	modByID := map[string]*Module{}

	for _, p := range files {
		raw, err := os.ReadFile(p)
		if err != nil {
			dief("failed to read %s: %v", p, err)
		}
		fm, body, has, err := parseFrontmatter(raw)
		if err != nil {
			dief("%s: %v", p, err)
		}
		if !has {
			dief("%s: missing frontmatter (v0.1 requires YAML frontmatter with id)", p)
		}
		if strings.TrimSpace(fm.ID) == "" {
			dief("%s: frontmatter missing required field: id", p)
		}
		m := &Module{
			Path:  p,
			Layer: layerIndexFromPath(filepath.ToSlash(p)),
			Front: fm,
			Body:  body,
		}
		if _, exists := modByID[fm.ID]; exists {
			dief("duplicate module id %q (found at %s)", fm.ID, p)
		}
		modByID[fm.ID] = m
	}

	return modByID
}

func parseKeyedTag(t string) (group, value string, ok bool) {
	i := strings.IndexByte(t, ':')
	if i <= 0 || i >= len(t)-1 {
		return "", "", false
	}
	return t[:i], t[i+1:], true
}

func validateExclusiveGroups(r Rules, mods []*Module) error {
	groupValues := map[string]map[string]bool{}
	for _, m := range mods {
		for _, t := range m.Front.Tags {
			g, v, ok := parseKeyedTag(t)
			if !ok {
				return fmt.Errorf("module %s has invalid tag %q (expected group:value)", m.Front.ID, t)
			}
			if groupValues[g] == nil {
				groupValues[g] = map[string]bool{}
			}
			groupValues[g][v] = true
		}
	}

	excl := map[string]bool{}
	for _, g := range r.ExclusiveGroups {
		excl[g] = true
	}

	for g, vals := range groupValues {
		if !excl[g] || len(vals) <= 1 {
			continue
		}
		var vs []string
		for v := range vals {
			vs = append(vs, v)
		}
		sort.Strings(vs)
		return fmt.Errorf("conflicting tags in group %q: %s", g, strings.Join(vs, ", "))
	}

	return nil
}

func inSlice(xs []string, s string) bool {
	for _, x := range xs {
		if x == s {
			return true
		}
	}
	return false
}

func expandRequires(selectedIDs []string, all map[string]*Module) ([]string, map[string]bool, error) {
	const (
		unvisited = 0
		visiting  = 1
		done      = 2
	)

	state := map[string]int{}
	stack := []string{}
	pos := map[string]int{}
	out := []string{}
	inOut := map[string]bool{}
	fromReq := map[string]bool{}

	var dfs func(id string, rootSelected bool) error
	dfs = func(id string, rootSelected bool) error {
		m, ok := all[id]
		if !ok {
			return fmt.Errorf("required module not found: %s", id)
		}

		switch state[id] {
		case done:
			if rootSelected {
				m.Selected = true
				fromReq[id] = false
			}
			return nil
		case visiting:
			i := pos[id]
			cycle := append(append([]string{}, stack[i:]...), id)
			return fmt.Errorf("circular requires: %s", strings.Join(cycle, " -> "))
		}

		state[id] = visiting
		pos[id] = len(stack)
		stack = append(stack, id)

		reqs := append([]string{}, m.Front.Requires...)
		sort.Strings(reqs)
		for _, r := range reqs {
			if err := dfs(r, false); err != nil {
				return err
			}
			if !inSlice(selectedIDs, r) {
				fromReq[r] = true
			}
		}

		stack = stack[:len(stack)-1]
		delete(pos, id)
		state[id] = done

		if !inOut[id] {
			inOut[id] = true
			out = append(out, id)
		}
		if rootSelected {
			m.Selected = true
			fromReq[id] = false
		}
		return nil
	}

	ids := append([]string{}, selectedIDs...)
	sort.Strings(ids)
	for _, id := range ids {
		if err := dfs(id, true); err != nil {
			return nil, nil, err
		}
	}

	return out, fromReq, nil
}

func render(mods []*Module, vars map[string]string) string {
	var b strings.Builder
	for i, m := range mods {
		if i > 0 {
			b.WriteString("\n\n")
		}
		body := m.Body
		for k, v := range vars {
			body = strings.ReplaceAll(body, "{{"+k+"}}", v)
		}
		b.WriteString(strings.TrimRight(body, "\n"))
	}
	return strings.TrimRight(b.String(), "\n") + "\n"
}

func sha256Hex(s string) string {
	h := sha256.Sum256([]byte(s))
	return hex.EncodeToString(h[:])
}

func main() {
	var (
		conservative = flag.Bool("conservative", false, "include traits/conservative")
		creative     = flag.Bool("creative", false, "include traits/creative")
		terse        = flag.Bool("terse", false, "include traits/terse")
		verbose      = flag.Bool("verbose", false, "include traits/verbose")
		revisions    = flag.Int("revisions", -1, "revision budget (enables policies/revisions)")
		contract     = flag.String("contract", "markdown", "contract module (code|markdown)")
		promptsDir   = flag.String("prompts", "prompts", "prompts directory")
		outPath      = flag.String("out", "", "write output to file")
		list         = flag.Bool("list", false, "list available modules")
		explain      = flag.Bool("explain", false, "explain resolution steps to stderr")
		withHash     = flag.Bool("hash", false, "prepend prompt-id hash header")
	)

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `usage:
  ppc [flags] <mode>

modes:
  explore | build | ship  (loads prompts/modes/<mode>.md)

examples:
  ppc --conservative --revisions 1 --contract markdown explore
  ppc --conservative --revisions 1 --contract code --explain ship
  ppc --creative --out AGENTS.md explore

flags:`)
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	modByID := loadModules(*promptsDir)

	if *list {
		var ids []string
		for id := range modByID {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		for _, id := range ids {
			m := modByID[id]
			desc := strings.TrimSpace(m.Front.Desc)
			if desc == "" {
				desc = "(no desc)"
			}
			fmt.Printf("%-22s  %s\n", id, desc)
		}
		return
	}

	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	mode := args[0]

	selectedIDs := []string{"base", "modes/" + mode, "contracts/" + *contract}
	if *conservative {
		selectedIDs = append(selectedIDs, "traits/conservative")
	}
	if *creative {
		selectedIDs = append(selectedIDs, "traits/creative")
	}
	if *terse {
		selectedIDs = append(selectedIDs, "traits/terse")
	}
	if *verbose {
		selectedIDs = append(selectedIDs, "traits/verbose")
	}
	if *revisions >= 0 {
		selectedIDs = append(selectedIDs, "policies/revisions")
	}

	rules := loadRules(*promptsDir)

	closureIDs, fromReq, err := expandRequires(selectedIDs, modByID)
	if err != nil {
		dief("%v", err)
	}

	var mods []*Module
	for _, id := range closureIDs {
		m := modByID[id]
		m.FromReq = fromReq[id]
		m.Selected = inSlice(selectedIDs, id)
		mods = append(mods, m)
	}

	if err := validateExclusiveGroups(rules, mods); err != nil {
		dief("%v", err)
	}

	sort.Slice(mods, func(i, j int) bool {
		a, b := mods[i], mods[j]
		if a.Layer != b.Layer {
			return a.Layer < b.Layer
		}
		if a.Front.Priority != b.Front.Priority {
			return a.Front.Priority < b.Front.Priority
		}
		return a.Front.ID < b.Front.ID
	})

	vars := map[string]string{"mode": mode}
	if *revisions >= 0 {
		vars["revisions"] = strconv.Itoa(*revisions)
	}

	out := render(mods, vars)
	if *withHash {
		h := sha256Hex(out)
		out = fmt.Sprintf("<!-- prompt-id: sha256:%s -->\n\n%s", h, out)
	}

	if *explain {
		var buf bytes.Buffer
		buf.WriteString("PPC explain\n")

		buf.WriteString("Selected IDs:\n")
		sel := append([]string{}, selectedIDs...)
		sort.Strings(sel)
		for _, id := range sel {
			buf.WriteString("  - " + id + "\n")
		}

		buf.WriteString("Closure IDs (after requires):\n")
		cls := append([]string{}, closureIDs...)
		sort.Strings(cls)
		for _, id := range cls {
			note := ""
			if fromReq[id] && !inSlice(selectedIDs, id) {
				note = " (required)"
			}
			buf.WriteString("  - " + id + note + "\n")
		}

		buf.WriteString("Final order:\n")
		for _, m := range mods {
			note := ""
			if m.FromReq && !m.Selected {
				note = " (required)"
			}
			buf.WriteString(fmt.Sprintf("  - [%d] %s prio=%d%s\n", m.Layer, m.Front.ID, m.Front.Priority, note))
		}

		io.Copy(os.Stderr, &buf)
	}

	if *outPath != "" {
		if err := os.WriteFile(*outPath, []byte(out), 0o644); err != nil {
			dief("failed to write %s: %v", *outPath, err)
		}
	}

	fmt.Print(out)
}
