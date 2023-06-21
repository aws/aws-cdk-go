package awsec2


// Describes the options for Verified Access logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   verifiedAccessLogsProperty := &VerifiedAccessLogsProperty{
//   	CloudWatchLogs: &CloudWatchLogsProperty{
//   		Enabled: jsii.Boolean(false),
//   		LogGroup: jsii.String("logGroup"),
//   	},
//   	KinesisDataFirehose: &KinesisDataFirehoseProperty{
//   		DeliveryStream: jsii.String("deliveryStream"),
//   		Enabled: jsii.Boolean(false),
//   	},
//   	S3: &S3Property{
//   		BucketName: jsii.String("bucketName"),
//   		BucketOwner: jsii.String("bucketOwner"),
//   		Enabled: jsii.Boolean(false),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
type CfnVerifiedAccessInstance_VerifiedAccessLogsProperty struct {
	// CloudWatch Logs logging destination.
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Kinesis logging destination.
	KinesisDataFirehose interface{} `field:"optional" json:"kinesisDataFirehose" yaml:"kinesisDataFirehose"`
	// Amazon S3 logging options.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

