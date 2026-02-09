package awscdkgluealpha


// Partition projection type.
//
// Determines how Athena projects partition values.
// See: https://docs.aws.amazon.com/athena/latest/ug/partition-projection-supported-types.html
//
// Experimental.
type PartitionProjectionType string

const (
	// Project partition values as integers within a range.
	// Experimental.
	PartitionProjectionType_INTEGER PartitionProjectionType = "INTEGER"
	// Project partition values as dates within a range.
	// Experimental.
	PartitionProjectionType_DATE PartitionProjectionType = "DATE"
	// Project partition values from an explicit list of values.
	// Experimental.
	PartitionProjectionType_ENUM PartitionProjectionType = "ENUM"
	// Project partition values that are injected at query time.
	// Experimental.
	PartitionProjectionType_INJECTED PartitionProjectionType = "INJECTED"
)

