package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type CompileOptions struct {
	Conservative bool
	Creative     bool
	Terse        bool
	Verbose      bool
	Revisions    int
	Contract     string
	PromptsDir   string
	Explain      bool
	WithHash     bool
}

type CompileResult struct {
	Output string
}

func compilePrompt(mode string, opts CompileOptions) (*CompileResult, error) {
	// Load modules and rules
	modByID := loadModules(opts.PromptsDir)
	rules := loadRules(opts.PromptsDir)

	// Build selected IDs
	selectedIDs := []string{"base", "modes/" + mode, "contracts/" + opts.Contract}
	if opts.Conservative {
		selectedIDs = append(selectedIDs, "traits/conservative")
	}
	if opts.Creative {
		selectedIDs = append(selectedIDs, "traits/creative")
	}
	if opts.Terse {
		selectedIDs = append(selectedIDs, "traits/terse")
	}
	if opts.Verbose {
		selectedIDs = append(selectedIDs, "traits/verbose")
	}
	if opts.Revisions >= 0 {
		selectedIDs = append(selectedIDs, "policies/revisions")
	}

	// Expand requires
	closureIDs, fromReq, err := expandRequires(selectedIDs, modByID)
	if err != nil {
		return nil, err
	}

	// Build module list
	var mods []*Module
	for _, id := range closureIDs {
		m := modByID[id]
		m.FromReq = fromReq[id]
		m.Selected = inSlice(selectedIDs, id)
		mods = append(mods, m)
	}

	// Validate exclusive groups
	if err := validateExclusiveGroups(rules, mods); err != nil {
		return nil, err
	}

	// Sort modules
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

	// Prepare variables for rendering
	vars := map[string]string{"mode": mode}
	if opts.Revisions >= 0 {
		vars["revisions"] = strconv.Itoa(opts.Revisions)
	}

	// Render output
	out := render(mods, vars)
	if opts.WithHash {
		h := sha256Hex(out)
		out = fmt.Sprintf("<!-- prompt-id: sha256:%s -->\n\n%s", h, out)
	}

	// Handle explain output (if needed by caller)
	if opts.Explain {
		explainOutput(mods, selectedIDs, closureIDs, fromReq)
	}

	return &CompileResult{Output: out}, nil
}

// Helper: Print explain output to stderr
func explainOutput(mods []*Module, selectedIDs []string, closureIDs []string, fromReq map[string]bool) {
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
