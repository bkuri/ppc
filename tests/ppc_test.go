package tests

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestBasicCompile(t *testing.T) {
	cmd := exec.Command("./build-prompt",
		"--conservative",
		"--revisions", "1",
		"--contract", "markdown",
		"explore",
	)
	cmd.Dir = ".."
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		t.Fatalf("command failed: %v\n%s", err, out.String())
	}

	s := out.String()
	if !strings.Contains(s, "Mode: Explore") {
		t.Fatalf("expected explore mode output, got:\n%s", s)
	}

	if !strings.Contains(s, "Trait: Conservative") {
		t.Fatalf("expected conservative trait, got:\n%s", s)
	}
}
