package awsgreengrassv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ioTJobAbortCriteriaProperty := &ioTJobAbortCriteriaProperty{
//   	action: jsii.String("action"),
//   	failureType: jsii.String("failureType"),
//   	minNumberOfExecutedThings: jsii.Number(123),
//   	thresholdPercentage: jsii.Number(123),
//   }
//
type CfnDeployment_IoTJobAbortCriteriaProperty struct {
	// `CfnDeployment.IoTJobAbortCriteriaProperty.Action`.
	Action *string `field:"required" json:"action" yaml:"action"`
	// `CfnDeployment.IoTJobAbortCriteriaProperty.FailureType`.
	FailureType *string `field:"required" json:"failureType" yaml:"failureType"`
	// `CfnDeployment.IoTJobAbortCriteriaProperty.MinNumberOfExecutedThings`.
	MinNumberOfExecutedThings *float64 `field:"required" json:"minNumberOfExecutedThings" yaml:"minNumberOfExecutedThings"`
	// `CfnDeployment.IoTJobAbortCriteriaProperty.ThresholdPercentage`.
	ThresholdPercentage *float64 `field:"required" json:"thresholdPercentage" yaml:"thresholdPercentage"`
}

