package awssagemaker


// Describes the Docker container for the model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var modelInput interface{}
//
//   modelPackageContainerDefinitionProperty := &modelPackageContainerDefinitionProperty{
//   	image: jsii.String("image"),
//
//   	// the properties below are optional
//   	containerHostname: jsii.String("containerHostname"),
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	framework: jsii.String("framework"),
//   	frameworkVersion: jsii.String("frameworkVersion"),
//   	imageDigest: jsii.String("imageDigest"),
//   	modelDataUrl: jsii.String("modelDataUrl"),
//   	modelInput: modelInput,
//   	nearestModelName: jsii.String("nearestModelName"),
//   	productId: jsii.String("productId"),
//   }
//
type CfnModelPackage_ModelPackageContainerDefinitionProperty struct {
	// The Amazon EC2 Container Registry (Amazon ECR) path where inference code is stored.
	//
	// If you are using your own custom algorithm instead of an algorithm provided by SageMaker, the inference code must meet SageMaker requirements. SageMaker supports both `registry/repository[:tag]` and `registry/repository[@digest]` image path formats. For more information, see [Using Your Own Algorithms with Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html) .
	Image *string `field:"required" json:"image" yaml:"image"`
	// The DNS host name for the Docker container.
	ContainerHostname *string `field:"optional" json:"containerHostname" yaml:"containerHostname"`
	// The environment variables to set in the Docker container.
	//
	// Each key and value in the `Environment` string to string map can have length of up to 1024. We support up to 16 entries in the map.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The machine learning framework of the model package container image.
	Framework *string `field:"optional" json:"framework" yaml:"framework"`
	// The framework version of the Model Package Container Image.
	FrameworkVersion *string `field:"optional" json:"frameworkVersion" yaml:"frameworkVersion"`
	// An MD5 hash of the training algorithm that identifies the Docker image used for training.
	ImageDigest *string `field:"optional" json:"imageDigest" yaml:"imageDigest"`
	// The Amazon S3 path where the model artifacts, which result from model training, are stored.
	//
	// This path must point to a single `gzip` compressed tar archive ( `.tar.gz` suffix).
	//
	// > The model artifacts must be in an S3 bucket that is in the same region as the model package.
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
	// A structure with Model Input details.
	ModelInput interface{} `field:"optional" json:"modelInput" yaml:"modelInput"`
	// The name of a pre-trained machine learning benchmarked by Amazon SageMaker Inference Recommender model that matches your model.
	//
	// You can find a list of benchmarked models by calling `ListModelMetadata` .
	NearestModelName *string `field:"optional" json:"nearestModelName" yaml:"nearestModelName"`
	// The AWS Marketplace product ID of the model package.
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
}

