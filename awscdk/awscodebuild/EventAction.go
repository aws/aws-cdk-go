package awscodebuild


// The types of webhook event actions.
//
// Example:
//   gitHubSource := codebuild.Source_GitHub(&GitHubSourceProps{
//   	Owner: jsii.String("awslabs"),
//   	Repo: jsii.String("aws-cdk"),
//   	 // optional, default: undefined if unspecified will create organization webhook
//   	Webhook: jsii.Boolean(true),
//   	 // optional, default: true if `webhookFilters` were provided, false otherwise
//   	WebhookTriggersBatchBuild: jsii.Boolean(true),
//   	 // optional, default is false
//   	WebhookFilters: []filterGroup{
//   		codebuild.*filterGroup_InEventOf(codebuild.EventAction_PUSH).AndBranchIs(jsii.String("main")).AndCommitMessageIs(jsii.String("the commit message")),
//   		codebuild.*filterGroup_*InEventOf(codebuild.EventAction_RELEASED).*AndBranchIs(jsii.String("main")),
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
	// Closing a Pull Request.
	EventAction_PULL_REQUEST_CLOSED EventAction = "PULL_REQUEST_CLOSED"
	// Merging a Pull Request.
	EventAction_PULL_REQUEST_MERGED EventAction = "PULL_REQUEST_MERGED"
	// Re-opening a previously closed Pull Request.
	//
	// Note that this event is only supported for GitHub and GitHubEnterprise sources.
	EventAction_PULL_REQUEST_REOPENED EventAction = "PULL_REQUEST_REOPENED"
	// A release is created in the repository.
	//
	// Works with GitHub only.
	EventAction_RELEASED EventAction = "RELEASED"
	// A prerelease is created in the repository.
	//
	// Works with GitHub only.
	EventAction_PRERELEASED EventAction = "PRERELEASED"
	// A workflow job is queued in the repository.
	//
	// Works with GitHub only.
	EventAction_WORKFLOW_JOB_QUEUED EventAction = "WORKFLOW_JOB_QUEUED"
)

