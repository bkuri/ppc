package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
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

func runExplore(args []string, promptsDir string) {
	fs := flag.NewFlagSet("explore", flag.ExitOnError)

	conservative := fs.Bool("conservative", false, "include traits/conservative")
	creative := fs.Bool("creative", false, "include traits/creative")
	terse := fs.Bool("terse", false, "include traits/terse")
	verbose := fs.Bool("verbose", false, "include traits/verbose")
	revisions := fs.Int("revisions", -1, "revision budget (enables policies/revisions)")
	contract := fs.String("contract", "markdown", "contract module (code|markdown)")
	outPath := fs.String("out", "", "write output to file")
	explain := fs.Bool("explain", false, "explain resolution steps to stderr")
	withHash := fs.Bool("hash", false, "prepend prompt-id hash header")
	proDir := fs.String("prompts", promptsDir, "prompts directory")

	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, `usage:
  ppc explore [flags]

Explore mode generates a prompt for exploration tasks.

flags:`)
		fs.PrintDefaults()
	}

	fs.Parse(args)

	opts := CompileOptions{
		Conservative: *conservative,
		Creative:     *creative,
		Terse:        *terse,
		Verbose:      *verbose,
		Revisions:    *revisions,
		Contract:     *contract,
		PromptsDir:   *proDir,
		Explain:      *explain,
		WithHash:     *withHash,
	}

	result, err := compilePrompt("explore", opts)
	if err != nil {
		dief("%v", err)
	}

	if *outPath != "" {
		if err := os.WriteFile(*outPath, []byte(result.Output), 0o644); err != nil {
			dief("failed to write %s: %v", *outPath, err)
		}
	}

	fmt.Print(result.Output)
}

func runBuild(args []string, promptsDir string) {
	fs := flag.NewFlagSet("build", flag.ExitOnError)

	conservative := fs.Bool("conservative", false, "include traits/conservative")
	creative := fs.Bool("creative", false, "include traits/creative")
	terse := fs.Bool("terse", false, "include traits/terse")
	verbose := fs.Bool("verbose", false, "include traits/verbose")
	revisions := fs.Int("revisions", -1, "revision budget (enables policies/revisions)")
	contract := fs.String("contract", "markdown", "contract module (code|markdown)")
	outPath := fs.String("out", "", "write output to file")
	explain := fs.Bool("explain", false, "explain resolution steps to stderr")
	withHash := fs.Bool("hash", false, "prepend prompt-id hash header")
	proDir := fs.String("prompts", promptsDir, "prompts directory")

	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, `usage:
  ppc build [flags]

Build mode generates a prompt for building/implementing features.

flags:`)
		fs.PrintDefaults()
	}

	fs.Parse(args)

	opts := CompileOptions{
		Conservative: *conservative,
		Creative:     *creative,
		Terse:        *terse,
		Verbose:      *verbose,
		Revisions:    *revisions,
		Contract:     *contract,
		PromptsDir:   *proDir,
		Explain:      *explain,
		WithHash:     *withHash,
	}

	result, err := compilePrompt("build", opts)
	if err != nil {
		dief("%v", err)
	}

	if *outPath != "" {
		if err := os.WriteFile(*outPath, []byte(result.Output), 0o644); err != nil {
			dief("failed to write %s: %v", *outPath, err)
		}
	}

	fmt.Print(result.Output)
}

func runShip(args []string, promptsDir string) {
	fs := flag.NewFlagSet("ship", flag.ExitOnError)

	conservative := fs.Bool("conservative", false, "include traits/conservative")
	creative := fs.Bool("creative", false, "include traits/creative")
	terse := fs.Bool("terse", false, "include traits/terse")
	verbose := fs.Bool("verbose", false, "include traits/verbose")
	revisions := fs.Int("revisions", -1, "revision budget (enables policies/revisions)")
	contract := fs.String("contract", "markdown", "contract module (code|markdown)")
	outPath := fs.String("out", "", "write output to file")
	explain := fs.Bool("explain", false, "explain resolution steps to stderr")
	withHash := fs.Bool("hash", false, "prepend prompt-id hash header")
	proDir := fs.String("prompts", promptsDir, "prompts directory")

	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, `usage:
  ppc ship [flags]

Ship mode generates a prompt for release/deployment tasks.

flags:`)
		fs.PrintDefaults()
	}

	fs.Parse(args)

	opts := CompileOptions{
		Conservative: *conservative,
		Creative:     *creative,
		Terse:        *terse,
		Verbose:      *verbose,
		Revisions:    *revisions,
		Contract:     *contract,
		PromptsDir:   *proDir,
		Explain:      *explain,
		WithHash:     *withHash,
	}

	result, err := compilePrompt("ship", opts)
	if err != nil {
		dief("%v", err)
	}

	if *outPath != "" {
		if err := os.WriteFile(*outPath, []byte(result.Output), 0o644); err != nil {
			dief("failed to write %s: %v", *outPath, err)
		}
	}

	fmt.Print(result.Output)
}

func printGlobalUsage() {
	fmt.Fprintln(os.Stderr, `usage:
  ppc <subcommand> [flags]

subcommands:
  explore    Generate prompt for exploration mode
  build      Generate prompt for build mode
  ship       Generate prompt for shipping mode
  doctor     Validate module structure and dependencies

global flags:
  --list     List all available modules
  --help     Show this help message

examples:
  ppc explore --conservative --revisions 1 --contract markdown
  ppc build --conservative --revisions 1 --contract code --explain
  ppc ship --creative --out AGENTS.md --hash
  ppc doctor --strict --json

run 'ppc <subcommand> --help' for subcommand-specific options`)
}

func handleListModules(promptsDir string) {
	modByID := loadModules(promptsDir)
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
}

func main() {
	if len(os.Args) < 2 {
		printGlobalUsage()
		os.Exit(1)
	}

	subcommand := os.Args[1]
	args := os.Args[2:]

	// Default prompts directory
	promptsDir := "prompts"

	// Handle global meta-flags first
	if subcommand == "--list" {
		handleListModules(promptsDir)
		os.Exit(0)
	}

	if subcommand == "--help" || subcommand == "-h" || subcommand == "help" {
		printGlobalUsage()
		os.Exit(0)
	}

	// Dispatch to subcommand
	switch subcommand {
	case "explore":
		runExplore(args, promptsDir)
	case "build":
		runBuild(args, promptsDir)
	case "ship":
		runShip(args, promptsDir)
	case "doctor":
		// Doctor has special flag parsing
		fs := flag.NewFlagSet("doctor", flag.ExitOnError)
		strict := fs.Bool("strict", false, "treat warnings as errors")
		jsonOut := fs.Bool("json", false, "output machine-readable JSON")
		proDir := fs.String("prompts", promptsDir, "prompts directory")
		fs.Usage = func() {
			fmt.Fprintln(os.Stderr, `usage:
  ppc doctor [flags]

Checks module integrity, requires, cycles, and tag/rules sanity.

flags:`)
			fs.PrintDefaults()
		}
		fs.Parse(args)
		os.Exit(runDoctor(*proDir, *strict, *jsonOut))

	default:
		fmt.Fprintf(os.Stderr, "unknown subcommand: %s\n", subcommand)
		printGlobalUsage()
		os.Exit(1)
	}
}
