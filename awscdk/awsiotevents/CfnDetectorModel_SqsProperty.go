package awsiotevents


// Sends information about the detector model instance and the event that triggered the action to an Amazon SQS queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqsProperty := &SqsProperty{
//   	QueueUrl: jsii.String("queueUrl"),
//
//   	// the properties below are optional
//   	Payload: &PayloadProperty{
//   		ContentExpression: jsii.String("contentExpression"),
//   		Type: jsii.String("type"),
//   	},
//   	UseBase64: jsii.Boolean(false),
//   }
//
type CfnDetectorModel_SqsProperty struct {
	// The URL of the SQS queue where the data is written.
	QueueUrl *string `field:"required" json:"queueUrl" yaml:"queueUrl"`
	// You can configure the action payload when you send a message to an Amazon SQS queue.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
	// Set this to TRUE if you want the data to be base-64 encoded before it is written to the queue.
	//
	// Otherwise, set this to FALSE.
	UseBase64 interface{} `field:"optional" json:"useBase64" yaml:"useBase64"`
}

