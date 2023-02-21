package awscodepipelineactions


// If and how the GitHub source action should be triggered.
type GitHubTrigger string

const (
	GitHubTrigger_NONE GitHubTrigger = "NONE"
	GitHubTrigger_POLL GitHubTrigger = "POLL"
	GitHubTrigger_WEBHOOK GitHubTrigger = "WEBHOOK"
)

