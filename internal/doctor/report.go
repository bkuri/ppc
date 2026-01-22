package doctor

import (
	"encoding/json"
	"fmt"
	"os"
)

// DoctorReport represents the JSON output structure
type DoctorReport struct {
	Status   string   `json:"status"`
	Modules  int      `json:"modules"`
	Errors   []string `json:"errors,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

// printDoctorJSON outputs doctor results as JSON
// Returns exit code: 0=ok, 2=failed
func printDoctorJSON(moduleCount int, errs, warns []string, strict bool) int {
	status := "ok"
	exitCode := 0

	if len(errs) > 0 {
		status = "failed"
		exitCode = 2
	} else if strict && len(warns) > 0 {
		status = "failed"
		exitCode = 2
	}

	report := DoctorReport{
		Status:   status,
		Modules:  moduleCount,
		Errors:   errs,
		Warnings: warns,
	}

	b, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "json marshal error: %v\n", err)
		return 2
	}

	fmt.Println(string(b))
	return exitCode
}
