package app

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/voku/vokuknow/internal/config"
	"github.com/voku/vokuknow/internal/doctor"
	"github.com/voku/vokuknow/internal/fs"
	"github.com/voku/vokuknow/internal/generator"
	"github.com/voku/vokuknow/internal/lint"
	"github.com/voku/vokuknow/internal/schema"
)

type App struct {
	RepoRoot string
}

func New(repoRoot string) *App {
	return &App{RepoRoot: repoRoot}
}

func (a *App) Init() ([]string, error) {
	dirs := []string{
		filepath.Join(a.RepoRoot, config.RootDir, "schema"),
		filepath.Join(a.RepoRoot, config.RootDir, "skills"),
		filepath.Join(a.RepoRoot, config.RootDir, "skills", "vokuknow"),
		filepath.Join(a.RepoRoot, config.RootDir, "prompts"),
		filepath.Join(a.RepoRoot, config.RootDir, "prompts", "examples"),
		filepath.Join(a.RepoRoot, config.RootDir, "policy"),
		filepath.Join(a.RepoRoot, config.RootDir, "templates"),
		filepath.Join(a.RepoRoot, config.RootDir, "memory", "discoveries"),
		filepath.Join(a.RepoRoot, config.RootDir, "memory", "claims"),
		filepath.Join(a.RepoRoot, config.RootDir, "memory", "claims", "private"),
		filepath.Join(a.RepoRoot, config.RootDir, "memory", "digests"),
		filepath.Join(a.RepoRoot, config.RootDir, "memory", "handoffs"),
		filepath.Join(a.RepoRoot, config.RootDir, "claims", "shared"),
		filepath.Join(a.RepoRoot, config.RootDir, "claims", "private"),
		filepath.Join(a.RepoRoot, config.RootDir, "sources", "raw"),
		filepath.Join(a.RepoRoot, config.RootDir, "sources", "normalized"),
		filepath.Join(a.RepoRoot, config.RootDir, "digests"),
		filepath.Join(a.RepoRoot, config.RootDir, "audit"),
		filepath.Join(a.RepoRoot, config.RootDir, "build"),
		filepath.Join(a.RepoRoot, config.RootDir, "state"),
	}
	for _, d := range dirs {
		if err := fs.EnsureDir(d); err != nil {
			return nil, err
		}
	}

	starter := starterFiles()
	written := make([]string, 0, len(starter))
	keys := make([]string, 0, len(starter))
	for path := range starter {
		keys = append(keys, path)
	}
	sort.Strings(keys)
	for _, rel := range keys {
		changed, err := fs.WriteIfChanged(filepath.Join(a.RepoRoot, rel), []byte(starter[rel]))
		if err != nil {
			return nil, err
		}
		if changed {
			written = append(written, rel)
		}
	}
	if _, err := a.Build(); err != nil {
		return nil, err
	}
	return written, nil
}

func (a *App) Build() ([]string, error) {
	s, err := schema.Load(a.RepoRoot)
	if err != nil {
		return nil, err
	}
	artifacts, err := generator.GenerateAll(a.RepoRoot, s)
	if err != nil {
		return nil, err
	}
	changed := make([]string, 0, len(artifacts))
	for _, artifact := range artifacts {
		didChange, err := fs.WriteIfChanged(artifact.Path, artifact.Content)
		if err != nil {
			return nil, err
		}
		if didChange {
			changed = append(changed, strings.TrimPrefix(filepath.ToSlash(artifact.Path), filepath.ToSlash(a.RepoRoot)+"/"))
		}
	}
	sort.Strings(changed)
	return changed, nil
}

func (a *App) Lint() error {
	result := lint.Run(a.RepoRoot)
	if result.HasErrors() {
		return fmt.Errorf("lint failed:\n- %s", strings.Join(result.Errors, "\n- "))
	}
	return nil
}

func (a *App) Doctor() (doctor.Report, error) {
	report := doctor.Run(a.RepoRoot)
	if len(report.Issues) > 0 {
		return report, fmt.Errorf("doctor found %d issue(s)", len(report.Issues))
	}
	return report, nil
}

func (a *App) ShowSkill(name string) (string, error) {
	b, err := os.ReadFile(filepath.Join(a.RepoRoot, config.RootDir, "skills", name+".md"))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (a *App) ShowPrompt(name string) (string, error) {
	b, err := os.ReadFile(filepath.Join(a.RepoRoot, config.RootDir, "prompts", name+".md"))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (a *App) GenerateSkill(name string) (string, error) {
	s, err := schema.Load(a.RepoRoot)
	if err != nil {
		return "", err
	}
	content, err := generator.GenerateSkill(name, s)
	if err != nil {
		return "", err
	}
	path := filepath.Join(a.RepoRoot, config.RootDir, "skills", name+".md")
	if _, err := fs.WriteIfChanged(path, content); err != nil {
		return "", err
	}
	return strings.TrimPrefix(filepath.ToSlash(path), filepath.ToSlash(a.RepoRoot)+"/"), nil
}

func (a *App) GeneratePrompt(name string) (string, error) {
	s, err := schema.Load(a.RepoRoot)
	if err != nil {
		return "", err
	}
	content, err := generator.GeneratePrompt(name, s)
	if err != nil {
		return "", err
	}
	path := filepath.Join(a.RepoRoot, config.RootDir, "prompts", name+".md")
	if _, err := fs.WriteIfChanged(path, content); err != nil {
		return "", err
	}
	return strings.TrimPrefix(filepath.ToSlash(path), filepath.ToSlash(a.RepoRoot)+"/"), nil
}
