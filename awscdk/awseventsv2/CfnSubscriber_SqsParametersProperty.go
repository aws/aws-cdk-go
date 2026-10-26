package awseventsv2


// Parameters for sending events to an Amazon SQS queue target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqsParametersProperty := &SqsParametersProperty{
//   	DelaySeconds: jsii.String("delaySeconds"),
//   	MessageAttributes: map[string]interface{}{
//   		"messageAttributesKey": &MessageAttributeValueProperty{
//   			"dataType": jsii.String("dataType"),
//
//   			// the properties below are optional
//   			"binaryValue": jsii.String("binaryValue"),
//   			"stringValue": jsii.String("stringValue"),
//   		},
//   	},
//   	MessageDeduplicationId: jsii.String("messageDeduplicationId"),
//   	MessageGroupId: jsii.String("messageGroupId"),
//   	MessageSystemAttributes: map[string]interface{}{
//   		"messageSystemAttributesKey": &MessageAttributeValueProperty{
//   			"dataType": jsii.String("dataType"),
//
//   			// the properties below are optional
//   			"binaryValue": jsii.String("binaryValue"),
//   			"stringValue": jsii.String("stringValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-sqsparameters.html
//
type CfnSubscriber_SqsParametersProperty struct {
	// The delay in seconds for the message, written as a string.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-sqsparameters.html#cfn-eventsv2-subscriber-sqsparameters-delayseconds
	//
	DelaySeconds *string `field:"optional" json:"delaySeconds" yaml:"delaySeconds"`
	// A map of message attribute names to values, with at most 50 attributes.
	//
	// Amazon SQS and Amazon SNS enforce their own per-message attribute limits on delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-sqsparameters.html#cfn-eventsv2-subscriber-sqsparameters-messageattributes
	//
	MessageAttributes interface{} `field:"optional" json:"messageAttributes" yaml:"messageAttributes"`
	// The message deduplication ID to use when the target is a FIFO queue.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-sqsparameters.html#cfn-eventsv2-subscriber-sqsparameters-messagededuplicationid
	//
	MessageDeduplicationId *string `field:"optional" json:"messageDeduplicationId" yaml:"messageDeduplicationId"`
	// The message group ID to use when the target is a FIFO queue.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-sqsparameters.html#cfn-eventsv2-subscriber-sqsparameters-messagegroupid
	//
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
	// A map of message attribute names to values, with at most 50 attributes.
	//
	// Amazon SQS and Amazon SNS enforce their own per-message attribute limits on delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-sqsparameters.html#cfn-eventsv2-subscriber-sqsparameters-messagesystemattributes
	//
	MessageSystemAttributes interface{} `field:"optional" json:"messageSystemAttributes" yaml:"messageSystemAttributes"`
}

