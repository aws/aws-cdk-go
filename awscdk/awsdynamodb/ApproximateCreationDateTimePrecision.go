package awsdynamodb


// The precision associated with the DynamoDB write timestamps that will be replicated to Kinesis.
//
// The default setting for record timestamp precision is microseconds. You can change this setting at any time.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-kinesisstreamspecification.html#aws-properties-dynamodb-table-kinesisstreamspecification-properties
//
type ApproximateCreationDateTimePrecision string

const (
	// Millisecond precision.
	ApproximateCreationDateTimePrecision_MILLISECOND ApproximateCreationDateTimePrecision = "MILLISECOND"
	// Microsecond precision.
	ApproximateCreationDateTimePrecision_MICROSECOND ApproximateCreationDateTimePrecision = "MICROSECOND"
)

