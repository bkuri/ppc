package loader

import (
	"fmt"
	"strings"

	errtypes "github.com/bkuri/ppc/internal/error"
	"github.com/bkuri/ppc/internal/model"
	"gopkg.in/yaml.v3"
)

// ParseFrontmatter extracts YAML frontmatter and body from raw content
func ParseFrontmatter(raw []byte) (model.Frontmatter, string, bool, errtypes.SrcError) {
	s := string(raw)
	if !strings.HasPrefix(s, "---\n") && !strings.HasPrefix(s, "---\r\n") {
		return model.Frontmatter{}, strings.TrimRight(s, "\n"), false, errtypes.SrcError{}
	}

	idx := strings.Index(s[4:], "\n---\n")
	delimLen := len("\n---\n")
	if idx == -1 {
		idx = strings.Index(s[4:], "\r\n---\r\n")
		delimLen = len("\r\n---\r\n")
	}
	if idx == -1 {
		return model.Frontmatter{}, "", false,
			errtypes.New("", "", "frontmatter start found but missing closing ---")
	}

	yml := s[4 : 4+idx]
	body := s[4+idx+delimLen:]
	body = strings.TrimLeft(body, "\r\n")
	body = strings.TrimRight(body, "\n")

	var fm model.Frontmatter
	if err := yaml.Unmarshal([]byte(yml), &fm); err != nil {
		return model.Frontmatter{}, "", false,
			errtypes.New("", "", fmt.Sprintf("invalid YAML frontmatter: %v", err))
	}
	return fm, body, true, errtypes.SrcError{}
}
