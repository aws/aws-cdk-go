package awsinternetmonitor


// Configuration information for other locations that you choose to publish Amazon CloudWatch Internet Monitor internet measurements to, such as Amazon S3.
//
// The measurements are also published to Amazon CloudWatch Logs.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-s3config.html
//
type CfnMonitor_S3ConfigProperty struct {
	// The Amazon S3 bucket name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-s3config.html#cfn-internetmonitor-monitor-s3config-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The Amazon S3 bucket prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-s3config.html#cfn-internetmonitor-monitor-s3config-bucketprefix
	//
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The status of publishing Internet Monitor internet measurements to an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-s3config.html#cfn-internetmonitor-monitor-s3config-logdeliverystatus
	//
	LogDeliveryStatus *string `field:"optional" json:"logDeliveryStatus" yaml:"logDeliveryStatus"`
}

