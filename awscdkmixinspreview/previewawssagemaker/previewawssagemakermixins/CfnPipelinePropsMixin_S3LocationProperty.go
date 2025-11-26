package previewawssagemakermixins


// The location of the pipeline definition stored in Amazon S3.
//
// If specified, SageMaker will retrieve the pipeline definition from this location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	ETag: jsii.String("eTag"),
//   	Key: jsii.String("key"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-s3location.html
//
type CfnPipelinePropsMixin_S3LocationProperty struct {
	// The name of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-s3location.html#cfn-sagemaker-pipeline-s3location-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// A file checksum of the pipeline definition file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-s3location.html#cfn-sagemaker-pipeline-s3location-etag
	//
	ETag *string `field:"optional" json:"eTag" yaml:"eTag"`
	// The object key (or key name) which uniquely identifies the object in an S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-s3location.html#cfn-sagemaker-pipeline-s3location-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The version ID of the pipeline definition file.
	//
	// If not specified, Amazon SageMaker will retrieve the latest version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-s3location.html#cfn-sagemaker-pipeline-s3location-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

