package awsquicksight


// Formatting configuration for gradient color.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionalFormattingGradientColorProperty := &ConditionalFormattingGradientColorProperty{
//   	Color: &GradientColorProperty{
//   		Stops: []interface{}{
//   			&GradientStopProperty{
//   				GradientOffset: jsii.Number(123),
//
//   				// the properties below are optional
//   				Color: jsii.String("color"),
//   				DataValue: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Expression: jsii.String("expression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattinggradientcolor.html
//
type CfnDashboard_ConditionalFormattingGradientColorProperty struct {
	// Determines the color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattinggradientcolor.html#cfn-quicksight-dashboard-conditionalformattinggradientcolor-color
	//
	Color interface{} `field:"required" json:"color" yaml:"color"`
	// The expression that determines the formatting configuration for gradient color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattinggradientcolor.html#cfn-quicksight-dashboard-conditionalformattinggradientcolor-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

