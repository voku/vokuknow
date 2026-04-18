package schema

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/voku/vokuknow/internal/config"
	"gopkg.in/yaml.v3"
)

func Load(repoRoot string) (*Schema, error) {
	schemaDir := filepath.Join(repoRoot, config.RootDir, "schema")
	for _, file := range config.RequiredSchemaFiles {
		if _, err := os.Stat(filepath.Join(schemaDir, file)); err != nil {
			return nil, fmt.Errorf("missing required schema file %s: %w", file, err)
		}
	}

	agentsPath := filepath.Join(schemaDir, "AGENTS.md")
	agents, err := os.ReadFile(agentsPath)
	if err != nil {
		return nil, fmt.Errorf("read AGENTS.md: %w", err)
	}

	s := &Schema{AgentsMarkdown: string(agents)}
	if err := parseYAMLFile(filepath.Join(schemaDir, "entity_types.yaml"), &EntityTypesFile{}, func(v any) {
		s.EntityTypes = v.(*EntityTypesFile).Entities
	}); err != nil {
		return nil, err
	}
	if err := parseYAMLFile(filepath.Join(schemaDir, "relation_types.yaml"), &RelationTypesFile{}, func(v any) {
		s.RelationTypes = v.(*RelationTypesFile).Relations
	}); err != nil {
		return nil, err
	}
	if err := parseYAMLFile(filepath.Join(schemaDir, "claim_policies.yaml"), &s.ClaimPolicies, nil); err != nil {
		return nil, err
	}
	if err := parseYAMLFile(filepath.Join(schemaDir, "source_policies.yaml"), &s.SourcePolicies, nil); err != nil {
		return nil, err
	}
	if err := parseYAMLFile(filepath.Join(schemaDir, "privacy_policies.yaml"), &s.Privacy, nil); err != nil {
		return nil, err
	}
	profiles := &PromptProfilesFile{}
	if err := parseYAMLFile(filepath.Join(schemaDir, "prompt_profiles.yaml"), profiles, nil); err != nil {
		return nil, err
	}
	s.PromptProfiles = profiles.Profiles

	sort.Slice(s.EntityTypes, func(i, j int) bool { return s.EntityTypes[i].Name < s.EntityTypes[j].Name })
	sort.Slice(s.RelationTypes, func(i, j int) bool { return s.RelationTypes[i].Name < s.RelationTypes[j].Name })
	sort.Slice(s.SourcePolicies.Authorities, func(i, j int) bool {
		return s.SourcePolicies.Authorities[i].Name < s.SourcePolicies.Authorities[j].Name
	})
	sort.Strings(s.ClaimPolicies.RequiredProvenanceFields)
	sort.Strings(s.Privacy.RedactionPatterns)
	sort.Strings(s.Privacy.PrivatePaths)

	if err := Validate(s); err != nil {
		return nil, err
	}

	return s, nil
}

func Validate(s *Schema) error {
	if strings.TrimSpace(s.AgentsMarkdown) == "" {
		return errors.New("AGENTS.md is empty")
	}
	requiredTerms := []string{"provenance", "audit", "private", "shared", "contradiction", "crystallization", "anti-drift"}
	lower := strings.ToLower(s.AgentsMarkdown)
	for _, term := range requiredTerms {
		if !strings.Contains(lower, term) {
			return fmt.Errorf("AGENTS.md missing required guidance term: %s", term)
		}
	}
	if len(s.EntityTypes) == 0 {
		return errors.New("entity_types.yaml must define at least one entity")
	}
	if len(s.RelationTypes) == 0 {
		return errors.New("relation_types.yaml must define at least one relation")
	}
	if len(s.ClaimPolicies.RequiredProvenanceFields) == 0 {
		return errors.New("claim_policies.yaml must define required_provenance_fields")
	}
	if s.ClaimPolicies.ContradictionStrategy == "" {
		return errors.New("claim_policies.yaml must define contradiction_strategy")
	}
	if len(s.Privacy.PrivatePaths) == 0 {
		return errors.New("privacy_policies.yaml must define private_paths")
	}
	if _, ok := s.PromptProfiles["default"]; !ok {
		return errors.New("prompt_profiles.yaml must include profiles.default")
	}
	return nil
}

func parseYAMLFile(path string, target any, post func(any)) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", filepath.Base(path), err)
	}
	if err := yaml.Unmarshal(b, target); err != nil {
		return fmt.Errorf("parse %s: %w", filepath.Base(path), err)
	}
	if post != nil {
		post(target)
	}
	return nil
}
