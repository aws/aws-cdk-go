package awseventsv2


// Parameters for publishing events to an Amazon SNS topic target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snsParametersProperty := &SnsParametersProperty{
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
//   	MessageStructure: jsii.String("messageStructure"),
//   	Subject: jsii.String("subject"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-snsparameters.html
//
type CfnSubscriber_SnsParametersProperty struct {
	// A map of message attribute names to values, with at most 50 attributes.
	//
	// Amazon SQS and Amazon SNS enforce their own per-message attribute limits on delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-snsparameters.html#cfn-eventsv2-subscriber-snsparameters-messageattributes
	//
	MessageAttributes interface{} `field:"optional" json:"messageAttributes" yaml:"messageAttributes"`
	// The message deduplication ID to use when the target is a FIFO topic.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-snsparameters.html#cfn-eventsv2-subscriber-snsparameters-messagededuplicationid
	//
	MessageDeduplicationId *string `field:"optional" json:"messageDeduplicationId" yaml:"messageDeduplicationId"`
	// The message group ID to use when the target is a FIFO topic.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-snsparameters.html#cfn-eventsv2-subscriber-snsparameters-messagegroupid
	//
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
	// Set to json to send a different message per delivery protocol.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-snsparameters.html#cfn-eventsv2-subscriber-snsparameters-messagestructure
	//
	MessageStructure *string `field:"optional" json:"messageStructure" yaml:"messageStructure"`
	// The subject line to use for email-protocol subscriptions.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-snsparameters.html#cfn-eventsv2-subscriber-snsparameters-subject
	//
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

