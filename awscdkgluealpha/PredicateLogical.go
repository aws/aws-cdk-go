package awscdkgluealpha


// Experimental.
type PredicateLogical string

const (
	// All conditions must be true for the predicate to be true.
	// Experimental.
	PredicateLogical_AND PredicateLogical = "AND"
	// At least one condition must be true for the predicate to be true.
	// Experimental.
	PredicateLogical_ANY PredicateLogical = "ANY"
)

