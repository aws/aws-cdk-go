package awscleanrooms


// Column in configured table that can be used in aggregate function in query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aggregateColumnProperty := &AggregateColumnProperty{
//   	ColumnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	Function: jsii.String("function"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregatecolumn.html
//
type CfnConfiguredTable_AggregateColumnProperty struct {
	// Column names in configured table of aggregate columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregatecolumn.html#cfn-cleanrooms-configuredtable-aggregatecolumn-columnnames
	//
	ColumnNames *[]*string `field:"required" json:"columnNames" yaml:"columnNames"`
	// Aggregation function that can be applied to aggregate column in query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregatecolumn.html#cfn-cleanrooms-configuredtable-aggregatecolumn-function
	//
	Function *string `field:"required" json:"function" yaml:"function"`
}

