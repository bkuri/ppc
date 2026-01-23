package doctor

import (
	"fmt"
	"os"

	"github.com/bkuri/ppc/internal/graph"
	"github.com/bkuri/ppc/internal/model"
)

// printDoctorGraph outputs graph in DOT format
// Returns exit code: 0=ok, 2=failed
func printDoctorGraph(modByID map[string]*model.Module, rules *model.Rules,
	reachable map[string]bool, outPath string) int {
	dotOutput := graph.BuildDOT(modByID, rules, reachable)

	if outPath != "" {
		if err := os.WriteFile(outPath, []byte(dotOutput), 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "failed to write graph: %v\n", err)
			return 2
		}
	} else {
		fmt.Print(dotOutput)
	}

	return 0
}
