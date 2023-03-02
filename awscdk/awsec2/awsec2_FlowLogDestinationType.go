package awsec2


// The available destination types for Flow Logs.
// Experimental.
type FlowLogDestinationType string

const (
	// Send flow logs to CloudWatch Logs Group.
	// Experimental.
	FlowLogDestinationType_CLOUD_WATCH_LOGS FlowLogDestinationType = "CLOUD_WATCH_LOGS"
	// Send flow logs to S3 Bucket.
	// Experimental.
	FlowLogDestinationType_S3 FlowLogDestinationType = "S3"
)

