// Package doctor provides module validation
package doctor

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bkuri/ppc/internal/loader"
	"github.com/bkuri/ppc/internal/resolver"
)

// RunDoctor validates module structure and dependencies
// Returns exit code: 0=ok, 2=failed
func RunDoctor(promptsDir string, strict bool, jsonOut bool) int {
	modByID, err := loader.LoadModules(promptsDir)
	if err != nil {
		fmt.Println("doctor: FAILED")
		fmt.Println("errors:")
		fmt.Printf("  - %v\n", err)
		return 2
	}

	rules, err := loader.LoadRules(promptsDir)
	if err != nil {
		fmt.Println("doctor: FAILED")
		fmt.Println("errors:")
		fmt.Printf("  - %v\n", err)
		return 2
	}

	var errs []string
	var warns []string

	// Validate tag format
	groupVals := map[string]map[string]bool{}
	for _, m := range modByID {
		for _, t := range m.Front.Tags {
			g, v, ok := resolver.ParseKeyedTag(t)
			if !ok {
				errs = append(errs, fmt.Sprintf("module %s has invalid tag %q (expected group:value)", m.Front.ID, t))
				continue
			}
			if groupVals[g] == nil {
				groupVals[g] = map[string]bool{}
			}
			groupVals[g][v] = true
		}
	}

	// Validate requires targets exist
	for _, m := range modByID {
		for _, r := range m.Front.Requires {
			if _, ok := modByID[r]; !ok {
				errs = append(errs, fmt.Sprintf("requires target not found: %s (referenced by %s)", r, m.Front.ID))
			}
		}
	}

	// Check for circular dependencies
	const (
		unvisited = 0
		visiting  = 1
		done      = 2
	)
	state := map[string]int{}
	stack := []string{}
	pos := map[string]int{}

	var visit func(id string) error
	visit = func(id string) error {
		switch state[id] {
		case done:
			return nil
		case visiting:
			i := pos[id]
			cycle := append(append([]string{}, stack[i:]...), id)
			return fmt.Errorf("circular requires: %s", strings.Join(cycle, " -> "))
		}
		state[id] = visiting
		pos[id] = len(stack)
		stack = append(stack, id)

		reqs := append([]string{}, modByID[id].Front.Requires...)
		sort.Strings(reqs)
		for _, r := range reqs {
			if _, ok := modByID[r]; !ok {
				continue
			}
			if err := visit(r); err != nil {
				return err
			}
		}

		stack = stack[:len(stack)-1]
		delete(pos, id)
		state[id] = done
		return nil
	}

	ids := make([]string, 0, len(modByID))
	for id := range modByID {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		if err := visit(id); err != nil {
			errs = append(errs, err.Error())
			break
		}
	}

	// Validate exclusive groups
	if len(rules.ExclusiveGroups) == 0 {
		warns = append(warns, "rules.yml: exclusive_groups is empty")
	}
	for _, g := range rules.ExclusiveGroups {
		if groupVals[g] == nil {
			warns = append(warns, fmt.Sprintf("exclusive group %q never appears in any module tags", g))
		}
	}

	// Check for unreachable modules
	entry := map[string]bool{"base": true}
	for id := range modByID {
		if strings.HasPrefix(id, "modes/") || strings.HasPrefix(id, "contracts/") {
			entry[id] = true
		}
	}

	reachable := map[string]bool{}
	var mark func(id string)
	mark = func(id string) {
		if reachable[id] {
			return
		}
		reachable[id] = true
		reqs := append([]string{}, modByID[id].Front.Requires...)
		sort.Strings(reqs)
		for _, r := range reqs {
			if _, ok := modByID[r]; ok {
				mark(r)
			}
		}
	}

	entryIDs := make([]string, 0, len(entry))
	for id := range entry {
		entryIDs = append(entryIDs, id)
	}
	sort.Strings(entryIDs)
	for _, id := range entryIDs {
		if _, ok := modByID[id]; ok {
			mark(id)
		} else if id == "base" {
			errs = append(errs, "missing required entrypoint module: base")
		}
	}

	var dead []string
	for id := range modByID {
		if !reachable[id] {
			dead = append(dead, id)
		}
	}
	sort.Strings(dead)
	if len(dead) > 0 {
		warns = append(warns, fmt.Sprintf("unreachable modules (%d): %s", len(dead), strings.Join(dead, ", ")))
	}

	// Output results
	if jsonOut {
		return printDoctorJSON(len(modByID), errs, warns, strict)
	}

	if len(errs) == 0 {
		fmt.Printf("doctor: OK (%d modules)\n", len(modByID))
		if len(warns) > 0 {
			fmt.Println("warnings:")
			for _, w := range warns {
				fmt.Println("  - " + w)
			}
			if strict {
				fmt.Println("doctor: FAILED (strict mode; warnings treated as errors)")
				return 2
			}
		}
		return 0
	}

	fmt.Println("doctor: FAILED")
	fmt.Println("errors:")
	for _, e := range errs {
		fmt.Println("  - " + e)
	}
	if len(warns) > 0 {
		fmt.Println("warnings:")
		for _, w := range warns {
			fmt.Println("  - " + w)
		}
	}
	return 2
}
