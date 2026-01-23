package main

import (
	"fmt"

	compilepkg "github.com/bkuri/ppc/internal/compile"
	"github.com/bkuri/ppc/internal/loader"
	profilepkg "github.com/bkuri/ppc/internal/profile"
	"github.com/bkuri/ppc/internal/resolver"
)

type ResolvedConfig struct {
	Mode       string
	Contract   string
	Revisions  int
	Traits     []string
	Vars       map[string]string
	PromptsDir string
}

func NewResolvedConfigFromProfile(profileName string) (*ResolvedConfig, error) {
	profile, err := profilepkg.LoadProfile(profileName)
	if err != nil {
		return nil, err
	}

	cfg := ResolvedConfig{
		Mode:       profile.Mode,
		Contract:   profile.Contract,
		Revisions:  -1,
		Traits:     profile.Traits,
		Vars:       map[string]string{},
		PromptsDir: "",
	}

	if profile.Revisions != nil {
		cfg.Revisions = *profile.Revisions
	}

	if profile.Vars != nil {
		cfg.Vars = profile.Vars
	}

	return &cfg, nil
}

func NewResolvedConfigFromDefaults(mode string, contract string) ResolvedConfig {
	return ResolvedConfig{
		Mode:       mode,
		Contract:   contract,
		Revisions:  -1,
		Traits:     []string{},
		Vars:       map[string]string{},
		PromptsDir: "",
	}
}

func (c *ResolvedConfig) ApplyCLIOverrides(conservative, creative, terse, verbose *bool, revisions *int, contract *string) (*ResolvedConfig, error) {
	cfg := *c

	if conservative != nil && *conservative {
		cfg.Traits = append(cfg.Traits, "traits/conservative")
	}
	if creative != nil && *creative {
		cfg.Traits = append(cfg.Traits, "traits/creative")
	}
	if terse != nil && *terse {
		cfg.Traits = append(cfg.Traits, "traits/terse")
	}
	if verbose != nil && *verbose {
		cfg.Traits = append(cfg.Traits, "traits/verbose")
	}
	if revisions != nil {
		cfg.Revisions = *revisions
	}
	if contract != nil && *contract != "" {
		cfg.Contract = *contract
	}

	return &cfg, validateExclusiveGroups(cfg.Traits)
}

func (c *ResolvedConfig) ToCompileOptions() compilepkg.CompileOptions {
	return compilepkg.CompileOptions{
		Mode:       c.Mode,
		Contract:   c.Contract,
		Traits:     c.Traits,
		PromptsDir: c.PromptsDir,
		Vars:       c.Vars,
	}
}

func validateExclusiveGroups(traits []string) error {
	modByID, _ := loader.LoadModules("prompts")

	for _, t := range traits {
		if _, exists := modByID[t]; !exists {
			continue
		}

		for _, tag := range modByID[t].Front.Tags {
			group, val, ok := resolver.ParseKeyedTag(tag)
			if !ok {
				continue
			}

			for _, other := range traits {
				if other == t {
					continue
				}

				otherModule, exists := modByID[other]
				if !exists {
					continue
				}

				for _, otherTag := range otherModule.Front.Tags {
					otherGroup, otherVal, ok2 := resolver.ParseKeyedTag(otherTag)
					if !ok2 {
						continue
					}

					if otherGroup == group && otherVal != val {
						return fmt.Errorf("exclusive group %q violated:\n  %s sets %s:%s\n  %s sets %s:%s",
							group,
							t, group, val,
							other, group, otherVal)
					}
				}
			}
		}
	}

	return nil
}
