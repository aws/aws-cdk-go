package awsmsk


// Details of the log delivery for the replicator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   replicatorLogDeliveryProperty := &ReplicatorLogDeliveryProperty{
//   	CloudWatchLogs: &CloudWatchLogsProperty{
//   		Enabled: jsii.Boolean(false),
//   		LogGroup: jsii.String("logGroup"),
//   	},
//   	Firehose: &FirehoseProperty{
//   		DeliveryStream: jsii.String("deliveryStream"),
//   		Enabled: jsii.Boolean(false),
//   	},
//   	S3: &S3Property{
//   		Bucket: jsii.String("bucket"),
//   		Enabled: jsii.Boolean(false),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicatorlogdelivery.html
//
type CfnReplicatorPropsMixin_ReplicatorLogDeliveryProperty struct {
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

