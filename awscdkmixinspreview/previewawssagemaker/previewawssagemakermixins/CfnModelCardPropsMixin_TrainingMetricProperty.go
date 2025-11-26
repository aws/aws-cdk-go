package previewawssagemakermixins


// A result from a SageMaker AI training job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   trainingMetricProperty := &TrainingMetricProperty{
//   	Name: jsii.String("name"),
//   	Notes: jsii.String("notes"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingmetric.html
//
type CfnModelCardPropsMixin_TrainingMetricProperty struct {
	// The name of the result from the SageMaker AI training job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingmetric.html#cfn-sagemaker-modelcard-trainingmetric-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Any additional notes describing the result of the training job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingmetric.html#cfn-sagemaker-modelcard-trainingmetric-notes
	//
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// The value of a result from the SageMaker AI training job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingmetric.html#cfn-sagemaker-modelcard-trainingmetric-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

