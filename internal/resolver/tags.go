package resolver

import "strings"

// ParseKeyedTag parses a keyed tag format (group:value)
// Returns (group, value, ok)
func ParseKeyedTag(t string) (group, value string, ok bool) {
	i := strings.IndexByte(t, ':')
	if i <= 0 || i >= len(t)-1 {
		return "", "", false
	}
	return t[:i], t[i+1:], true
}

// Contains checks if a string exists in a slice
func Contains(xs []string, s string) bool {
	for _, x := range xs {
		if x == s {
			return true
		}
	}
	return false
}
