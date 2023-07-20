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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingmetric.html
//
type CfnModelCard_TrainingMetricProperty struct {
	// The name of the result from the SageMaker training job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingmetric.html#cfn-sagemaker-modelcard-trainingmetric-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of a result from the SageMaker training job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingmetric.html#cfn-sagemaker-modelcard-trainingmetric-value
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// Any additional notes describing the result of the training job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingmetric.html#cfn-sagemaker-modelcard-trainingmetric-notes
	//
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
}

