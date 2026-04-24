package awsmsk


// Configuration for log delivery for the replicator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   logDeliveryProperty := &LogDeliveryProperty{
//   	ReplicatorLogDelivery: &ReplicatorLogDeliveryProperty{
//   		CloudWatchLogs: &CloudWatchLogsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		Firehose: &FirehoseProperty{
//   			DeliveryStream: jsii.String("deliveryStream"),
//   			Enabled: jsii.Boolean(false),
//   		},
//   		S3: &S3Property{
//   			Bucket: jsii.String("bucket"),
//   			Enabled: jsii.Boolean(false),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-logdelivery.html
//
type CfnReplicatorPropsMixin_LogDeliveryProperty struct {
	// Details of the log delivery for the replicator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-logdelivery.html#cfn-msk-replicator-logdelivery-replicatorlogdelivery
	//
	ReplicatorLogDelivery interface{} `field:"optional" json:"replicatorLogDelivery" yaml:"replicatorLogDelivery"`
}

