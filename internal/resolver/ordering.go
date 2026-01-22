package resolver

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bkuri/ppc/internal/model"
)

// ValidateExclusiveGroups checks that no exclusive group has conflicting values
func ValidateExclusiveGroups(r *model.Rules, mods []*model.Module) error {
	groupValues := map[string]map[string]bool{}
	for _, m := range mods {
		for _, t := range m.Front.Tags {
			g, v, ok := ParseKeyedTag(t)
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

// SortModules orders modules by: (1) layer, (2) priority, (3) id
func SortModules(mods []*model.Module) []*model.Module {
	sorted := append([]*model.Module{}, mods...)
	sort.Slice(sorted, func(i, j int) bool {
		a, b := sorted[i], sorted[j]
		if a.Layer != b.Layer {
			return a.Layer < b.Layer
		}
		if a.Front.Priority != b.Front.Priority {
			return a.Front.Priority < b.Front.Priority
		}
		return a.Front.ID < b.Front.ID
	})
	return sorted
}
