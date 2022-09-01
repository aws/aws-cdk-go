package awscodebuild


// The types of webhook event actions.
//
// Example:
//   gitHubSource := codebuild.source.gitHub(&gitHubSourceProps{
//   	owner: jsii.String("awslabs"),
//   	repo: jsii.String("aws-cdk"),
//   	webhook: jsii.Boolean(true),
//   	 // optional, default: true if `webhookFilters` were provided, false otherwise
//   	webhookTriggersBatchBuild: jsii.Boolean(true),
//   	 // optional, default is false
//   	webhookFilters: []filterGroup{
//   		codebuild.*filterGroup.inEventOf(codebuild.eventAction_PUSH).andBranchIs(jsii.String("main")).andCommitMessageIs(jsii.String("the commit message")),
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

