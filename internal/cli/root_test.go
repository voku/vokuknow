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
