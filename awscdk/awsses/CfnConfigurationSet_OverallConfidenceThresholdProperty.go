package awsses


// The overall confidence threshold settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   overallConfidenceThresholdProperty := &OverallConfidenceThresholdProperty{
//   	ConfidenceVerdictThreshold: jsii.String("confidenceVerdictThreshold"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-overallconfidencethreshold.html
//
type CfnConfigurationSet_OverallConfidenceThresholdProperty struct {
	// The confidence verdict threshold level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-overallconfidencethreshold.html#cfn-ses-configurationset-overallconfidencethreshold-confidenceverdictthreshold
	//
	ConfidenceVerdictThreshold *string `field:"required" json:"confidenceVerdictThreshold" yaml:"confidenceVerdictThreshold"`
}

