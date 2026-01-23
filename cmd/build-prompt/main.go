package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/bkuri/ppc/internal/compile"
	"github.com/bkuri/ppc/internal/doctor"
	profilepkg "github.com/bkuri/ppc/internal/profile"
)

// dief prints error to stderr and exits
func dief(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(2)
}

// buildTraits converts boolean CLI flags to explicit module IDs
func buildTraits(conservative, creative, terse, verbose bool) []string {
	var traits []string
	if conservative {
		traits = append(traits, "traits/conservative")
	}
	if creative {
		traits = append(traits, "traits/creative")
	}
	if terse {
		traits = append(traits, "traits/terse")
	}
	if verbose {
		traits = append(traits, "traits/verbose")
	}
	return traits
}

// explainOutput prints compilation metadata to stderr (CLI concern)
func explainOutput(meta compile.CompileMeta) {
	fmt.Fprintln(os.Stderr, "PPC explain")

	fmt.Fprintln(os.Stderr, "Selected IDs:")
	sel := append([]string{}, meta.SelectedIDs...)
	sort.Strings(sel)
	for _, id := range sel {
		fmt.Fprintf(os.Stderr, "  - %s\n", id)
	}

	fmt.Fprintln(os.Stderr, "Closure IDs (after requires):")
	cls := append([]string{}, meta.ClosureIDs...)
	sort.Strings(cls)
	for _, id := range cls {
		fmt.Fprintln(os.Stderr, "  - "+id)
	}

	fmt.Fprintln(os.Stderr, "Final order:")
	for _, id := range meta.Order {
		fmt.Fprintf(os.Stderr, "  - %s\n", id)
	}
}

func runExplore(args []string, promptsDir string) {
	fs := flag.NewFlagSet("explore", flag.ExitOnError)

	profile := fs.String("profile", "", "load preset configuration (e.g., ship)")
	conservative := fs.Bool("conservative", false, "include traits/conservative")
	creative := fs.Bool("creative", false, "include traits/creative")
	terse := fs.Bool("terse", false, "include traits/terse")
	verbose := fs.Bool("verbose", false, "include traits/verbose")
	revisions := fs.Int("revisions", -1, "revision budget (enables policies/revisions)")
	contract := fs.String("contract", "markdown", "contract module (code|markdown)")
	outPath := fs.String("out", "", "write output to file")
	explain := fs.Bool("explain", false, "explain resolution steps to stderr")
	withHash := fs.Bool("hash", false, "prepend prompt-id hash header")
	proDir := fs.String("prompts", promptsDir, "prompts directory")

	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, `usage:
  ppc explore [flags]

Explore mode generates a prompt for exploration tasks.

flags:`)
		fs.PrintDefaults()
	}

	fs.Parse(args)

	var opts *compile.CompileOptions
	var err error

	if *profile != "" {
		prof, err := profile.LoadProfile(*profile)
		if err != nil {
			dief("profile error: %v", err)
		}

		mergeOpts := profile.MergeOptions{
			Conservative: toBoolPtr(*conservative),
			Creative:     toBoolPtr(*creative),
			Terse:        toBoolPtr(*terse),
			Verbose:      toBoolPtr(*verbose),
			Revisions:    toIntPtr(*revisions),
			Contract:     toStringPtr(*contract),
		}

		opts, err = profile.Merge(prof, mergeOpts)
		if err != nil {
			dief("merge error: %v", err)
		}
		opts.PromptsDir = *proDir
	} else {
		traits := buildTraits(*conservative, *creative, *terse, *verbose)

		vars := map[string]string{"mode": "explore"}
		if *revisions >= 0 {
			vars["revisions"] = strconv.Itoa(*revisions)
		}

		opts = &compile.CompileOptions{
			Mode:       "explore",
			Contract:   *contract,
			Traits:     traits,
			PromptsDir: *proDir,
			Vars:       vars,
		}
	}

	out, meta, err := compile.Compile(opts)
	if err != nil {
		dief("%v", err)
	}

	if *withHash {
		out = fmt.Sprintf("<!-- prompt-id: sha256:%s -->\n\n%s", meta.Hash, out)
	}

	if *explain {
		explainOutput(meta)
	}

	if *outPath != "" {
		if err := os.WriteFile(*outPath, []byte(out), 0o644); err != nil {
			dief("failed to write %s: %v", *outPath, err)
		}
	}
	fmt.Print(out)
}

func runBuild(args []string, promptsDir string) {
	fs := flag.NewFlagSet("build", flag.ExitOnError)

	conservative := fs.Bool("conservative", false, "include traits/conservative")
	creative := fs.Bool("creative", false, "include traits/creative")
	terse := fs.Bool("terse", false, "include traits/terse")
	verbose := fs.Bool("verbose", false, "include traits/verbose")
	revisions := fs.Int("revisions", -1, "revision budget (enables policies/revisions)")
	contract := fs.String("contract", "markdown", "contract module (code|markdown)")
	outPath := fs.String("out", "", "write output to file")
	explain := fs.Bool("explain", false, "explain resolution steps to stderr")
	withHash := fs.Bool("hash", false, "prepend prompt-id hash header")
	proDir := fs.String("prompts", promptsDir, "prompts directory")

	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, `usage:
  ppc build [flags]

Build mode generates a prompt for building/implementing features.

flags:`)
		fs.PrintDefaults()
	}

	fs.Parse(args)

	traits := buildTraits(*conservative, *creative, *terse, *verbose)

	vars := map[string]string{"mode": "build"}
	if *revisions >= 0 {
		vars["revisions"] = strconv.Itoa(*revisions)
	}

	opts := compile.CompileOptions{
		Mode:       "build",
		Contract:   *contract,
		Traits:     traits,
		PromptsDir: *proDir,
		Vars:       vars,
	}

	out, meta, err := compile.Compile(opts)
	if err != nil {
		dief("%v", err)
	}

	if *withHash {
		out = fmt.Sprintf("<!-- prompt-id: sha256:%s -->\n\n%s", meta.Hash, out)
	}

	if *explain {
		explainOutput(meta)
	}

	if *outPath != "" {
		if err := os.WriteFile(*outPath, []byte(out), 0o644); err != nil {
			dief("failed to write %s: %v", *outPath, err)
		}
	}
	fmt.Print(out)
}

func runShip(args []string, promptsDir string) {
	fs := flag.NewFlagSet("ship", flag.ExitOnError)

	conservative := fs.Bool("conservative", false, "include traits/conservative")
	creative := fs.Bool("creative", false, "include traits/creative")
	terse := fs.Bool("terse", false, "include traits/terse")
	verbose := fs.Bool("verbose", false, "include traits/verbose")
	revisions := fs.Int("revisions", -1, "revision budget (enables policies/revisions)")
	contract := fs.String("contract", "markdown", "contract module (code|markdown)")
	outPath := fs.String("out", "", "write output to file")
	explain := fs.Bool("explain", false, "explain resolution steps to stderr")
	withHash := fs.Bool("hash", false, "prepend prompt-id hash header")
	proDir := fs.String("prompts", promptsDir, "prompts directory")

	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, `usage:
  ppc ship [flags]

Ship mode generates a prompt for release/deployment tasks.

flags:`)
		fs.PrintDefaults()
	}

	fs.Parse(args)

	traits := buildTraits(*conservative, *creative, *terse, *verbose)

	vars := map[string]string{"mode": "ship"}
	if *revisions >= 0 {
		vars["revisions"] = strconv.Itoa(*revisions)
	}

	opts := compile.CompileOptions{
		Mode:       "ship",
		Contract:   *contract,
		Traits:     traits,
		PromptsDir: *proDir,
		Vars:       vars,
	}

	out, meta, err := compile.Compile(opts)
	if err != nil {
		dief("%v", err)
	}

	if *withHash {
		out = fmt.Sprintf("<!-- prompt-id: sha256:%s -->\n\n%s", meta.Hash, out)
	}

	if *explain {
		explainOutput(meta)
	}

	if *outPath != "" {
		if err := os.WriteFile(*outPath, []byte(out), 0o644); err != nil {
			dief("failed to write %s: %v", *outPath, err)
		}
	}
	fmt.Print(out)
}

func printGlobalUsage() {
	fmt.Fprintln(os.Stderr, `usage:
  ppc <subcommand> [flags]

subcommands:
  explore    Generate prompt for exploration mode
  build      Generate prompt for build mode
  ship       Generate prompt for shipping mode
  doctor     Validate module structure and dependencies

global flags:
  --list     List all available modules
  --help     Show this help message

examples:
  ppc explore --conservative --revisions 1 --contract markdown
  ppc build --conservative --revisions 1 --contract code --explain
  ppc ship --creative --out AGENTS.md --hash
  ppc doctor --strict --json

run 'ppc <subcommand> --help' for subcommand-specific options`)
}

func handleListModules(promptsDir string) {
	modByID, err := loader.LoadModules(promptsDir)
	if err != nil {
		dief("%v", err)
	}

	var ids []string
	for id := range modByID {
		ids = append(ids, id)
	}

	sort.Strings(ids)
	for _, id := range ids {
		m := modByID[id]
		desc := m.Front.Desc
		if desc == "" {
			desc = "(no desc)"
		}
		fmt.Printf("%-22s  %s\n", id, desc)
	}
}

func main() {
	if len(os.Args) < 2 {
		printGlobalUsage()
		os.Exit(1)
	}

	subcommand := os.Args[1]
	args := os.Args[2:]

	// Default prompts directory
	promptsDir := "prompts"

	// Handle global meta-flags first
	if subcommand == "--list" {
		handleListModules(promptsDir)
		os.Exit(0)
	}

	if subcommand == "--help" || subcommand == "-h" || subcommand == "help" {
		printGlobalUsage()
		os.Exit(0)
	}

	// Dispatch to subcommand
	switch subcommand {
	case "explore":
		runExplore(args, promptsDir)
	case "build":
		runBuild(args, promptsDir)
	case "ship":
		runShip(args, promptsDir)
	case "doctor":
		fs := flag.NewFlagSet("doctor", flag.ExitOnError)
		strict := fs.Bool("strict", false, "treat warnings as errors")
		jsonOut := fs.Bool("json", false, "output machine-readable JSON")
		withStats := fs.Bool("stats", false, "include module statistics in JSON output")
		graphOut := fs.Bool("graph", false, "output Graphviz DOT format")
		outPath := fs.String("out", "", "write output to file")
		proDir := fs.String("prompts", promptsDir, "prompts directory")
		fs.Usage = func() {
			fmt.Fprintln(os.Stderr, `usage:
  ppc doctor [flags]

Checks module integrity, requires, cycles, and tag/rules sanity.

flags:`)
			fs.PrintDefaults()
		}
		fs.Parse(args)
		os.Exit(doctor.RunDoctor(*proDir, *strict, *jsonOut, *withStats, *graphOut, *outPath))

	default:
		fmt.Fprintf(os.Stderr, "unknown subcommand: %s\n", subcommand)
		printGlobalUsage()
		os.Exit(1)
	}
}
