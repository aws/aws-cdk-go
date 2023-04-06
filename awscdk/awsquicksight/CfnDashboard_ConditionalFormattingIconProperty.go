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
type CfnDashboard_ConditionalFormattingIconProperty struct {
	// Determines the custom condition for an icon set.
	CustomCondition interface{} `field:"optional" json:"customCondition" yaml:"customCondition"`
	// Formatting configuration for icon set.
	IconSet interface{} `field:"optional" json:"iconSet" yaml:"iconSet"`
}

