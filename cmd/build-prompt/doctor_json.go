package main

import (
	"encoding/json"
	"fmt"
)

type DoctorReport struct {
	Status   string   `json:"status"`
	Modules  int      `json:"modules"`
	Errors   []string `json:"errors,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

func printDoctorJSON(
	modCount int,
	errors []string,
	warnings []string,
	strict bool,
) int {
	status := "ok"
	if len(errors) > 0 || (strict && len(warnings) > 0) {
		status = "fail"
	}

	rep := DoctorReport{
		Status:   status,
		Modules: modCount,
		Errors:   errors,
		Warnings: warnings,
	}

	b, _ := json.MarshalIndent(rep, "", "  ")
	fmt.Println(string(b))

	if status == "fail" {
		return 2
	}
	return 0
}
