package doctor

import (
	"fmt"

	"github.com/voku/vokuknow/internal/lint"
)

type Report struct {
	Checks []string
	Issues []string
}

func Run(repoRoot string) Report {
	report := Report{Checks: []string{"structure", "schema", "generated artifacts", "policy consistency"}}
	lintResult := lint.Run(repoRoot)
	report.Issues = append(report.Issues, lintResult.Errors...)
	for _, w := range lintResult.Warnings {
		report.Issues = append(report.Issues, fmt.Sprintf("warning: %s", w))
	}
	return report
}
