package awsiotevents


// Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firehoseProperty := &firehoseProperty{
//   	deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   	separator: jsii.String("separator"),
//   }
//
type CfnAlarmModel_FirehoseProperty struct {
	// The name of the Kinesis Data Firehose delivery stream where the data is written.
	DeliveryStreamName *string `field:"required" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// You can configure the action payload when you send a message to an Amazon Kinesis Data Firehose delivery stream.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
	// A character separator that is used to separate records written to the Kinesis Data Firehose delivery stream.
	//
	// Valid values are: '\n' (newline), '\t' (tab), '\r\n' (Windows newline), ',' (comma).
	Separator *string `field:"optional" json:"separator" yaml:"separator"`
}

