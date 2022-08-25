package awsappflow


// Determines how Amazon AppFlow handles the success response that it gets from the connector after placing data.
//
// For example, this setting would determine where to write the response from the destination connector upon a successful insert operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   successResponseHandlingConfigProperty := &successResponseHandlingConfigProperty{
//   	bucketName: jsii.String("bucketName"),
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   }
//
type CfnFlow_SuccessResponseHandlingConfigProperty struct {
	// The name of the Amazon S3 bucket.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The Amazon S3 bucket prefix.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
}

