package cli

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/voku/vokuknow/internal/app"
)

func Execute(args []string, out, errOut io.Writer) error {
	var repoRoot string
	root := &cobra.Command{
		Use:   "vokuknow",
		Short: "Generate deterministic repo-local knowledge contracts",
	}
	root.SetOut(out)
	root.SetErr(errOut)
	root.SetArgs(args)
	root.PersistentFlags().StringVar(&repoRoot, "repo", ".", "repository root path")

	root.AddCommand(initCmd(&repoRoot, out))
	root.AddCommand(buildCmd(&repoRoot, out))
	root.AddCommand(lintCmd(&repoRoot, out))
	root.AddCommand(doctorCmd(&repoRoot, out))
	root.AddCommand(showCmd(&repoRoot, out))
	root.AddCommand(skillCmd(&repoRoot, out))
	root.AddCommand(promptCmd(&repoRoot, out))

	return root.Execute()
}

func initCmd(repoRoot *string, out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize repo-local vokuknow structure and starter schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			a := app.New(*repoRoot)
			files, err := a.Init()
			if err != nil {
				return err
			}
			fmt.Fprintf(out, "initialized vokuknow; wrote %d starter files\n", len(files))
			return nil
		},
	}
}

func buildCmd(repoRoot *string, out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "Build deterministic skills, prompts, and manifests from schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			a := app.New(*repoRoot)
			changed, err := a.Build()
			if err != nil {
				return err
			}
			fmt.Fprintf(out, "build complete; changed %d artifact(s)\n", len(changed))
			return nil
		},
	}
}

func lintCmd(repoRoot *string, out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "lint",
		Short: "Lint schema and generated artifacts",
		RunE: func(cmd *cobra.Command, args []string) error {
			a := app.New(*repoRoot)
			if err := a.Lint(); err != nil {
				return err
			}
			fmt.Fprintln(out, "lint passed")
			return nil
		},
	}
}

func doctorCmd(repoRoot *string, out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "doctor",
		Short: "Run broad health checks",
		RunE: func(cmd *cobra.Command, args []string) error {
			a := app.New(*repoRoot)
			report, err := a.Doctor()
			for _, check := range report.Checks {
				fmt.Fprintf(out, "check: %s\n", check)
			}
			for _, issue := range report.Issues {
				fmt.Fprintf(out, "issue: %s\n", issue)
			}
			return err
		},
	}
}

func showCmd(repoRoot *string, out io.Writer) *cobra.Command {
	show := &cobra.Command{Use: "show", Short: "Show generated artifacts"}
	show.AddCommand(&cobra.Command{
		Use:   "skill <name>",
		Args:  cobra.ExactArgs(1),
		Short: "Show generated skill content",
		RunE: func(cmd *cobra.Command, args []string) error {
			a := app.New(*repoRoot)
			content, err := a.ShowSkill(args[0])
			if err != nil {
				return err
			}
			fmt.Fprint(out, content)
			return nil
		},
	})
	show.AddCommand(&cobra.Command{
		Use:   "prompt <name>",
		Args:  cobra.ExactArgs(1),
		Short: "Show generated prompt content",
		RunE: func(cmd *cobra.Command, args []string) error {
			a := app.New(*repoRoot)
			content, err := a.ShowPrompt(args[0])
			if err != nil {
				return err
			}
			fmt.Fprint(out, content)
			return nil
		},
	})
	return show
}

func skillCmd(repoRoot *string, out io.Writer) *cobra.Command {
	skill := &cobra.Command{Use: "skill", Short: "Skill operations"}
	skill.AddCommand(&cobra.Command{
		Use:   "generate <name>",
		Args:  cobra.ExactArgs(1),
		Short: "Generate one skill file",
		RunE: func(cmd *cobra.Command, args []string) error {
			a := app.New(*repoRoot)
			path, err := a.GenerateSkill(args[0])
			if err != nil {
				return err
			}
			fmt.Fprintf(out, "generated %s\n", path)
			return nil
		},
	})
	return skill
}

func promptCmd(repoRoot *string, out io.Writer) *cobra.Command {
	prompt := &cobra.Command{Use: "prompt", Short: "Prompt operations"}
	prompt.AddCommand(&cobra.Command{
		Use:   "generate <name>",
		Args:  cobra.ExactArgs(1),
		Short: "Generate one prompt file",
		RunE: func(cmd *cobra.Command, args []string) error {
			a := app.New(*repoRoot)
			path, err := a.GeneratePrompt(args[0])
			if err != nil {
				return err
			}
			fmt.Fprintf(out, "generated %s\n", path)
			return nil
		},
	})
	return prompt
}
