package profile

import (
	"fmt"
	"github.com/bkuri/ppc/internal/compile"
	"github.com/bkuri/ppc/internal/loader"
	"github.com/bkuri/ppc/internal/model"
	"github.com/bkuri/ppc/internal/resolver"
)

// MergeOptions represents CLI flags that can override a profile
type MergeOptions struct {
	Conservative *bool
	Creative     *bool
	Terse        *bool
	Verbose      *bool
	Revisions    *int
	Contract     *string
}

// Merge merges a profile with CLI flag overrides
// Semantics:
// 1. Start with profile values
// 2. Override with explicit CLI flags (if set)
// 3. Validate exclusive group consistency (conflicts = error)
//
// Returns: merged CompileOptions, or error if exclusive groups conflict
func Merge(prof *Profile, opts MergeOptions) (*compile.CompileOptions, error) {
	traits := append([]string{}, prof.Traits...)

	if opts.Conservative != nil && *opts.Conservative {
		traits = append(traits, "traits/conservative")
	}
	if opts.Creative != nil && *opts.Creative {
		traits = append(traits, "traits/creative")
	}
	if opts.Terse != nil && *opts.Terse {
		traits = append(traits, "traits/terse")
	}
	if opts.Verbose != nil && *opts.Verbose {
		traits = append(traits, "traits/verbose")
	}

	rules, _ := loader.LoadRules("prompts")
	if err := validateExclusiveGroups(traits, rules); err != nil {
		return nil, err
	}

	contract := prof.Contract
	if opts.Contract != nil && *opts.Contract != "" {
		contract = *opts.Contract
	}

	vars := make(map[string]string)
	if prof.Vars != nil {
		for k, v := range prof.Vars {
			vars[k] = v
		}
	}

	return &compile.CompileOptions{
		Mode:       prof.Mode,
		Contract:   contract,
		Traits:     traits,
		PromptsDir: "prompts",
		Vars:       vars,
	}, nil
}

func varsFromProfile(prof *Profile) map[string]string {
	if prof.Vars == nil {
		return make(map[string]string)
	}
	return prof.Vars
}

// validateExclusiveGroups checks if final trait set violates exclusive group rules
// Returns error with specific conflict details
func validateExclusiveGroups(traits []string, rules *model.Rules) error {
	modByID, _ := loader.LoadModules("prompts")

	tagsByGroup := make(map[string][]string)
	traitsByGroup := make(map[string][]string)

	for _, t := range traits {
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
