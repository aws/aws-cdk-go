package awssagemaker


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
	// The name of the S3 bucket where the PipelineDefinition file is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-s3location.html#cfn-sagemaker-pipeline-s3location-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The file name of the PipelineDefinition file (Amazon S3 object name).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-s3location.html#cfn-sagemaker-pipeline-s3location-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The Amazon S3 ETag (a file checksum) of the PipelineDefinition file.
	//
	// If you don't specify a value, SageMaker skips ETag validation of your PipelineDefinition file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-s3location.html#cfn-sagemaker-pipeline-s3location-etag
	//
	ETag *string `field:"optional" json:"eTag" yaml:"eTag"`
	// For versioning-enabled buckets, a specific version of the PipelineDefinition file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-s3location.html#cfn-sagemaker-pipeline-s3location-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

