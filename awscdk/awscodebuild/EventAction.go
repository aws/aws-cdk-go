package awscodebuild


// The types of webhook event actions.
//
// Example:
//   gitHubSource := codebuild.Source_GitHub(&GitHubSourceProps{
//   	Owner: jsii.String("awslabs"),
//   	Repo: jsii.String("aws-cdk"),
//   	Webhook: jsii.Boolean(true),
//   	 // optional, default: true if `webhookFilters` were provided, false otherwise
//   	WebhookTriggersBatchBuild: jsii.Boolean(true),
//   	 // optional, default is false
//   	WebhookFilters: []filterGroup{
//   		codebuild.*filterGroup_InEventOf(codebuild.EventAction_PUSH).AndBranchIs(jsii.String("main")).AndCommitMessageIs(jsii.String("the commit message")),
//   	},
//   })
//
type EventAction string

const (
	// A push (of a branch, or a tag) to the repository.
	EventAction_PUSH EventAction = "PUSH"
	// Creating a Pull Request.
	EventAction_PULL_REQUEST_CREATED EventAction = "PULL_REQUEST_CREATED"
	// Updating a Pull Request.
	EventAction_PULL_REQUEST_UPDATED EventAction = "PULL_REQUEST_UPDATED"
	// Merging a Pull Request.
	EventAction_PULL_REQUEST_MERGED EventAction = "PULL_REQUEST_MERGED"
	// Re-opening a previously closed Pull Request.
	//
	// Note that this event is only supported for GitHub and GitHubEnterprise sources.
	EventAction_PULL_REQUEST_REOPENED EventAction = "PULL_REQUEST_REOPENED"
)

