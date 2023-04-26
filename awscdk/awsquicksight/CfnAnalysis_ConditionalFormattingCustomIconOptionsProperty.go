package awsquicksight


// Custom icon options for an icon set.
//
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
type CfnAnalysis_ConditionalFormattingCustomIconOptionsProperty struct {
	// Determines the type of icon.
	Icon *string `field:"optional" json:"icon" yaml:"icon"`
	// Determines the Unicode icon type.
	UnicodeIcon *string `field:"optional" json:"unicodeIcon" yaml:"unicodeIcon"`
}

