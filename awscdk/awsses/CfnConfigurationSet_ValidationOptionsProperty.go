package awsses


// An object that contains information about the validation options for your account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   validationOptionsProperty := &ValidationOptionsProperty{
//   	ConditionThreshold: &ConditionThresholdProperty{
//   		ConditionThresholdEnabled: jsii.String("conditionThresholdEnabled"),
//
//   		// the properties below are optional
//   		OverallConfidenceThreshold: &OverallConfidenceThresholdProperty{
//   			ConfidenceVerdictThreshold: jsii.String("confidenceVerdictThreshold"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-validationoptions.html
//
type CfnConfigurationSet_ValidationOptionsProperty struct {
	// The condition threshold settings for suppression validation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-validationoptions.html#cfn-ses-configurationset-validationoptions-conditionthreshold
	//
	ConditionThreshold interface{} `field:"required" json:"conditionThreshold" yaml:"conditionThreshold"`
}

