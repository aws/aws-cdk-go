package awsgreengrassv2


// Contains criteria that define when and how to cancel a job.
//
// The deployment stops if the following conditions are true:
//
// - The number of things that receive the deployment exceeds the `minNumberOfExecutedThings` .
// - The percentage of failures with type `failureType` exceeds the `thresholdPercentage` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ioTJobAbortCriteriaProperty := &IoTJobAbortCriteriaProperty{
//   	Action: jsii.String("action"),
//   	FailureType: jsii.String("failureType"),
//   	MinNumberOfExecutedThings: jsii.Number(123),
//   	ThresholdPercentage: jsii.Number(123),
//   }
//
type CfnDeployment_IoTJobAbortCriteriaProperty struct {
	// The action to perform when the criteria are met.
	Action *string `field:"required" json:"action" yaml:"action"`
	// The type of job deployment failure that can cancel a job.
	FailureType *string `field:"required" json:"failureType" yaml:"failureType"`
	// The minimum number of things that receive the configuration before the job can cancel.
	MinNumberOfExecutedThings *float64 `field:"required" json:"minNumberOfExecutedThings" yaml:"minNumberOfExecutedThings"`
	// The minimum percentage of `failureType` failures that occur before the job can cancel.
	//
	// This parameter supports up to two digits after the decimal (for example, you can specify `10.9` or `10.99` , but not `10.999` ).
	ThresholdPercentage *float64 `field:"required" json:"thresholdPercentage" yaml:"thresholdPercentage"`
}

