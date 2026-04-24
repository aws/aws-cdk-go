package awsmsk


// Configuration for log delivery for the replicator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logDeliveryProperty := &LogDeliveryProperty{
//   	ReplicatorLogDelivery: &ReplicatorLogDeliveryProperty{
//   		CloudWatchLogs: &CloudWatchLogsProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		Firehose: &FirehoseProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			DeliveryStream: jsii.String("deliveryStream"),
//   		},
//   		S3: &S3Property{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			Bucket: jsii.String("bucket"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-logdelivery.html
//
type CfnReplicator_LogDeliveryProperty struct {
	// Details of the log delivery for the replicator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-logdelivery.html#cfn-msk-replicator-logdelivery-replicatorlogdelivery
	//
	ReplicatorLogDelivery interface{} `field:"optional" json:"replicatorLogDelivery" yaml:"replicatorLogDelivery"`
}

