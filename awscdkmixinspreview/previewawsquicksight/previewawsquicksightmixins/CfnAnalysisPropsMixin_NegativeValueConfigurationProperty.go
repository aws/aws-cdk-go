package previewawsquicksightmixins


// The options that determine the negative value configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   negativeValueConfigurationProperty := &NegativeValueConfigurationProperty{
//   	DisplayMode: jsii.String("displayMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-negativevalueconfiguration.html
//
type CfnAnalysisPropsMixin_NegativeValueConfigurationProperty struct {
	// Determines the display mode of the negative value configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-negativevalueconfiguration.html#cfn-quicksight-analysis-negativevalueconfiguration-displaymode
	//
	DisplayMode *string `field:"optional" json:"displayMode" yaml:"displayMode"`
}

