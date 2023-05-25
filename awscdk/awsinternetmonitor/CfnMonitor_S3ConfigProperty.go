package awsinternetmonitor


// The configuration for publishing Amazon CloudWatch Internet Monitor internet measurements to Amazon S3.
//
// The configuration includes the bucket name and (optionally) bucket prefix for the S3 bucket to store the measurements, and the delivery status. The delivery status is `ENABLED` if you choose to deliver internet measurements to S3 logs, and `DISABLED` otherwise.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ConfigProperty := &S3ConfigProperty{
//   	BucketName: jsii.String("bucketName"),
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   	LogDeliveryStatus: jsii.String("logDeliveryStatus"),
//   }
//
type CfnMonitor_S3ConfigProperty struct {
	// The Amazon S3 bucket name for internet measurements publishing.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// An optional Amazon S3 bucket prefix for internet measurements publishing.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The status of publishing Internet Monitor internet measurements to an Amazon S3 bucket.
	//
	// The delivery status is `ENABLED` if you choose to deliver internet measurements to an S3 bucket, and `DISABLED` otherwise.
	LogDeliveryStatus *string `field:"optional" json:"logDeliveryStatus" yaml:"logDeliveryStatus"`
}

