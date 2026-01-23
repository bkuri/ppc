package profile

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// LoadProfile loads a built-in profile by name
// Example: LoadProfile("ship") → profiles/ship.yml
// Built-in profiles: explore, build, ship
func LoadProfile(name string) (*Profile, error) {
	path := filepath.Join("profiles", name+".yml")
	return LoadProfileFromFile(path)
}

// LoadProfileFromFile loads a profile from an arbitrary file path
// Supports both absolute and relative paths
// Example: LoadProfileFromFile("./custom.yml")
func LoadProfileFromFile(path string) (*Profile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("profile %q: %w", path, err)
	}

	var p Profile
	if err := yaml.Unmarshal(data, &p); err != nil {
		return nil, fmt.Errorf("profile %q: invalid YAML: %w", path, err)
	}

	if err := p.Validate(); err != nil {
		return nil, err
	}

	return &p, nil
}
