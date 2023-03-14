package awscodepipelineactions


// If and how the GitHub source action should be triggered.
// Experimental.
type GitHubTrigger string

const (
	// Experimental.
	GitHubTrigger_NONE GitHubTrigger = "NONE"
	// Experimental.
	GitHubTrigger_POLL GitHubTrigger = "POLL"
	// Experimental.
	GitHubTrigger_WEBHOOK GitHubTrigger = "WEBHOOK"
)

