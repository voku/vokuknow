package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/voku/vokuknow/internal/schema"
)

func testSchema() *schema.Schema {
	return &schema.Schema{
		AgentsMarkdown: "provenance audit private shared contradiction crystallization anti-drift",
		EntityTypes:    []schema.TypeDef{{Name: "claim", Description: "c"}},
		RelationTypes:  []schema.TypeDef{{Name: "supports", Description: "s"}},
		ClaimPolicies: schema.ClaimPolicies{
			RequiredProvenanceFields: []string{"source_ref", "observed_at"},
			ContradictionStrategy:    "prefer recent and authoritative",
			UpdateRule:               "update matching keys",
			CrystallizationRequired:  true,
		},
		SourcePolicies: schema.SourcePolicies{Authorities: []schema.Authority{{Name: "docs", Weight: 10}}},
		Privacy: schema.PrivacyPolicies{
			RedactionPatterns: []string{"token"},
			PrivatePaths:      []string{".vokuknow/memory/claims/private"},
			PromotionRule:     "manual",
		},
		PromptProfiles: map[string]schema.PromptProfile{
			"default": {GoalTemplate: "Goal for %s", DoneTemplate: "Done for %s"},
		},
	}
}

func TestGenerateSkillGolden(t *testing.T) {
	got, err := GenerateSkill("wiki-operator", testSchema())
	if err != nil {
		t.Fatal(err)
	}
	want, err := os.ReadFile(filepath.Join("testdata", "wiki-operator.golden.md"))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(want) {
		t.Fatalf("skill output mismatch\n--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}

func TestGeneratePromptGolden(t *testing.T) {
	got, err := GeneratePrompt("ingest", testSchema())
	if err != nil {
		t.Fatal(err)
	}
	want, err := os.ReadFile(filepath.Join("testdata", "ingest.golden.md"))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(want) {
		t.Fatalf("prompt output mismatch\n--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}

func TestGenerateAllDeterministic(t *testing.T) {
	repo := t.TempDir()
	first, err := GenerateAll(repo, testSchema())
	if err != nil {
		t.Fatal(err)
	}
	second, err := GenerateAll(repo, testSchema())
	if err != nil {
		t.Fatal(err)
	}
	if len(first) != len(second) {
		t.Fatalf("artifact length mismatch: %d vs %d", len(first), len(second))
	}
	for i := range first {
		if first[i].Path != second[i].Path || string(first[i].Content) != string(second[i].Content) {
			t.Fatalf("artifact mismatch at index %d", i)
		}
	}
}
