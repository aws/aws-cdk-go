package awsquicksight


// The formatting configuration for the color.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionalFormattingColorProperty := &ConditionalFormattingColorProperty{
//   	Gradient: &ConditionalFormattingGradientColorProperty{
//   		Color: &GradientColorProperty{
//   			Stops: []interface{}{
//   				&GradientStopProperty{
//   					GradientOffset: jsii.Number(123),
//
//   					// the properties below are optional
//   					Color: jsii.String("color"),
//   					DataValue: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Expression: jsii.String("expression"),
//   	},
//   	Solid: &ConditionalFormattingSolidColorProperty{
//   		Expression: jsii.String("expression"),
//
//   		// the properties below are optional
//   		Color: jsii.String("color"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingcolor.html
//
type CfnAnalysis_ConditionalFormattingColorProperty struct {
	// Formatting configuration for gradient color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingcolor.html#cfn-quicksight-analysis-conditionalformattingcolor-gradient
	//
	Gradient interface{} `field:"optional" json:"gradient" yaml:"gradient"`
	// Formatting configuration for solid color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingcolor.html#cfn-quicksight-analysis-conditionalformattingcolor-solid
	//
	Solid interface{} `field:"optional" json:"solid" yaml:"solid"`
}

