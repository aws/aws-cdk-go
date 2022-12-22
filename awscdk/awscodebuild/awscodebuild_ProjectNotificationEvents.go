package awscodebuild


// The list of event types for AWS Codebuild.
// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-buildproject
//
// Experimental.
type ProjectNotificationEvents string

const (
	// Trigger notification when project build state failed.
	// Experimental.
	ProjectNotificationEvents_BUILD_FAILED ProjectNotificationEvents = "BUILD_FAILED"
	// Trigger notification when project build state succeeded.
	// Experimental.
	ProjectNotificationEvents_BUILD_SUCCEEDED ProjectNotificationEvents = "BUILD_SUCCEEDED"
	// Trigger notification when project build state in progress.
	// Experimental.
	ProjectNotificationEvents_BUILD_IN_PROGRESS ProjectNotificationEvents = "BUILD_IN_PROGRESS"
	// Trigger notification when project build state stopped.
	// Experimental.
	ProjectNotificationEvents_BUILD_STOPPED ProjectNotificationEvents = "BUILD_STOPPED"
	// Trigger notification when project build phase failure.
	// Experimental.
	ProjectNotificationEvents_BUILD_PHASE_FAILED ProjectNotificationEvents = "BUILD_PHASE_FAILED"
	// Trigger notification when project build phase success.
	// Experimental.
	ProjectNotificationEvents_BUILD_PHASE_SUCCEEDED ProjectNotificationEvents = "BUILD_PHASE_SUCCEEDED"
)

