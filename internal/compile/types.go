// Package compile provides the core compilation API
package compile

// CompileOptions specifies the compilation parameters
type CompileOptions struct {
	Mode       string   // "explore", "build", "ship"
	Contract   string   // "markdown", "code"
	Traits     []string // e.g., "traits/conservative", "traits/terse"
	PromptsDir string
	Vars       map[string]string // e.g., {"mode": "explore", "revisions": "1"}
}

// CompileMeta provides metadata about the compilation
type CompileMeta struct {
	SelectedIDs []string // Root selected modules
	ClosureIDs  []string // After requires expansion
	Order       []string // Final module order (IDs)
	Hash        string   // SHA256 of output
}
