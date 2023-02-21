package awsstepfunctions


// Defines the S3 bucket location where a state machine definition is stored.
//
// The state machine definition must be a JSON or YAML file.
//
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
//   	version: jsii.String("version"),
//   }
//
type CfnStateMachine_S3LocationProperty struct {
	// The name of the S3 bucket where the state machine definition JSON or YAML file is stored.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the state machine definition file (Amazon S3 object name).
	Key *string `field:"required" json:"key" yaml:"key"`
	// For versioning-enabled buckets, a specific version of the state machine definition.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

