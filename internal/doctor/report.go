package doctor

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bkuri/ppc/internal/loader"
	"github.com/bkuri/ppc/internal/model"
	"github.com/bkuri/ppc/internal/resolver"
)

// DoctorStats represents module statistics
type DoctorStats struct {
	Modules     int            `json:"modules"`
	ByLayer     map[string]int `json:"by_layer"`
	Unreachable int            `json:"unreachable"`
	Tags        int            `json:"tags"`
	Groups      int            `json:"groups"`
	Orphaned    int            `json:"orphaned"`
}

// printDoctorJSON outputs doctor results as JSON
// Returns exit code: 0=ok, 2=failed
func printDoctorJSON(moduleCount int, errs, warns []string, strict bool, statsRequested bool) int {
	status := "ok"
	exitCode := 0

	if len(errs) > 0 {
		status = "failed"
		exitCode = 2
	} else if strict && len(warns) > 0 {
		status = "failed"
		exitCode = 2
	}

	// Calculate statistics if requested
	var stats *DoctorStats
	if statsRequested {
		// Calculate by layer
		byLayer := map[string]int{"base": 0, "modes": 0, "traits": 0, "policies": 0, "contracts": 0}
		for _, m := range modByID {
			if _, ok := entry[m]; ok {
				byLayer[layerFromID(m.Layer)]++
			}
		}

		// Count unique tags
		tagValues := map[string]bool{}
		for _, m := range modByID {
			for _, t := range m.Front.Tags {
				if g, v, ok := parseKeyedTag(t); ok {
					tagValues[g+v] = true
				}
			}
		}

		// Count unreachable
		unreachable := 0
		for id := range modByID {
			if !reachable[id] {
				unreachable++
			}
		}

		// Count orphaned requirements (requirements non-existent)
		orphaned := 0
		for _, m := range modByID {
			for _, r := range m.Front.Requires {
				if _, ok := modByID[r]; !ok {
					orphaned++
				}
			}
		}

		stats = &DoctorStats{
			Modules:     moduleCount,
			ByLayer:     byLayer,
			Unreachable: unreachable,
			Tags:        len(tagValues),
			Groups:      len(rules.ExclusiveGroups),
			Orphaned:    orphaned,
		}
	}

	report := DoctorReport{
		Status:   status,
		Modules:  moduleCount,
		Errors:   errs,
		Warnings: warns,
		Stats:    stats,
	}

	b, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "json marshal error: %v\n", err)
		return 2
	}

	fmt.Println(string(b))
	return exitCode
}

// printDoctorJSON outputs doctor results as JSON
// Returns exit code: 0=ok, 2=failed
func printDoctorJSON(moduleCount int, errs, warns []string, strict bool, stats *DoctorStats) int {
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
		Stats:    stats,
	}

	b, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "json marshal error: %v\n", err)
		return 2
	}

	fmt.Println(string(b))
	return exitCode
}

// printDoctorJSON outputs doctor results as JSON
// Returns exit code: 0=ok, 2=failed
func printDoctorJSON(moduleCount int, errs, warns []string, strict bool, statsRequested bool, stats *DoctorStats) int {
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
		Stats:    stats,
	}

	b, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "json marshal error: %v\n", err)
		return 2
	}

	fmt.Println(string(b))
	return exitCode
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
