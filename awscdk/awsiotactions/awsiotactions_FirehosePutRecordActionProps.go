package awsiotactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Configuration properties of an action for the Kinesis Data Firehose stream.
//
// Example:
//   import firehose "github.com/aws/aws-cdk-go/awscdk"
//   import destinations "github.com/aws/aws-cdk-go/awscdk"
//
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   stream := firehose.NewDeliveryStream(this, jsii.String("MyStream"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		destinations.NewS3Bucket(bucket),
//   	},
//   })
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewFirehosePutRecordAction(stream, &firehosePutRecordActionProps{
//   			batchMode: jsii.Boolean(true),
//   			recordSeparator: actions.firehoseRecordSeparator_NEWLINE,
//   		}),
//   	},
//   })
//
// Experimental.
type FirehosePutRecordActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Whether to deliver the Kinesis Data Firehose stream as a batch by using `PutRecordBatch`.
	//
	// When batchMode is true and the rule's SQL statement evaluates to an Array, each Array
	// element forms one record in the PutRecordBatch request. The resulting array can't have
	// more than 500 records.
	// Experimental.
	BatchMode *bool `field:"optional" json:"batchMode" yaml:"batchMode"`
	// A character separator that will be used to separate records written to the Kinesis Data Firehose stream.
	// Experimental.
	RecordSeparator FirehoseRecordSeparator `field:"optional" json:"recordSeparator" yaml:"recordSeparator"`
}

