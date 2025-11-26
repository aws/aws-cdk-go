package previewawsquicksightmixins


// Determines the custom condition for an icon set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   conditionalFormattingCustomIconConditionProperty := &ConditionalFormattingCustomIconConditionProperty{
//   	Color: jsii.String("color"),
//   	DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   		IconDisplayOption: jsii.String("iconDisplayOption"),
//   	},
//   	Expression: jsii.String("expression"),
//   	IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   		Icon: jsii.String("icon"),
//   		UnicodeIcon: jsii.String("unicodeIcon"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingcustomiconcondition.html
//
type CfnAnalysisPropsMixin_ConditionalFormattingCustomIconConditionProperty struct {
	// Determines the color of the icon.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingcustomiconcondition.html#cfn-quicksight-analysis-conditionalformattingcustomiconcondition-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Determines the icon display configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingcustomiconcondition.html#cfn-quicksight-analysis-conditionalformattingcustomiconcondition-displayconfiguration
	//
	DisplayConfiguration interface{} `field:"optional" json:"displayConfiguration" yaml:"displayConfiguration"`
	// The expression that determines the condition of the icon set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingcustomiconcondition.html#cfn-quicksight-analysis-conditionalformattingcustomiconcondition-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Custom icon options for an icon set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingcustomiconcondition.html#cfn-quicksight-analysis-conditionalformattingcustomiconcondition-iconoptions
	//
	IconOptions interface{} `field:"optional" json:"iconOptions" yaml:"iconOptions"`
}

