package awsec2


// The available destination types for Flow Logs.
type FlowLogDestinationType string

const (
	// Send flow logs to CloudWatch Logs Group.
	FlowLogDestinationType_CLOUD_WATCH_LOGS FlowLogDestinationType = "CLOUD_WATCH_LOGS"
	// Send flow logs to S3 Bucket.
	FlowLogDestinationType_S3 FlowLogDestinationType = "S3"
	// Send flow logs to Kinesis Data Firehose.
	FlowLogDestinationType_KINESIS_DATA_FIREHOSE FlowLogDestinationType = "KINESIS_DATA_FIREHOSE"
)

