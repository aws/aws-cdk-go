package awsrobomaker


// Information about a source configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceConfigProperty := &sourceConfigProperty{
//   	architecture: jsii.String("architecture"),
//   	s3Bucket: jsii.String("s3Bucket"),
//   	s3Key: jsii.String("s3Key"),
//   }
//
type CfnSimulationApplication_SourceConfigProperty struct {
	// The target processor architecture for the application.
	Architecture *string `field:"required" json:"architecture" yaml:"architecture"`
	// The Amazon S3 bucket name.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The s3 object key.
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
}

