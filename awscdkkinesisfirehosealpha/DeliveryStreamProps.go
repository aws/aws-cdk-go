package awscdkkinesisfirehosealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a new delivery stream.
//
// Example:
//   import firehose "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//   import destinations "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
//
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   stream := firehose.NewDeliveryStream(this, jsii.String("MyStream"), &DeliveryStreamProps{
//   	Destinations: []iDestination{
//   		destinations.NewS3Bucket(bucket),
//   	},
//   })
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewFirehosePutRecordAction(stream, &FirehosePutRecordActionProps{
//   			BatchMode: jsii.Boolean(true),
//   			RecordSeparator: actions.FirehoseRecordSeparator_NEWLINE,
//   		}),
//   	},
//   })
//
// Experimental.
type DeliveryStreamProps struct {
	// The destinations that this delivery stream will deliver data to.
	//
	// Only a singleton array is supported at this time.
	// Experimental.
	Destinations *[]IDestination `field:"required" json:"destinations" yaml:"destinations"`
	// A name for the delivery stream.
	// Experimental.
	DeliveryStreamName *string `field:"optional" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// Indicates the type of customer master key (CMK) to use for server-side encryption, if any.
	// Experimental.
	Encryption StreamEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// Customer managed key to server-side encrypt data in the stream.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The IAM role associated with this delivery stream.
	//
	// Assumed by Kinesis Data Firehose to read from sources and encrypt data server-side.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The Kinesis data stream to use as a source for this delivery stream.
	// Experimental.
	SourceStream awskinesis.IStream `field:"optional" json:"sourceStream" yaml:"sourceStream"`
}

