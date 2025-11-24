package mixinsawsquicksight


// Determines the icon display configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   conditionalFormattingIconDisplayConfigurationProperty := &ConditionalFormattingIconDisplayConfigurationProperty{
//   	IconDisplayOption: jsii.String("iconDisplayOption"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingicondisplayconfiguration.html
//
type CfnAnalysisPropsMixin_ConditionalFormattingIconDisplayConfigurationProperty struct {
	// Determines the icon display configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingicondisplayconfiguration.html#cfn-quicksight-analysis-conditionalformattingicondisplayconfiguration-icondisplayoption
	//
	IconDisplayOption *string `field:"optional" json:"iconDisplayOption" yaml:"iconDisplayOption"`
}

