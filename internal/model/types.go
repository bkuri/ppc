// Package model defines core data structures
package model

import (
	"path/filepath"
	"strings"
)

// Rules defines validation rules for modules
type Rules struct {
	ExclusiveGroups []string `yaml:"exclusive_groups"`
}

// Frontmatter represents the YAML frontmatter of a module
type Frontmatter struct {
	ID       string   `yaml:"id"`
	Desc     string   `yaml:"desc"`
	Priority int      `yaml:"priority"`
	Tags     []string `yaml:"tags"`
	Requires []string `yaml:"requires"`
}

// Module represents a compiled module with metadata
type Module struct {
	Path     string
	Layer    int
	Front    Frontmatter
	Body     string
	FromReq  bool
	Selected bool
}

// LayerOrder defines the canonical layer precedence
var LayerOrder = []string{"base", "modes", "traits", "policies", "contracts"}

// LayerIndexFromPath returns the layer index for a module path
func LayerIndexFromPath(p string) int {
	parts := strings.Split(filepath.ToSlash(p), "/")
	for i, s := range LayerOrder {
		for _, part := range parts {
			if part == s {
				return i
			}
		}
	}
	return 0
}
