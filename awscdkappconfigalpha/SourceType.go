package awscdkappconfigalpha


// Defines the source type for event destinations.
// Experimental.
type SourceType string

const (
	// Experimental.
	SourceType_LAMBDA SourceType = "LAMBDA"
	// Experimental.
	SourceType_SQS SourceType = "SQS"
	// Experimental.
	SourceType_SNS SourceType = "SNS"
	// Experimental.
	SourceType_EVENTS SourceType = "EVENTS"
)

