package config

const RootDir = ".vokuknow"

var RequiredSchemaFiles = []string{
	"AGENTS.md",
	"entity_types.yaml",
	"relation_types.yaml",
	"claim_policies.yaml",
	"source_policies.yaml",
	"privacy_policies.yaml",
	"prompt_profiles.yaml",
}

var DefaultSkills = []string{"wiki-operator", "crystallizer", "contradiction-resolver"}

var DefaultPrompts = []string{"ingest", "query", "update-claim", "resolve-claim", "crystallize", "handoff", "lint"}
