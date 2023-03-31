package awsiot


// Describes an action to publish data to an Amazon SQS queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqsActionProperty := &sqsActionProperty{
//   	queueUrl: jsii.String("queueUrl"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	useBase64: jsii.Boolean(false),
//   }
//
type CfnTopicRule_SqsActionProperty struct {
	// The URL of the Amazon SQS queue.
	QueueUrl *string `field:"required" json:"queueUrl" yaml:"queueUrl"`
	// The ARN of the IAM role that grants access.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Specifies whether to use Base64 encoding.
	UseBase64 interface{} `field:"optional" json:"useBase64" yaml:"useBase64"`
}

