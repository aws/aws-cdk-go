package awsquicksight


// The data options for an axis.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataDriven interface{}
//
//   axisDataOptionsProperty := &AxisDataOptionsProperty{
//   	DateAxisOptions: &DateAxisOptionsProperty{
//   		MissingDateVisibility: jsii.String("missingDateVisibility"),
//   	},
//   	NumericAxisOptions: &NumericAxisOptionsProperty{
//   		Range: &AxisDisplayRangeProperty{
//   			DataDriven: dataDriven,
//   			MinMax: &AxisDisplayMinMaxRangeProperty{
//   				Maximum: jsii.Number(123),
//   				Minimum: jsii.Number(123),
//   			},
//   		},
//   		Scale: &AxisScaleProperty{
//   			Linear: &AxisLinearScaleProperty{
//   				StepCount: jsii.Number(123),
//   				StepSize: jsii.Number(123),
//   			},
//   			Logarithmic: &AxisLogarithmicScaleProperty{
//   				Base: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-axisdataoptions.html
//
type CfnAnalysis_AxisDataOptionsProperty struct {
	// The options for an axis with a date field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-axisdataoptions.html#cfn-quicksight-analysis-axisdataoptions-dateaxisoptions
	//
	DateAxisOptions interface{} `field:"optional" json:"dateAxisOptions" yaml:"dateAxisOptions"`
	// The options for an axis with a numeric field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-axisdataoptions.html#cfn-quicksight-analysis-axisdataoptions-numericaxisoptions
	//
	NumericAxisOptions interface{} `field:"optional" json:"numericAxisOptions" yaml:"numericAxisOptions"`
}

