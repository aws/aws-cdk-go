package awsquicksight


// The numeric equality type drill down filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numericEqualityDrillDownFilterProperty := &NumericEqualityDrillDownFilterProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericequalitydrilldownfilter.html
//
type CfnAnalysis_NumericEqualityDrillDownFilterProperty struct {
	// The column that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericequalitydrilldownfilter.html#cfn-quicksight-analysis-numericequalitydrilldownfilter-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The value of the double input numeric drill down filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericequalitydrilldownfilter.html#cfn-quicksight-analysis-numericequalitydrilldownfilter-value
	//
	// Default: - 0.
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

