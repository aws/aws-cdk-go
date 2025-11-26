package previewawsquicksightmixins


// Formatting configuration for solid color.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   conditionalFormattingSolidColorProperty := &ConditionalFormattingSolidColorProperty{
//   	Color: jsii.String("color"),
//   	Expression: jsii.String("expression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattingsolidcolor.html
//
type CfnDashboardPropsMixin_ConditionalFormattingSolidColorProperty struct {
	// Determines the color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattingsolidcolor.html#cfn-quicksight-dashboard-conditionalformattingsolidcolor-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The expression that determines the formatting configuration for solid color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattingsolidcolor.html#cfn-quicksight-dashboard-conditionalformattingsolidcolor-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

