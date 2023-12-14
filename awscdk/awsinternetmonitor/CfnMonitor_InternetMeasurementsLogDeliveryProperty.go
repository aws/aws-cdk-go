package awsinternetmonitor


// Publish internet measurements to an Amazon S3 bucket in addition to CloudWatch Logs.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-internetmeasurementslogdelivery.html
//
type CfnMonitor_InternetMeasurementsLogDeliveryProperty struct {
	// The configuration information for publishing Internet Monitor internet measurements to Amazon S3.
	//
	// The configuration includes the bucket name and (optionally) prefix for the S3 bucket to store the measurements, and the delivery status. The delivery status is `ENABLED` or `DISABLED` , depending on whether you choose to deliver internet measurements to S3 logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-internetmeasurementslogdelivery.html#cfn-internetmonitor-monitor-internetmeasurementslogdelivery-s3config
	//
	S3Config interface{} `field:"optional" json:"s3Config" yaml:"s3Config"`
}

