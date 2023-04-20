package awsquicksight


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
type CfnDashboard_ConditionalFormattingCustomIconConditionProperty struct {
	// `CfnDashboard.ConditionalFormattingCustomIconConditionProperty.Expression`.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// `CfnDashboard.ConditionalFormattingCustomIconConditionProperty.IconOptions`.
	IconOptions interface{} `field:"required" json:"iconOptions" yaml:"iconOptions"`
	// `CfnDashboard.ConditionalFormattingCustomIconConditionProperty.Color`.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// `CfnDashboard.ConditionalFormattingCustomIconConditionProperty.DisplayConfiguration`.
	DisplayConfiguration interface{} `field:"optional" json:"displayConfiguration" yaml:"displayConfiguration"`
}

