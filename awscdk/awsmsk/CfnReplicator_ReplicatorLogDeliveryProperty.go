package awsmsk


// Details of the log delivery for the replicator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicatorLogDeliveryProperty := &ReplicatorLogDeliveryProperty{
//   	CloudWatchLogs: &CloudWatchLogsProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		LogGroup: jsii.String("logGroup"),
//   	},
//   	Firehose: &FirehoseProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		DeliveryStream: jsii.String("deliveryStream"),
//   	},
//   	S3: &S3Property{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		Bucket: jsii.String("bucket"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicatorlogdelivery.html
//
type CfnReplicator_ReplicatorLogDeliveryProperty struct {
	// Details about delivering logs to CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicatorlogdelivery.html#cfn-msk-replicator-replicatorlogdelivery-cloudwatchlogs
	//
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Details about delivering logs to Firehose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicatorlogdelivery.html#cfn-msk-replicator-replicatorlogdelivery-firehose
	//
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// Details about delivering logs to S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicatorlogdelivery.html#cfn-msk-replicator-replicatorlogdelivery-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

