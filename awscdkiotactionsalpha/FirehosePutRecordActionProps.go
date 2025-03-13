package awscdkiotactionsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration properties of an action for the Kinesis Data Firehose stream.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   stream := firehose.NewDeliveryStream(this, jsii.String("MyStream"), &DeliveryStreamProps{
//   	Destination: firehose.NewS3Bucket(bucket),
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
type FirehosePutRecordActionProps struct {
	// The IAM role that allows access to AWS service.
	// Default: a new role will be created.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Whether to deliver the Kinesis Data Firehose stream as a batch by using `PutRecordBatch`.
	//
	// When batchMode is true and the rule's SQL statement evaluates to an Array, each Array
	// element forms one record in the PutRecordBatch request. The resulting array can't have
	// more than 500 records.
	// Default: false.
	//
	// Experimental.
	BatchMode *bool `field:"optional" json:"batchMode" yaml:"batchMode"`
	// A character separator that will be used to separate records written to the Kinesis Data Firehose stream.
	// Default: - none -- the stream does not use a separator.
	//
	// Experimental.
	RecordSeparator FirehoseRecordSeparator `field:"optional" json:"recordSeparator" yaml:"recordSeparator"`
}

