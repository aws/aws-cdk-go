package previewawsioteventsmixins


// Sends information about the detector model instance and the event that triggered the action to an Amazon SQS queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sqsProperty := &SqsProperty{
//   	Payload: &PayloadProperty{
//   		ContentExpression: jsii.String("contentExpression"),
//   		Type: jsii.String("type"),
//   	},
//   	QueueUrl: jsii.String("queueUrl"),
//   	UseBase64: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-sqs.html
//
type CfnAlarmModelPropsMixin_SqsProperty struct {
	// You can configure the action payload when you send a message to an Amazon SQS queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-sqs.html#cfn-iotevents-alarmmodel-sqs-payload
	//
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
	// The URL of the SQS queue where the data is written.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-sqs.html#cfn-iotevents-alarmmodel-sqs-queueurl
	//
	QueueUrl *string `field:"optional" json:"queueUrl" yaml:"queueUrl"`
	// Set this to TRUE if you want the data to be base-64 encoded before it is written to the queue.
	//
	// Otherwise, set this to FALSE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-sqs.html#cfn-iotevents-alarmmodel-sqs-usebase64
	//
	UseBase64 interface{} `field:"optional" json:"useBase64" yaml:"useBase64"`
}

