package generator

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/voku/vokuknow/internal/config"
	"github.com/voku/vokuknow/internal/schema"
)

type Artifact struct {
	Path    string
	Content []byte
}

type Manifest struct {
	Version   int                `json:"version"`
	Artifacts []ManifestArtifact `json:"artifacts"`
}

type ManifestArtifact struct {
	Path   string `json:"path"`
	SHA256 string `json:"sha256"`
}

func GenerateAll(repoRoot string, s *schema.Schema) ([]Artifact, error) {
	artifacts := make([]Artifact, 0, len(config.DefaultSkills)+len(config.DefaultPrompts)+2)
	for _, name := range config.DefaultSkills {
		content, err := GenerateSkill(name, s)
		if err != nil {
			return nil, err
		}
		artifacts = append(artifacts, Artifact{Path: filepath.Join(repoRoot, config.RootDir, "skills", name+".md"), Content: content})
	}
	for _, name := range config.DefaultPrompts {
		content, err := GeneratePrompt(name, s)
		if err != nil {
			return nil, err
		}
		artifacts = append(artifacts, Artifact{Path: filepath.Join(repoRoot, config.RootDir, "prompts", name+".md"), Content: content})
	}
	instructionBytes, err := GenerateInstructionsJSON(s)
	if err != nil {
		return nil, err
	}
	artifacts = append(artifacts, Artifact{Path: filepath.Join(repoRoot, config.RootDir, "build", "instructions.json"), Content: instructionBytes})

	manifestBytes, err := GenerateManifestJSON(repoRoot, artifacts)
	if err != nil {
		return nil, err
	}
	artifacts = append(artifacts, Artifact{Path: filepath.Join(repoRoot, config.RootDir, "build", "manifest.json"), Content: manifestBytes})

	sort.Slice(artifacts, func(i, j int) bool { return artifacts[i].Path < artifacts[j].Path })
	return artifacts, nil
}

func GenerateSkill(name string, s *schema.Schema) ([]byte, error) {
	const tpl = `# Skill: {{.Name}}

## Purpose
{{.Purpose}}

## Repo locations
- Schema: .vokuknow/schema/
- Memory discoveries: .vokuknow/memory/discoveries/
- Memory claims: .vokuknow/memory/claims/
- Memory private claims: .vokuknow/memory/claims/private/
- Memory digests: .vokuknow/memory/digests/
- Memory handoffs: .vokuknow/memory/handoffs/
- Sources raw: .vokuknow/sources/raw/
- Sources normalized: .vokuknow/sources/normalized/
- Audit: .vokuknow/audit/

## Required workflow
1. Inspect existing memory artifacts and sources before writing new claims or digests.
2. Preserve provenance fields: {{.Provenance}}.
3. Follow contradiction strategy: {{.Contradiction}}.
4. Enforce private/shared boundary using private paths: {{.PrivatePaths}}.
5. Apply redaction rules before storing: {{.Redaction}}.
6. Log self-directed decisions and missing guidance in .vokuknow/audit/decision-log.md.
7. Append auditable entries for edits, promotions, deletions, and resolutions.
8. Crystallize finished work into digests when complete.

## Quality rules
- Use deterministic language.
- Prefer authoritative, recent, and multi-supported claims.
- Never leave unresolved placeholders.

## Anti-drift constraints
- Do not invent schema fields.
- Do not skip provenance.
- Keep shared/private policy consistent.
`
	t := template.Must(template.New("skill").Parse(tpl))
	data := map[string]string{
		"Name":          name,
		"Purpose":       skillPurpose(name),
		"Provenance":    strings.Join(s.ClaimPolicies.RequiredProvenanceFields, ", "),
		"Contradiction": s.ClaimPolicies.ContradictionStrategy,
		"PrivatePaths":  strings.Join(s.Privacy.PrivatePaths, ", "),
		"Redaction":     strings.Join(s.Privacy.RedactionPatterns, ", "),
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return nil, err
	}
	return []byte(strings.TrimSpace(buf.String()) + "\n"), nil
}

func GeneratePrompt(name string, s *schema.Schema) ([]byte, error) {
	profile := s.PromptProfiles["default"]
	const tpl = `# Prompt: {{.Name}}

## Task
{{.Task}}

## Goal
{{.Goal}}

## Required steps
1. Read .vokuknow/schema/AGENTS.md and relevant policy YAML.
2. Inspect existing memory artifacts, sources, and digests before changes.
3. Apply provenance requirements: {{.Provenance}}.
4. Enforce contradiction resolution: {{.Contradiction}}.
5. Enforce privacy boundary and redaction.
6. If you must decide because docs or skills are insufficient, append an entry to .vokuknow/audit/decision-log.md and either update the missing guidance in the same task or leave a concrete follow-up in that entry.
7. Write auditable updates with deterministic formatting.
8. If work is complete, crystallize reusable lessons into digests.

## Constraints
- Deterministic output only.
- No unresolved placeholders.
- Keep shared/private boundaries intact.
- Prefer recent authoritative sources with multiple observations.

## Done condition
{{.Done}}

## Output expectations
- Updated files under .vokuknow/ with clear diffs.
- Decision-log updates when missing guidance forced a material choice, plus either same-task guidance updates or explicit follow-up notes.
- Audit-safe edits with provenance.
`
	t := template.Must(template.New("prompt").Parse(tpl))
	data := map[string]string{
		"Name":          name,
		"Task":          promptTask(name),
		"Goal":          fmt.Sprintf(profile.GoalTemplate, name),
		"Done":          fmt.Sprintf(profile.DoneTemplate, name),
		"Provenance":    strings.Join(s.ClaimPolicies.RequiredProvenanceFields, ", "),
		"Contradiction": s.ClaimPolicies.ContradictionStrategy,
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return nil, err
	}
	return []byte(strings.TrimSpace(buf.String()) + "\n"), nil
}

func GenerateInstructionsJSON(s *schema.Schema) ([]byte, error) {
	payload := map[string]any{
		"schema": map[string]any{
			"entity_types":   s.EntityTypes,
			"relation_types": s.RelationTypes,
		},
		"policies": map[string]any{
			"claims":  s.ClaimPolicies,
			"sources": s.SourcePolicies,
			"privacy": s.Privacy,
		},
		"skills":  config.DefaultSkills,
		"prompts": config.DefaultPrompts,
	}
	return marshalJSON(payload)
}

func GenerateManifestJSON(repoRoot string, artifacts []Artifact) ([]byte, error) {
	items := make([]ManifestArtifact, 0, len(artifacts))
	for _, a := range artifacts {
		h := sha256.Sum256(a.Content)
		rel := strings.TrimPrefix(filepath.ToSlash(a.Path), filepath.ToSlash(repoRoot)+"/")
		items = append(items, ManifestArtifact{Path: rel, SHA256: hex.EncodeToString(h[:])})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].Path < items[j].Path })
	return marshalJSON(Manifest{Version: 1, Artifacts: items})
}

func marshalJSON(v any) ([]byte, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(b, '\n'), nil
}

func skillPurpose(name string) string {
	switch name {
	case "wiki-operator":
		return "Operate the repo-local wiki with deterministic claim/query/update behavior."
	case "crystallizer":
		return "Distill completed work into durable digests and reusable lessons."
	case "contradiction-resolver":
		return "Resolve conflicting claims using policy-driven precedence and visible outcomes."
	default:
		return "Execute schema-driven wiki operations deterministically."
	}
}

func promptTask(name string) string {
	switch name {
	case "ingest":
		return "Ingest new source material into normalized wiki knowledge artifacts."
	case "query":
		return "Query the existing wiki and return evidence-backed answers."
	case "update-claim":
		return "Update or create a claim according to policy and provenance requirements."
	case "resolve-claim":
		return "Resolve contradictory claims and keep both default outcome and human override visible."
	case "crystallize":
		return "Crystallize completed work into durable digests and reusable patterns."
	case "handoff":
		return "Prepare deterministic handoff notes from current wiki state and audit trail."
	case "lint":
		return "Lint the knowledge contract and report policy or structure violations."
	default:
		return "Execute deterministic wiki workflow."
	}
}
