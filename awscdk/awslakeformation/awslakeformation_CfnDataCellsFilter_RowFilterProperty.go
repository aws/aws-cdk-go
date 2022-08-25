package awslakeformation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var allRowsWildcard interface{}
//
//   rowFilterProperty := &rowFilterProperty{
//   	allRowsWildcard: allRowsWildcard,
//   	filterExpression: jsii.String("filterExpression"),
//   }
//
type CfnDataCellsFilter_RowFilterProperty struct {
	// `CfnDataCellsFilter.RowFilterProperty.AllRowsWildcard`.
	AllRowsWildcard interface{} `field:"optional" json:"allRowsWildcard" yaml:"allRowsWildcard"`
	// `CfnDataCellsFilter.RowFilterProperty.FilterExpression`.
	FilterExpression *string `field:"optional" json:"filterExpression" yaml:"filterExpression"`
}

