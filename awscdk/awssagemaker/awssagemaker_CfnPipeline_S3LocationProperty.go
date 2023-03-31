package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	eTag: jsii.String("eTag"),
//   	version: jsii.String("version"),
//   }
//
type CfnPipeline_S3LocationProperty struct {
	// `CfnPipeline.S3LocationProperty.Bucket`.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnPipeline.S3LocationProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnPipeline.S3LocationProperty.ETag`.
	ETag *string `field:"optional" json:"eTag" yaml:"eTag"`
	// `CfnPipeline.S3LocationProperty.Version`.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

