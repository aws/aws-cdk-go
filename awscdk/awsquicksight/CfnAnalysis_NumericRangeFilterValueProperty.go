package awsquicksight


// The value input pf the numeric range filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numericRangeFilterValueProperty := &NumericRangeFilterValueProperty{
//   	Parameter: jsii.String("parameter"),
//   	StaticValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericrangefiltervalue.html
//
type CfnAnalysis_NumericRangeFilterValueProperty struct {
	// The parameter that is used in the numeric range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericrangefiltervalue.html#cfn-quicksight-analysis-numericrangefiltervalue-parameter
	//
	Parameter *string `field:"optional" json:"parameter" yaml:"parameter"`
	// The static value of the numeric range filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericrangefiltervalue.html#cfn-quicksight-analysis-numericrangefiltervalue-staticvalue
	//
	StaticValue *float64 `field:"optional" json:"staticValue" yaml:"staticValue"`
}

