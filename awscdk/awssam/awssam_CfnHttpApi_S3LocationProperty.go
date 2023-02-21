package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   	version: jsii.Number(123),
//   }
//
type CfnHttpApi_S3LocationProperty struct {
	// `CfnHttpApi.S3LocationProperty.Bucket`.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnHttpApi.S3LocationProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnHttpApi.S3LocationProperty.Version`.
	Version *float64 `field:"required" json:"version" yaml:"version"`
}

