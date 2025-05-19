package awsquicksight


// The value of a time range filter.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeRangeFilterValueProperty := &TimeRangeFilterValueProperty{
//   	Parameter: jsii.String("parameter"),
//   	RollingDate: &RollingDateConfigurationProperty{
//   		Expression: jsii.String("expression"),
//
//   		// the properties below are optional
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	StaticValue: jsii.String("staticValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefiltervalue.html
//
type CfnAnalysis_TimeRangeFilterValueProperty struct {
	// The parameter type input value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefiltervalue.html#cfn-quicksight-analysis-timerangefiltervalue-parameter
	//
	Parameter *string `field:"optional" json:"parameter" yaml:"parameter"`
	// The rolling date input value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefiltervalue.html#cfn-quicksight-analysis-timerangefiltervalue-rollingdate
	//
	RollingDate interface{} `field:"optional" json:"rollingDate" yaml:"rollingDate"`
	// The static input value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefiltervalue.html#cfn-quicksight-analysis-timerangefiltervalue-staticvalue
	//
	StaticValue *string `field:"optional" json:"staticValue" yaml:"staticValue"`
}

