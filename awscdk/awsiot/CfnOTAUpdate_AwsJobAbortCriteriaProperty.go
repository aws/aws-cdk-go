package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsJobAbortCriteriaProperty := &AwsJobAbortCriteriaProperty{
//   	Action: jsii.String("action"),
//   	FailureType: jsii.String("failureType"),
//   	MinNumberOfExecutedThings: jsii.Number(123),
//   	ThresholdPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobabortcriteria.html
//
type CfnOTAUpdate_AwsJobAbortCriteriaProperty struct {
	// The type of job action to take to initiate the job abort.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobabortcriteria.html#cfn-iot-otaupdate-awsjobabortcriteria-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The type of job execution failures that can initiate a job abort.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobabortcriteria.html#cfn-iot-otaupdate-awsjobabortcriteria-failuretype
	//
	FailureType *string `field:"required" json:"failureType" yaml:"failureType"`
	// The minimum number of things which must receive job execution notifications before the job can be aborted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobabortcriteria.html#cfn-iot-otaupdate-awsjobabortcriteria-minnumberofexecutedthings
	//
	MinNumberOfExecutedThings *float64 `field:"required" json:"minNumberOfExecutedThings" yaml:"minNumberOfExecutedThings"`
	// The minimum percentage of job execution failures that must occur to initiate the job abort.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobabortcriteria.html#cfn-iot-otaupdate-awsjobabortcriteria-thresholdpercentage
	//
	ThresholdPercentage *float64 `field:"required" json:"thresholdPercentage" yaml:"thresholdPercentage"`
}

