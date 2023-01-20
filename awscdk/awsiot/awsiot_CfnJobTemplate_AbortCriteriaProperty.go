package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   abortCriteriaProperty := &abortCriteriaProperty{
//   	action: jsii.String("action"),
//   	failureType: jsii.String("failureType"),
//   	minNumberOfExecutedThings: jsii.Number(123),
//   	thresholdPercentage: jsii.Number(123),
//   }
//
type CfnJobTemplate_AbortCriteriaProperty struct {
	// `CfnJobTemplate.AbortCriteriaProperty.Action`.
	Action *string `field:"required" json:"action" yaml:"action"`
	// `CfnJobTemplate.AbortCriteriaProperty.FailureType`.
	FailureType *string `field:"required" json:"failureType" yaml:"failureType"`
	// `CfnJobTemplate.AbortCriteriaProperty.MinNumberOfExecutedThings`.
	MinNumberOfExecutedThings *float64 `field:"required" json:"minNumberOfExecutedThings" yaml:"minNumberOfExecutedThings"`
	// `CfnJobTemplate.AbortCriteriaProperty.ThresholdPercentage`.
	ThresholdPercentage *float64 `field:"required" json:"thresholdPercentage" yaml:"thresholdPercentage"`
}

