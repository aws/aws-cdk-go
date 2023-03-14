package awssagemaker


// A result from a SageMaker training job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trainingMetricProperty := &TrainingMetricProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.Number(123),
//
//   	// the properties below are optional
//   	Notes: jsii.String("notes"),
//   }
//
type CfnModelCard_TrainingMetricProperty struct {
	// The name of the result from the SageMaker training job.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of a result from the SageMaker training job.
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// Any additional notes describing the result of the training job.
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
}

