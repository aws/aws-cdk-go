package awsinternetmonitor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   internetMeasurementsLogDeliveryProperty := &InternetMeasurementsLogDeliveryProperty{
//   	S3Config: &S3ConfigProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		LogDeliveryStatus: jsii.String("logDeliveryStatus"),
//   	},
//   }
//
type CfnMonitor_InternetMeasurementsLogDeliveryProperty struct {
	// The configuration information for publishing Amazon CloudWatch Internet Monitor internet measurements to Amazon S3.
	//
	// The configuration includes the bucket name and (optionally) bucket prefix for the S3 bucket to store the measurements, and the delivery status. The delivery status is `ENABLED` if you choose to deliver internet measurements to an S3 bucket, and `DISABLED` otherwise.
	S3Config interface{} `field:"optional" json:"s3Config" yaml:"s3Config"`
}

