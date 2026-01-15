package previewawssesmixins


// The overall confidence threshold settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   overallConfidenceThresholdProperty := &OverallConfidenceThresholdProperty{
//   	ConfidenceVerdictThreshold: jsii.String("confidenceVerdictThreshold"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-overallconfidencethreshold.html
//
type CfnConfigurationSetPropsMixin_OverallConfidenceThresholdProperty struct {
	// The confidence verdict threshold level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-overallconfidencethreshold.html#cfn-ses-configurationset-overallconfidencethreshold-confidenceverdictthreshold
	//
	ConfidenceVerdictThreshold *string `field:"optional" json:"confidenceVerdictThreshold" yaml:"confidenceVerdictThreshold"`
}

