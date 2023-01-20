package awsdynamodb


// The set of attributes that are projected into the index.
// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Projection.html
//
type ProjectionType string

const (
	// Only the index and primary keys are projected into the index.
	ProjectionType_KEYS_ONLY ProjectionType = "KEYS_ONLY"
	// Only the specified table attributes are projected into the index.
	//
	// The list of projected attributes is in `nonKeyAttributes`.
	ProjectionType_INCLUDE ProjectionType = "INCLUDE"
	// All of the table attributes are projected into the index.
	ProjectionType_ALL ProjectionType = "ALL"
)

