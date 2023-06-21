package awscleanrooms


// Constraint on query output removing output rows that do not meet a minimum number of distinct values of a specified column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aggregationConstraintProperty := &AggregationConstraintProperty{
//   	ColumnName: jsii.String("columnName"),
//   	Minimum: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
type CfnConfiguredTable_AggregationConstraintProperty struct {
	// Column in aggregation constraint for which there must be a minimum number of distinct values in an output row for it to be in the query output.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The minimum number of distinct values that an output row must be an aggregation of.
	//
	// Minimum threshold of distinct values for a specified column that must exist in an output row for it to be in the query output.
	Minimum *float64 `field:"required" json:"minimum" yaml:"minimum"`
	// The type of aggregation the constraint allows.
	//
	// The only valid value is currently `COUNT_DISTINCT`.
	Type *string `field:"required" json:"type" yaml:"type"`
}

