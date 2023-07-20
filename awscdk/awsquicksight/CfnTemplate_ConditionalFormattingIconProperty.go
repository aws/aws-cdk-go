package awsquicksight


// The formatting configuration for the icon.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionalFormattingIconProperty := &ConditionalFormattingIconProperty{
//   	CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   		Expression: jsii.String("expression"),
//   		IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   			Icon: jsii.String("icon"),
//   			UnicodeIcon: jsii.String("unicodeIcon"),
//   		},
//
//   		// the properties below are optional
//   		Color: jsii.String("color"),
//   		DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   			IconDisplayOption: jsii.String("iconDisplayOption"),
//   		},
//   	},
//   	IconSet: &ConditionalFormattingIconSetProperty{
//   		Expression: jsii.String("expression"),
//
//   		// the properties below are optional
//   		IconSetType: jsii.String("iconSetType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattingicon.html
//
type CfnTemplate_ConditionalFormattingIconProperty struct {
	// Determines the custom condition for an icon set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattingicon.html#cfn-quicksight-template-conditionalformattingicon-customcondition
	//
	CustomCondition interface{} `field:"optional" json:"customCondition" yaml:"customCondition"`
	// Formatting configuration for icon set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattingicon.html#cfn-quicksight-template-conditionalformattingicon-iconset
	//
	IconSet interface{} `field:"optional" json:"iconSet" yaml:"iconSet"`
}

