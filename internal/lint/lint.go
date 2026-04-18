package lint

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/voku/vokuknow/internal/config"
	"github.com/voku/vokuknow/internal/generator"
	"github.com/voku/vokuknow/internal/schema"
)

type Result struct {
	Errors   []string
	Warnings []string
}

func (r Result) HasErrors() bool { return len(r.Errors) > 0 }

func Run(repoRoot string) Result {
	result := Result{}
	requiredDirs := []string{
		filepath.Join(repoRoot, config.RootDir, "schema"),
		filepath.Join(repoRoot, config.RootDir, "skills"),
		filepath.Join(repoRoot, config.RootDir, "prompts"),
		filepath.Join(repoRoot, config.RootDir, "claims", "shared"),
		filepath.Join(repoRoot, config.RootDir, "claims", "private"),
		filepath.Join(repoRoot, config.RootDir, "build"),
	}
	for _, dir := range requiredDirs {
		if _, err := os.Stat(dir); err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("missing required directory: %s", dir))
		}
	}

	s, err := schema.Load(repoRoot)
	if err != nil {
		result.Errors = append(result.Errors, err.Error())
		return result
	}

	if len(s.ClaimPolicies.RequiredProvenanceFields) == 0 {
		result.Errors = append(result.Errors, "missing provenance rules")
	}
	if !containsPath(s.Privacy.PrivatePaths, ".vokuknow/claims/private") {
		result.Errors = append(result.Errors, "privacy policy must include .vokuknow/claims/private path")
	}

	artifacts, err := generator.GenerateAll(repoRoot, s)
	if err != nil {
		result.Errors = append(result.Errors, err.Error())
		return result
	}
	sort.Slice(artifacts, func(i, j int) bool { return artifacts[i].Path < artifacts[j].Path })
	for _, artifact := range artifacts {
		existing, err := os.ReadFile(artifact.Path)
		if err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("missing generated artifact: %s", artifact.Path))
			continue
		}
		text := string(existing)
		if strings.Contains(text, "{{") || strings.Contains(text, "}}") {
			result.Errors = append(result.Errors, fmt.Sprintf("unresolved placeholder in %s", artifact.Path))
		}
		if text != string(artifact.Content) {
			result.Errors = append(result.Errors, fmt.Sprintf("drift detected in %s", artifact.Path))
		}
	}

	return result
}

func containsPath(paths []string, target string) bool {
	for _, p := range paths {
		if p == target {
			return true
		}
	}
	return false
}
