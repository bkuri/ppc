// Package compile provides the core compilation API
package compile

type CompileOptions struct {
	Mode       string
	Contract   string
	Traits     []string
	PromptsDir string
	VarsFile   string
	Vars       map[string]any
}

// CompileMeta provides metadata about the compilation
type CompileMeta struct {
	SelectedIDs []string // Root selected modules
	ClosureIDs  []string // After requires expansion
	Order       []string // Final module order (IDs)
	Hash        string   // SHA256 of output
}
