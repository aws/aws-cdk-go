package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retryCriteriaProperty := &retryCriteriaProperty{
//   	failureType: jsii.String("failureType"),
//   	numberOfRetries: jsii.Number(123),
//   }
//
type CfnJobTemplate_RetryCriteriaProperty struct {
	// `CfnJobTemplate.RetryCriteriaProperty.FailureType`.
	FailureType *string `field:"optional" json:"failureType" yaml:"failureType"`
	// `CfnJobTemplate.RetryCriteriaProperty.NumberOfRetries`.
	NumberOfRetries *float64 `field:"optional" json:"numberOfRetries" yaml:"numberOfRetries"`
}

