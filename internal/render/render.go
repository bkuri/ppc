// Package render provides output rendering
package render

import (
	"strings"

	"github.com/bkuri/ppc/internal/model"
)

// Render concatenates module bodies with variable substitution
// Ensures LF line endings and single trailing newline (canonical output)
func Render(mods []*model.Module, vars map[string]string) string {
	var b strings.Builder
	for i, m := range mods {
		if i > 0 {
			b.WriteString("\n\n")
		}
		body := m.Body
		for k, v := range vars {
			body = strings.ReplaceAll(body, "{{"+k+"}}", v)
		}
		b.WriteString(strings.TrimRight(body, "\n"))
	}
	// Ensure single trailing newline (canonical)
	return strings.TrimRight(b.String(), "\n") + "\n"
}
