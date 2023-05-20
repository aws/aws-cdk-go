package awscdkintegtestsalpha


// The type of assertion to perform.
// Experimental.
type AssertionType string

const (
	// Assert that two values are equal.
	// Experimental.
	AssertionType_EQUALS AssertionType = "EQUALS"
	// The keys and their values must be present in the target but the target can be a superset.
	// Experimental.
	AssertionType_OBJECT_LIKE AssertionType = "OBJECT_LIKE"
	// Matches the specified pattern with the array The set of elements must be in the same order as would be found.
	// Experimental.
	AssertionType_ARRAY_WITH AssertionType = "ARRAY_WITH"
)

