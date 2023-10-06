package awssagemaker


// The location of the pipeline definition stored in Amazon S3.
//
// If specified, SageMaker will retrieve the pipeline definition from this location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	ETag: jsii.String("eTag"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-s3location.html
//
type CfnPipeline_S3LocationProperty struct {
	// The name of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-s3location.html#cfn-sagemaker-pipeline-s3location-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The object key (or key name) which uniquely identifies the object in an S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-s3location.html#cfn-sagemaker-pipeline-s3location-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// A file checksum of the pipeline definition file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-s3location.html#cfn-sagemaker-pipeline-s3location-etag
	//
	ETag *string `field:"optional" json:"eTag" yaml:"eTag"`
	// The version ID of the pipeline definition file.
	//
	// If not specified, Amazon SageMaker will retrieve the latest version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-s3location.html#cfn-sagemaker-pipeline-s3location-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

