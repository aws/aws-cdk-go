package awsglue


// Job states emitted by Glue to CloudWatch Events.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types for more information.
//
// Experimental.
type JobState string

const (
	// State indicating job run succeeded.
	// Experimental.
	JobState_SUCCEEDED JobState = "SUCCEEDED"
	// State indicating job run failed.
	// Experimental.
	JobState_FAILED JobState = "FAILED"
	// State indicating job run timed out.
	// Experimental.
	JobState_TIMEOUT JobState = "TIMEOUT"
	// State indicating job is starting.
	// Experimental.
	JobState_STARTING JobState = "STARTING"
	// State indicating job is running.
	// Experimental.
	JobState_RUNNING JobState = "RUNNING"
	// State indicating job is stopping.
	// Experimental.
	JobState_STOPPING JobState = "STOPPING"
	// State indicating job stopped.
	// Experimental.
	JobState_STOPPED JobState = "STOPPED"
)

