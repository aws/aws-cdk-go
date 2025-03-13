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
//   s3LocationProperty := &S3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-s3location.html
//
type CfnStateMachine_S3LocationProperty struct {
	// The name of the S3 bucket where the state machine definition JSON or YAML file is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-s3location.html#cfn-stepfunctions-statemachine-s3location-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the state machine definition file (Amazon S3 object name).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-s3location.html#cfn-stepfunctions-statemachine-s3location-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// For versioning-enabled buckets, a specific version of the state machine definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-s3location.html#cfn-stepfunctions-statemachine-s3location-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

