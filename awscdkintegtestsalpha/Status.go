// CDK Integration Testing Constructs
package awscdkintegtestsalpha


// The status of the assertion.
// Experimental.
type Status string

const (
	// The assertion passed.
	// Experimental.
	Status_PASS Status = "PASS"
	// The assertion failed.
	// Experimental.
	Status_FAIL Status = "FAIL"
)

