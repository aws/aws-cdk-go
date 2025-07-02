package awslambda


// The type of destination.
type DestinationType string

const (
	// Failure.
	DestinationType_FAILURE DestinationType = "FAILURE"
	// Success.
	DestinationType_SUCCESS DestinationType = "SUCCESS"
)

