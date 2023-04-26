package awsquicksight


// Determines the icon display configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionalFormattingIconDisplayConfigurationProperty := &ConditionalFormattingIconDisplayConfigurationProperty{
//   	IconDisplayOption: jsii.String("iconDisplayOption"),
//   }
//
type CfnAnalysis_ConditionalFormattingIconDisplayConfigurationProperty struct {
	// Determines the icon display configuration.
	IconDisplayOption *string `field:"optional" json:"iconDisplayOption" yaml:"iconDisplayOption"`
}

