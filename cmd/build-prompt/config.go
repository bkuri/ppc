package main

import (
	"fmt"
	"strconv"

	"github.com/bkuri/ppc/internal/compile"
	"github.com/bkuri/ppc/internal/loader"
	"github.com/bkuri/ppc/internal/profile"
	"github.com/bkuri/ppc/internal/resolver"
)

// ResolvedConfig represents the final configuration after resolving
// profile defaults + CLI flag overrides.
// It's the intermediate format before converting to CompileOptions.
type ResolvedConfig struct {
	Mode       string            // "explore", "build", "ship"
	Contract   string            // "markdown", "code"
	Revisions  int               // -1 means not set
	Traits     []string          // e.g., ["traits/conservative"]
	Vars       map[string]string // custom variables
	PromptsDir string
}

// NewResolvedConfigFromMode creates a default ResolvedConfig for a given mode
func NewResolvedConfigFromMode(mode string) ResolvedConfig {
	return ResolvedConfig{
		Mode:       mode,
		Contract:   "markdown",
		Revisions:  -1,
		Traits:     []string{},
		Vars:       make(map[string]string),
		PromptsDir: "prompts",
	}
}

// NewResolvedConfigFromProfile loads a profile and returns a ResolvedConfig
func NewResolvedConfigFromProfile(profileName, mode string) (*ResolvedConfig, error) {
	prof, err := profile.LoadProfile(profileName)
	if err != nil {
		return nil, fmt.Errorf("load profile: %w", err)
	}

	cfg := ResolvedConfig{
		Mode:       prof.Mode, // Use profile's mode, not the requested mode
		Contract:   prof.Contract,
		Revisions:  -1,
		Traits:     append([]string{}, prof.Traits...),
		Vars:       make(map[string]string),
		PromptsDir: "prompts",
	}

	// Copy vars from profile
	if prof.Vars != nil {
		for k, v := range prof.Vars {
			cfg.Vars[k] = v
		}
	}

	// Copy revisions if set in profile
	if prof.Revisions != nil {
		cfg.Revisions = *prof.Revisions
	}

	return &cfg, nil
}

// ApplyCLIOverrides applies CLI flag overrides to the config
// Returns error if CLI flags conflict with exclusive groups
func (c *ResolvedConfig) ApplyCLIOverrides(conservative, creative, terse, verbose bool, revisions int, contract string, promptsDir string) (*ResolvedConfig, error) {
	// Make a copy to avoid modifying receiver
	cfg := *c

	// Apply mode-independent CLI overrides
	if contract != "" {
		cfg.Contract = contract
	}
	if promptsDir != "" {
		cfg.PromptsDir = promptsDir
	}
	if revisions >= 0 {
		cfg.Revisions = revisions
	}

	// Apply trait flags (only if explicitly set)
	if conservative {
		cfg.Traits = append(cfg.Traits, "traits/conservative")
	}
	if creative {
		cfg.Traits = append(cfg.Traits, "traits/creative")
	}
	if terse {
		cfg.Traits = append(cfg.Traits, "traits/terse")
	}
	if verbose {
		cfg.Traits = append(cfg.Traits, "traits/verbose")
	}

	// Validate exclusive groups after applying all traits
	if err := cfg.validateExclusiveGroups(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// validateExclusiveGroups checks if the trait set violates exclusive group rules
func (c *ResolvedConfig) validateExclusiveGroups() error {
	modByID, _ := loader.LoadModules(c.PromptsDir)

	tagsByGroup := make(map[string][]string)
	traitsByGroup := make(map[string][]string)

	for _, t := range c.Traits {
		if _, exists := modByID[t]; !exists {
			continue
		}

		for _, tag := range modByID[t].Front.Tags {
			group, val, ok := resolver.ParseKeyedTag(tag)
			if !ok {
				continue
			}

			for _, existing := range tagsByGroup[group] {
				if existing != val {
					return fmt.Errorf("exclusive group %q violated:\n  %s sets %s:%s\n  %s sets %s:%s",
						group,
						traitsByGroup[group][0], group, existing,
						t, group, val)
				}
			}

			tagsByGroup[group] = append(tagsByGroup[group], val)
			traitsByGroup[group] = append(traitsByGroup[group], t)
		}
	}

	return nil
}

// ToCompileOptions converts the ResolvedConfig to a CompileOptions struct
// Ready to pass to compile.Compile()
func (c *ResolvedConfig) ToCompileOptions() *compile.CompileOptions {
	vars := make(map[string]string)

	// Copy any existing vars
	for k, v := range c.Vars {
		vars[k] = v
	}

	// Set mode variable
	vars["mode"] = c.Mode

	// Set revisions variable if applicable
	if c.Revisions >= 0 {
		vars["revisions"] = strconv.Itoa(c.Revisions)
	}

	return &compile.CompileOptions{
		Mode:       c.Mode,
		Contract:   c.Contract,
		Traits:     c.Traits,
		PromptsDir: c.PromptsDir,
		Vars:       vars,
	}
}
