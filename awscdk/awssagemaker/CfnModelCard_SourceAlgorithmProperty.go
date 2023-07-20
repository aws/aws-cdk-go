package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceAlgorithmProperty := &SourceAlgorithmProperty{
//   	AlgorithmName: jsii.String("algorithmName"),
//
//   	// the properties below are optional
//   	ModelDataUrl: jsii.String("modelDataUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-sourcealgorithm.html
//
type CfnModelCard_SourceAlgorithmProperty struct {
	// The name of an algorithm that was used to create the model package.
	//
	// The algorithm must be either an algorithm resource in your SageMaker account or an algorithm in AWS Marketplace that you are subscribed to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-sourcealgorithm.html#cfn-sagemaker-modelcard-sourcealgorithm-algorithmname
	//
	AlgorithmName *string `field:"required" json:"algorithmName" yaml:"algorithmName"`
	// The Amazon S3 path where the model artifacts, which result from model training, are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-sourcealgorithm.html#cfn-sagemaker-modelcard-sourcealgorithm-modeldataurl
	//
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
}

