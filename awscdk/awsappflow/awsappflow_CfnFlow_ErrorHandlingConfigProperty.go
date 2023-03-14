package awsappflow


// The settings that determine how Amazon AppFlow handles an error when placing data in the destination.
//
// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. `ErrorHandlingConfig` is a part of the destination connector details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   errorHandlingConfigProperty := &errorHandlingConfigProperty{
//   	bucketName: jsii.String("bucketName"),
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	failOnFirstError: jsii.Boolean(false),
//   }
//
type CfnFlow_ErrorHandlingConfigProperty struct {
	// Specifies the name of the Amazon S3 bucket.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Specifies the Amazon S3 bucket prefix.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Specifies if the flow should fail after the first instance of a failure when attempting to place data in the destination.
	FailOnFirstError interface{} `field:"optional" json:"failOnFirstError" yaml:"failOnFirstError"`
}

