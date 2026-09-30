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
//
//   		// the properties below are optional
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		TopicIdentifier: jsii.String("topicIdentifier"),
//   	},
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-numericequalitydrilldownfilter.html
//
type CfnTemplate_NumericEqualityDrillDownFilterProperty struct {
	// The column that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-numericequalitydrilldownfilter.html#cfn-quicksight-template-numericequalitydrilldownfilter-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The value of the double input numeric drill down filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-numericequalitydrilldownfilter.html#cfn-quicksight-template-numericequalitydrilldownfilter-value
	//
	// Default: - 0.
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

