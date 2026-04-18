package schema

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadParsesAndSortsSchema(t *testing.T) {
	repo := t.TempDir()
	writeValidSchema(t, repo)

	s, err := Load(repo)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if got, want := s.EntityTypes[0].Name, "claim"; got != want {
		t.Fatalf("sorted entity order mismatch: got %q want %q", got, want)
	}
	if _, ok := s.PromptProfiles["default"]; !ok {
		t.Fatal("expected default prompt profile")
	}
}

func TestLoadInvalidYAML(t *testing.T) {
	repo := t.TempDir()
	writeValidSchema(t, repo)
	bad := filepath.Join(repo, ".vokuknow", "schema", "entity_types.yaml")
	if err := os.WriteFile(bad, []byte("entities: ["), 0o644); err != nil {
		t.Fatal(err)
	}

	_, err := Load(repo)
	if err == nil || !strings.Contains(err.Error(), "parse entity_types.yaml") {
		t.Fatalf("expected parse error, got %v", err)
	}
}

func writeValidSchema(t *testing.T, repo string) {
	t.Helper()
	dir := filepath.Join(repo, ".vokuknow", "schema")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{
		"AGENTS.md": `provenance audit private shared contradiction crystallization anti-drift`,
		"entity_types.yaml": `entities:
  - name: digest
    description: d
  - name: claim
    description: c
`,
		"relation_types.yaml": `relations:
  - name: supports
    description: s
`,
		"claim_policies.yaml": `required_provenance_fields:
  - source_ref
contradiction_strategy: policy
update_rule: update
crystallization_required: true
`,
		"source_policies.yaml": `authorities:
  - name: docs
    weight: 1
`,
		"privacy_policies.yaml": `redaction_patterns:
  - token
private_paths:
  - .vokuknow/memory/claims/private
promotion_rule: explicit
`,
		"prompt_profiles.yaml": `profiles:
  default:
    goal_template: "goal %s"
    done_template: "done %s"
`,
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}
