package awsquicksight


// Determines the custom condition for an icon set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionalFormattingCustomIconConditionProperty := &ConditionalFormattingCustomIconConditionProperty{
//   	Expression: jsii.String("expression"),
//   	IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   		Icon: jsii.String("icon"),
//   		UnicodeIcon: jsii.String("unicodeIcon"),
//   	},
//
//   	// the properties below are optional
//   	Color: jsii.String("color"),
//   	DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   		IconDisplayOption: jsii.String("iconDisplayOption"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattingcustomiconcondition.html
//
type CfnTemplate_ConditionalFormattingCustomIconConditionProperty struct {
	// The expression that determines the condition of the icon set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattingcustomiconcondition.html#cfn-quicksight-template-conditionalformattingcustomiconcondition-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Custom icon options for an icon set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattingcustomiconcondition.html#cfn-quicksight-template-conditionalformattingcustomiconcondition-iconoptions
	//
	IconOptions interface{} `field:"required" json:"iconOptions" yaml:"iconOptions"`
	// Determines the color of the icon.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattingcustomiconcondition.html#cfn-quicksight-template-conditionalformattingcustomiconcondition-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Determines the icon display configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattingcustomiconcondition.html#cfn-quicksight-template-conditionalformattingcustomiconcondition-displayconfiguration
	//
	DisplayConfiguration interface{} `field:"optional" json:"displayConfiguration" yaml:"displayConfiguration"`
}

