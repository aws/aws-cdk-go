package awssam


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
//   	Version: jsii.Number(123),
//   }
//
type CfnFunction_S3LocationProperty struct {
	// `CfnFunction.S3LocationProperty.Bucket`.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnFunction.S3LocationProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnFunction.S3LocationProperty.Version`.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

