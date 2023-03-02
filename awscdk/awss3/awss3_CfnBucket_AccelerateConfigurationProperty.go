package awss3


// Configures the transfer acceleration state for an Amazon S3 bucket.
//
// For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accelerateConfigurationProperty := &accelerateConfigurationProperty{
//   	accelerationStatus: jsii.String("accelerationStatus"),
//   }
//
type CfnBucket_AccelerateConfigurationProperty struct {
	// Specifies the transfer acceleration status of the bucket.
	AccelerationStatus *string `field:"required" json:"accelerationStatus" yaml:"accelerationStatus"`
}

