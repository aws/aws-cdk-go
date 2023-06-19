package awslambda


// The type of destination.
// Experimental.
type DestinationType string

const (
	// Failure.
	// Experimental.
	DestinationType_FAILURE DestinationType = "FAILURE"
	// Success.
	// Experimental.
	DestinationType_SUCCESS DestinationType = "SUCCESS"
)

