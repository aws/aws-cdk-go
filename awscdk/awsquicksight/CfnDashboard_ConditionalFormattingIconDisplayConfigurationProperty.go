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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattingicondisplayconfiguration.html
//
type CfnDashboard_ConditionalFormattingIconDisplayConfigurationProperty struct {
	// Determines the icon display configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattingicondisplayconfiguration.html#cfn-quicksight-dashboard-conditionalformattingicondisplayconfiguration-icondisplayoption
	//
	IconDisplayOption *string `field:"optional" json:"iconDisplayOption" yaml:"iconDisplayOption"`
}

