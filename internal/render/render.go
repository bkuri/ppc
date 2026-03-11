package render

import (
	"strings"

	"github.com/bkuri/ppc/internal/model"
	"github.com/bkuri/ppc/internal/substitute"
)

func Render(mods []*model.Module, vars substitute.Vars) string {
	var b strings.Builder
	for i, m := range mods {
		if i > 0 {
			b.WriteString("\n\n")
		}
		body := substitute.Substitute(m.Body, vars)
		b.WriteString(strings.TrimRight(body, "\n"))
	}
	return strings.TrimRight(b.String(), "\n") + "\n"
}
