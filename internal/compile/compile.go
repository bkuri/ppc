package compile

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/bkuri/ppc/internal/loader"
	"github.com/bkuri/ppc/internal/model"
	"github.com/bkuri/ppc/internal/render"
	"github.com/bkuri/ppc/internal/resolver"
)

// Compile performs the full prompt compilation pipeline
// Returns (output, metadata, error)
func Compile(opts CompileOptions) (string, CompileMeta, interface{}) {
	// Load modules and rules
	modByID, err := loader.LoadModules(opts.PromptsDir)
	if err != nil {
		return "", CompileMeta{}, err
	}

	rules, err := loader.LoadRules(opts.PromptsDir)
	if err != nil {
		return "", CompileMeta{}, err
	}

	// Build selected IDs from Mode, Contract, Traits
	selectedIDs := buildSelectedIDs(opts)

	// Expand requires (transitive closure)
	closureIDs, fromReq, err := resolver.ExpandRequires(selectedIDs, modByID)
	if err != nil {
		return "", CompileMeta{}, err
	}

	// Build module list with FromReq/Selected flags
	mods, order := buildModuleList(closureIDs, fromReq, selectedIDs, modByID)

	// Validate exclusive groups
	if err := resolver.ValidateExclusiveGroups(rules, mods); err != nil {
		return "", CompileMeta{}, err
	}

	// Sort modules by layer/priority/id
	sortedMods := resolver.SortModules(mods)

	// Render output (LF canonical, single trailing newline)
	out := render.Render(sortedMods, opts.Vars)

	// Compute hash (of canonical output)
	h := sha256.Sum256([]byte(out))
	hash := hex.EncodeToString(h[:])

	// Build metadata
	meta := CompileMeta{
		SelectedIDs: selectedIDs,
		ClosureIDs:  closureIDs,
		Order:       order,
		Hash:        hash,
	}

	return out, meta, nil
}

// buildSelectedIDs constructs the initial selection from options
func buildSelectedIDs(opts CompileOptions) []string {
	selectedIDs := []string{
		"base",
		"modes/" + opts.Mode,
		"contracts/" + opts.Contract,
	}

	// Traits already include "traits/" prefix (e.g., "traits/conservative")
	selectedIDs = append(selectedIDs, opts.Traits...)

	return selectedIDs
}

// buildModuleList populates modules with FromReq/Selected flags and extracts order
func buildModuleList(
	closureIDs []string,
	fromReq map[string]bool,
	selectedIDs []string,
	modByID map[string]*model.Module,
) ([]*model.Module, []string) {
	var mods []*model.Module
	var order []string

	for _, id := range closureIDs {
		m := modByID[id]
		m.FromReq = fromReq[id]
		m.Selected = resolver.Contains(selectedIDs, id)
		mods = append(mods, m)
		order = append(order, id)
	}

	return mods, order
}
