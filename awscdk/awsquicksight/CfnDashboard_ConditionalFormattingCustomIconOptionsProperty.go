package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionalFormattingCustomIconOptionsProperty := &ConditionalFormattingCustomIconOptionsProperty{
//   	Icon: jsii.String("icon"),
//   	UnicodeIcon: jsii.String("unicodeIcon"),
//   }
//
type CfnDashboard_ConditionalFormattingCustomIconOptionsProperty struct {
	// `CfnDashboard.ConditionalFormattingCustomIconOptionsProperty.Icon`.
	Icon *string `field:"optional" json:"icon" yaml:"icon"`
	// `CfnDashboard.ConditionalFormattingCustomIconOptionsProperty.UnicodeIcon`.
	UnicodeIcon *string `field:"optional" json:"unicodeIcon" yaml:"unicodeIcon"`
}

