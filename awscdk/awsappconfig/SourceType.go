package awsappconfig


// Defines the source type for event destinations.
type SourceType string

const (
	SourceType_LAMBDA SourceType = "LAMBDA"
	SourceType_SQS SourceType = "SQS"
	SourceType_SNS SourceType = "SNS"
	SourceType_EVENTS SourceType = "EVENTS"
)

