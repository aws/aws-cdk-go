package awslakeformation


// A PartiQL predicate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var allRowsWildcard interface{}
//
//   rowFilterProperty := &RowFilterProperty{
//   	AllRowsWildcard: allRowsWildcard,
//   	FilterExpression: jsii.String("filterExpression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-datacellsfilter-rowfilter.html
//
type CfnDataCellsFilter_RowFilterProperty struct {
	// A wildcard for all rows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-datacellsfilter-rowfilter.html#cfn-lakeformation-datacellsfilter-rowfilter-allrowswildcard
	//
	AllRowsWildcard interface{} `field:"optional" json:"allRowsWildcard" yaml:"allRowsWildcard"`
	// A filter expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-datacellsfilter-rowfilter.html#cfn-lakeformation-datacellsfilter-rowfilter-filterexpression
	//
	FilterExpression *string `field:"optional" json:"filterExpression" yaml:"filterExpression"`
}

