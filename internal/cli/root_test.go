package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLIInitBuildLint(t *testing.T) {
	repo := t.TempDir()
	var out bytes.Buffer
	var errOut bytes.Buffer

	if err := Execute([]string{"--repo", repo, "init"}, &out, &errOut); err != nil {
		t.Fatalf("init failed: %v", err)
	}
	if _, err := os.Stat(filepath.Join(repo, ".vokuknow", "schema", "AGENTS.md")); err != nil {
		t.Fatalf("missing AGENTS.md: %v", err)
	}
	for _, path := range []string{
		filepath.Join(repo, ".vokuknow", "skills", "vokuknow", "SKILL.md"),
		filepath.Join(repo, ".vokuknow", "templates", "discovery.md"),
		filepath.Join(repo, ".vokuknow", "templates", "decision-entry.md"),
		filepath.Join(repo, ".vokuknow", "policy", "memory_rules.md"),
		filepath.Join(repo, ".vokuknow", "prompts", "examples", "code-discovery.md"),
		filepath.Join(repo, ".vokuknow", "prompts", "examples", "decision-log.md"),
		filepath.Join(repo, ".vokuknow", "audit", "decision-log.md"),
		filepath.Join(repo, ".vokuknow", "memory", "discoveries"),
		filepath.Join(repo, ".vokuknow", "memory", "claims", "private"),
		filepath.Join(repo, ".vokuknow", "memory", "digests"),
		filepath.Join(repo, ".vokuknow", "memory", "handoffs"),
	} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("missing bootstrapped path %s: %v", path, err)
		}
	}

	out.Reset()
	if err := Execute([]string{"--repo", repo, "build"}, &out, &errOut); err != nil {
		t.Fatalf("build failed: %v", err)
	}
	if !strings.Contains(out.String(), "build complete") {
		t.Fatalf("unexpected build output: %q", out.String())
	}

	if err := Execute([]string{"--repo", repo, "lint"}, &out, &errOut); err != nil {
		t.Fatalf("lint failed: %v", err)
	}
}
