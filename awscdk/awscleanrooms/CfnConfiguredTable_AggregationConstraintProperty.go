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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregationconstraint.html
//
type CfnConfiguredTable_AggregationConstraintProperty struct {
	// Column in aggregation constraint for which there must be a minimum number of distinct values in an output row for it to be in the query output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregationconstraint.html#cfn-cleanrooms-configuredtable-aggregationconstraint-columnname
	//
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The minimum number of distinct values that an output row must be an aggregation of.
	//
	// Minimum threshold of distinct values for a specified column that must exist in an output row for it to be in the query output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregationconstraint.html#cfn-cleanrooms-configuredtable-aggregationconstraint-minimum
	//
	Minimum *float64 `field:"required" json:"minimum" yaml:"minimum"`
	// The type of aggregation the constraint allows.
	//
	// The only valid value is currently `COUNT_DISTINCT`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregationconstraint.html#cfn-cleanrooms-configuredtable-aggregationconstraint-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

