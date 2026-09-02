package awsiot


// The criteria that determine when and how a job abort takes place.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   abortCriteriaProperty := &AbortCriteriaProperty{
//   	Action: jsii.String("action"),
//   	FailureType: jsii.String("failureType"),
//   	MinNumberOfExecutedThings: jsii.Number(123),
//   	ThresholdPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-abortcriteria.html
//
type CfnJobPropsMixin_AbortCriteriaProperty struct {
	// The type of job action to take to initiate the job abort.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-abortcriteria.html#cfn-iot-job-abortcriteria-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The type of job execution failures that can initiate a job abort.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-abortcriteria.html#cfn-iot-job-abortcriteria-failuretype
	//
	FailureType *string `field:"optional" json:"failureType" yaml:"failureType"`
	// The minimum number of things which must receive job execution notifications before the job can be aborted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-abortcriteria.html#cfn-iot-job-abortcriteria-minnumberofexecutedthings
	//
	MinNumberOfExecutedThings *float64 `field:"optional" json:"minNumberOfExecutedThings" yaml:"minNumberOfExecutedThings"`
	// The minimum percentage of job execution failures that must occur to initiate the job abort.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-abortcriteria.html#cfn-iot-job-abortcriteria-thresholdpercentage
	//
	ThresholdPercentage *float64 `field:"optional" json:"thresholdPercentage" yaml:"thresholdPercentage"`
}

