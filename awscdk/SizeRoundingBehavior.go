package awscdk


// Rounding behaviour when converting between units of `Size`.
// Experimental.
type SizeRoundingBehavior string

const (
	// Fail the conversion if the result is not an integer.
	// Experimental.
	SizeRoundingBehavior_FAIL SizeRoundingBehavior = "FAIL"
	// If the result is not an integer, round it to the closest integer less than the result.
	// Experimental.
	SizeRoundingBehavior_FLOOR SizeRoundingBehavior = "FLOOR"
	// Don't round.
	//
	// Return even if the result is a fraction.
	// Experimental.
	SizeRoundingBehavior_NONE SizeRoundingBehavior = "NONE"
)

