package previewawssagemakermixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   containerProperty := &ContainerProperty{
//   	Image: jsii.String("image"),
//   	ModelDataUrl: jsii.String("modelDataUrl"),
//   	NearestModelName: jsii.String("nearestModelName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-container.html
//
type CfnModelCardPropsMixin_ContainerProperty struct {
	// Inference environment path.
	//
	// The Amazon EC2 Container Registry (Amazon ECR) path where inference code is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-container.html#cfn-sagemaker-modelcard-container-image
	//
	Image *string `field:"optional" json:"image" yaml:"image"`
	// The Amazon S3 path where the model artifacts, which result from model training, are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-container.html#cfn-sagemaker-modelcard-container-modeldataurl
	//
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
	// The name of a pre-trained machine learning benchmarked by Amazon SageMaker Inference Recommender model that matches your model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-container.html#cfn-sagemaker-modelcard-container-nearestmodelname
	//
	NearestModelName *string `field:"optional" json:"nearestModelName" yaml:"nearestModelName"`
}

