package schema

type TypeDef struct {
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
}

type EntityTypesFile struct {
	Entities []TypeDef `yaml:"entities"`
}

type RelationTypesFile struct {
	Relations []TypeDef `yaml:"relations"`
}

type ClaimPolicies struct {
	RequiredProvenanceFields []string `yaml:"required_provenance_fields" json:"required_provenance_fields"`
	ContradictionStrategy    string   `yaml:"contradiction_strategy" json:"contradiction_strategy"`
	UpdateRule               string   `yaml:"update_rule" json:"update_rule"`
	CrystallizationRequired  bool     `yaml:"crystallization_required" json:"crystallization_required"`
}

type Authority struct {
	Name   string `yaml:"name" json:"name"`
	Weight int    `yaml:"weight" json:"weight"`
}

type SourcePolicies struct {
	Authorities []Authority `yaml:"authorities" json:"authorities"`
}

type PrivacyPolicies struct {
	RedactionPatterns []string `yaml:"redaction_patterns" json:"redaction_patterns"`
	PrivatePaths      []string `yaml:"private_paths" json:"private_paths"`
	PromotionRule     string   `yaml:"promotion_rule" json:"promotion_rule"`
}

type PromptProfile struct {
	GoalTemplate string `yaml:"goal_template" json:"goal_template"`
	DoneTemplate string `yaml:"done_template" json:"done_template"`
}

type PromptProfilesFile struct {
	Profiles map[string]PromptProfile `yaml:"profiles"`
}

type Schema struct {
	AgentsMarkdown string                   `json:"agents_markdown"`
	EntityTypes    []TypeDef                `json:"entity_types"`
	RelationTypes  []TypeDef                `json:"relation_types"`
	ClaimPolicies  ClaimPolicies            `json:"claim_policies"`
	SourcePolicies SourcePolicies           `json:"source_policies"`
	Privacy        PrivacyPolicies          `json:"privacy_policies"`
	PromptProfiles map[string]PromptProfile `json:"prompt_profiles"`
}
