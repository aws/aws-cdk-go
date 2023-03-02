package awscodepipeline


// The list of event types for AWS Codepipeline Pipeline.
// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
//
// Experimental.
type PipelineNotificationEvents string

const (
	// Trigger notification when pipeline execution failed.
	// Experimental.
	PipelineNotificationEvents_PIPELINE_EXECUTION_FAILED PipelineNotificationEvents = "PIPELINE_EXECUTION_FAILED"
	// Trigger notification when pipeline execution canceled.
	// Experimental.
	PipelineNotificationEvents_PIPELINE_EXECUTION_CANCELED PipelineNotificationEvents = "PIPELINE_EXECUTION_CANCELED"
	// Trigger notification when pipeline execution started.
	// Experimental.
	PipelineNotificationEvents_PIPELINE_EXECUTION_STARTED PipelineNotificationEvents = "PIPELINE_EXECUTION_STARTED"
	// Trigger notification when pipeline execution resumed.
	// Experimental.
	PipelineNotificationEvents_PIPELINE_EXECUTION_RESUMED PipelineNotificationEvents = "PIPELINE_EXECUTION_RESUMED"
	// Trigger notification when pipeline execution succeeded.
	// Experimental.
	PipelineNotificationEvents_PIPELINE_EXECUTION_SUCCEEDED PipelineNotificationEvents = "PIPELINE_EXECUTION_SUCCEEDED"
	// Trigger notification when pipeline execution superseded.
	// Experimental.
	PipelineNotificationEvents_PIPELINE_EXECUTION_SUPERSEDED PipelineNotificationEvents = "PIPELINE_EXECUTION_SUPERSEDED"
	// Trigger notification when pipeline stage execution started.
	// Experimental.
	PipelineNotificationEvents_STAGE_EXECUTION_STARTED PipelineNotificationEvents = "STAGE_EXECUTION_STARTED"
	// Trigger notification when pipeline stage execution succeeded.
	// Experimental.
	PipelineNotificationEvents_STAGE_EXECUTION_SUCCEEDED PipelineNotificationEvents = "STAGE_EXECUTION_SUCCEEDED"
	// Trigger notification when pipeline stage execution resumed.
	// Experimental.
	PipelineNotificationEvents_STAGE_EXECUTION_RESUMED PipelineNotificationEvents = "STAGE_EXECUTION_RESUMED"
	// Trigger notification when pipeline stage execution canceled.
	// Experimental.
	PipelineNotificationEvents_STAGE_EXECUTION_CANCELED PipelineNotificationEvents = "STAGE_EXECUTION_CANCELED"
	// Trigger notification when pipeline stage execution failed.
	// Experimental.
	PipelineNotificationEvents_STAGE_EXECUTION_FAILED PipelineNotificationEvents = "STAGE_EXECUTION_FAILED"
	// Trigger notification when pipeline action execution succeeded.
	// Experimental.
	PipelineNotificationEvents_ACTION_EXECUTION_SUCCEEDED PipelineNotificationEvents = "ACTION_EXECUTION_SUCCEEDED"
	// Trigger notification when pipeline action execution failed.
	// Experimental.
	PipelineNotificationEvents_ACTION_EXECUTION_FAILED PipelineNotificationEvents = "ACTION_EXECUTION_FAILED"
	// Trigger notification when pipeline action execution canceled.
	// Experimental.
	PipelineNotificationEvents_ACTION_EXECUTION_CANCELED PipelineNotificationEvents = "ACTION_EXECUTION_CANCELED"
	// Trigger notification when pipeline action execution started.
	// Experimental.
	PipelineNotificationEvents_ACTION_EXECUTION_STARTED PipelineNotificationEvents = "ACTION_EXECUTION_STARTED"
	// Trigger notification when pipeline manual approval failed.
	// Experimental.
	PipelineNotificationEvents_MANUAL_APPROVAL_FAILED PipelineNotificationEvents = "MANUAL_APPROVAL_FAILED"
	// Trigger notification when pipeline manual approval needed.
	// Experimental.
	PipelineNotificationEvents_MANUAL_APPROVAL_NEEDED PipelineNotificationEvents = "MANUAL_APPROVAL_NEEDED"
	// Trigger notification when pipeline manual approval succeeded.
	// Experimental.
	PipelineNotificationEvents_MANUAL_APPROVAL_SUCCEEDED PipelineNotificationEvents = "MANUAL_APPROVAL_SUCCEEDED"
)

