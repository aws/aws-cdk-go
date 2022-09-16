package awscodebuild


// The list of event types for AWS Codebuild.
// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-buildproject
//
type ProjectNotificationEvents string

const (
	// Trigger notification when project build state failed.
	ProjectNotificationEvents_BUILD_FAILED ProjectNotificationEvents = "BUILD_FAILED"
	// Trigger notification when project build state succeeded.
	ProjectNotificationEvents_BUILD_SUCCEEDED ProjectNotificationEvents = "BUILD_SUCCEEDED"
	// Trigger notification when project build state in progress.
	ProjectNotificationEvents_BUILD_IN_PROGRESS ProjectNotificationEvents = "BUILD_IN_PROGRESS"
	// Trigger notification when project build state stopped.
	ProjectNotificationEvents_BUILD_STOPPED ProjectNotificationEvents = "BUILD_STOPPED"
	// Trigger notification when project build phase failure.
	ProjectNotificationEvents_BUILD_PHASE_FAILED ProjectNotificationEvents = "BUILD_PHASE_FAILED"
	// Trigger notification when project build phase success.
	ProjectNotificationEvents_BUILD_PHASE_SUCCEEDED ProjectNotificationEvents = "BUILD_PHASE_SUCCEEDED"
)

