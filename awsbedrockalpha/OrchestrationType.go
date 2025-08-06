package awsbedrockalpha


// Enum for orchestration types available for agents.
// Experimental.
type OrchestrationType string

const (
	// Default orchestration by the agent.
	// Experimental.
	OrchestrationType_DEFAULT OrchestrationType = "DEFAULT"
	// Custom orchestration using Lambda.
	// Experimental.
	OrchestrationType_CUSTOM_ORCHESTRATION OrchestrationType = "CUSTOM_ORCHESTRATION"
)

