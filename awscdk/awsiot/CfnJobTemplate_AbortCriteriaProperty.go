package awsiot


// The criteria that determine when and how a job abort takes place.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   abortCriteriaProperty := &AbortCriteriaProperty{
//   	Action: jsii.String("action"),
//   	FailureType: jsii.String("failureType"),
//   	MinNumberOfExecutedThings: jsii.Number(123),
//   	ThresholdPercentage: jsii.Number(123),
//   }
//
type CfnJobTemplate_AbortCriteriaProperty struct {
	// The type of job action to take to initiate the job abort.
	Action *string `field:"required" json:"action" yaml:"action"`
	// The type of job execution failures that can initiate a job abort.
	FailureType *string `field:"required" json:"failureType" yaml:"failureType"`
	// The minimum number of things which must receive job execution notifications before the job can be aborted.
	MinNumberOfExecutedThings *float64 `field:"required" json:"minNumberOfExecutedThings" yaml:"minNumberOfExecutedThings"`
	// The minimum percentage of job execution failures that must occur to initiate the job abort.
	//
	// AWS IoT Core supports up to two digits after the decimal (for example, 10.9 and 10.99, but not 10.999).
	ThresholdPercentage *float64 `field:"required" json:"thresholdPercentage" yaml:"thresholdPercentage"`
}

