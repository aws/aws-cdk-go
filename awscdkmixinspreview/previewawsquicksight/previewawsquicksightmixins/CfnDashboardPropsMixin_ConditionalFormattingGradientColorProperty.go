package previewawsquicksightmixins


// Formatting configuration for gradient color.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   conditionalFormattingGradientColorProperty := &ConditionalFormattingGradientColorProperty{
//   	Color: &GradientColorProperty{
//   		Stops: []interface{}{
//   			&GradientStopProperty{
//   				Color: jsii.String("color"),
//   				DataValue: jsii.Number(123),
//   				GradientOffset: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Expression: jsii.String("expression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattinggradientcolor.html
//
type CfnDashboardPropsMixin_ConditionalFormattingGradientColorProperty struct {
	// Determines the color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattinggradientcolor.html#cfn-quicksight-dashboard-conditionalformattinggradientcolor-color
	//
	Color interface{} `field:"optional" json:"color" yaml:"color"`
	// The expression that determines the formatting configuration for gradient color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattinggradientcolor.html#cfn-quicksight-dashboard-conditionalformattinggradientcolor-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

