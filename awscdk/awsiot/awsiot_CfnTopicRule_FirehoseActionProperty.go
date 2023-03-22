package awsiot


// Describes an action that writes data to an Amazon Kinesis Firehose stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firehoseActionProperty := &firehoseActionProperty{
//   	deliveryStreamName: jsii.String("deliveryStreamName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	batchMode: jsii.Boolean(false),
//   	separator: jsii.String("separator"),
//   }
//
type CfnTopicRule_FirehoseActionProperty struct {
	// The delivery stream name.
	DeliveryStreamName *string `field:"required" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// The IAM role that grants access to the Amazon Kinesis Firehose stream.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Whether to deliver the Kinesis Data Firehose stream as a batch by using [`PutRecordBatch`](https://docs.aws.amazon.com/firehose/latest/APIReference/API_PutRecordBatch.html) . The default value is `false` .
	//
	// When `batchMode` is `true` and the rule's SQL statement evaluates to an Array, each Array element forms one record in the [`PutRecordBatch`](https://docs.aws.amazon.com/firehose/latest/APIReference/API_PutRecordBatch.html) request. The resulting array can't have more than 500 records.
	BatchMode interface{} `field:"optional" json:"batchMode" yaml:"batchMode"`
	// A character separator that will be used to separate records written to the Firehose stream.
	//
	// Valid values are: '\n' (newline), '\t' (tab), '\r\n' (Windows newline), ',' (comma).
	Separator *string `field:"optional" json:"separator" yaml:"separator"`
}

