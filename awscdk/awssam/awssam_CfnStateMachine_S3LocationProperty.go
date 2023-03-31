package awssam


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
//   	version: jsii.Number(123),
//   }
//
type CfnStateMachine_S3LocationProperty struct {
	// `CfnStateMachine.S3LocationProperty.Bucket`.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnStateMachine.S3LocationProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnStateMachine.S3LocationProperty.Version`.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

