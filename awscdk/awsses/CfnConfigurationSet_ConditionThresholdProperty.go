package awsses


// The condition threshold settings for suppression validation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionThresholdProperty := &ConditionThresholdProperty{
//   	ConditionThresholdEnabled: jsii.String("conditionThresholdEnabled"),
//
//   	// the properties below are optional
//   	OverallConfidenceThreshold: &OverallConfidenceThresholdProperty{
//   		ConfidenceVerdictThreshold: jsii.String("confidenceVerdictThreshold"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-conditionthreshold.html
//
type CfnConfigurationSet_ConditionThresholdProperty struct {
	// Whether the condition threshold is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-conditionthreshold.html#cfn-ses-configurationset-conditionthreshold-conditionthresholdenabled
	//
	ConditionThresholdEnabled *string `field:"required" json:"conditionThresholdEnabled" yaml:"conditionThresholdEnabled"`
	// The overall confidence threshold settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-conditionthreshold.html#cfn-ses-configurationset-conditionthreshold-overallconfidencethreshold
	//
	OverallConfidenceThreshold interface{} `field:"optional" json:"overallConfidenceThreshold" yaml:"overallConfidenceThreshold"`
}

