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
type CfnDataCellsFilter_RowFilterProperty struct {
	// A wildcard for all rows.
	AllRowsWildcard interface{} `field:"optional" json:"allRowsWildcard" yaml:"allRowsWildcard"`
	// A filter expression.
	FilterExpression *string `field:"optional" json:"filterExpression" yaml:"filterExpression"`
}

