package tests

import (
	"sort"
	"testing"

	"github.com/bkuri/ppc/internal/model"
	"github.com/bkuri/ppc/internal/resolver"
)

func TestOrderingPrecedence(t *testing.T) {
	t.Run("layer first", func(t *testing.T) {
		mods := []*model.Module{
			{Layer: 2, Front: model.Frontmatter{ID: "b"}},
			{Layer: 0, Front: model.Frontmatter{ID: "a"}},
			{Layer: 1, Front: model.Frontmatter{ID: "c"}},
		}
		sorted := resolver.SortModules(mods)

		if sorted[0].Layer != 0 || sorted[1].Layer != 1 || sorted[2].Layer != 2 {
			t.Errorf("modules not ordered by layer: got %d, %d, %d",
				sorted[0].Layer, sorted[1].Layer, sorted[2].Layer)
		}
	})

	t.Run("priority within layer", func(t *testing.T) {
		mods := []*model.Module{
			{Layer: 0, Front: model.Frontmatter{ID: "b", Priority: 50}},
			{Layer: 0, Front: model.Frontmatter{ID: "a", Priority: 0}},
			{Layer: 0, Front: model.Frontmatter{ID: "c", Priority: 25}},
		}
		sorted := resolver.SortModules(mods)

		if sorted[0].Front.Priority != 0 ||
			sorted[1].Front.Priority != 25 ||
			sorted[2].Front.Priority != 50 {
			t.Errorf("modules not ordered by priority within layer: got %d, %d, %d",
				sorted[0].Front.Priority, sorted[1].Front.Priority, sorted[2].Front.Priority)
		}
	})

	t.Run("id determinism", func(t *testing.T) {
		mods := []*model.Module{
			{Layer: 0, Front: model.Frontmatter{ID: "c", Priority: 10}},
			{Layer: 0, Front: model.Frontmatter{ID: "a", Priority: 10}},
			{Layer: 0, Front: model.Frontmatter{ID: "b", Priority: 10}},
		}
		sorted := resolver.SortModules(mods)

		ids := []string{sorted[0].Front.ID, sorted[1].Front.ID, sorted[2].Front.ID}
		if !sort.StringsAreSorted(ids) {
			t.Errorf("modules not ordered by ID within same layer+priority: got %v", ids)
		}
	})
}

func TestKeyedTagParsing(t *testing.T) {
	tests := []struct {
		tag   string
		group string
		value string
		ok    bool
	}{
		{"risk:low", "risk", "low", true},
		{"tone:terse", "tone", "terse", true},
		{"novalue", "", "", false},
		{":onlycolon", "", "", false},
		{"colonatend:", "", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.tag, func(t *testing.T) {
			g, v, ok := resolver.ParseKeyedTag(tt.tag)
			if ok != tt.ok {
				t.Errorf("ParseKeyedTag(%q) ok = %v, want %v", tt.tag, ok, tt.ok)
			}
			if ok && (g != tt.group || v != tt.value) {
				t.Errorf("ParseKeyedTag(%q) = (%q, %q), want (%q, %q)",
					tt.tag, g, v, tt.group, tt.value)
			}
		})
	}
}
