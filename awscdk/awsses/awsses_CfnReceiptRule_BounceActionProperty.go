package awsses


// When included in a receipt rule, this action rejects the received email by returning a bounce response to the sender and, optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).
//
// For information about sending a bounce message in response to a received email, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-bounce.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bounceActionProperty := &bounceActionProperty{
//   	message: jsii.String("message"),
//   	sender: jsii.String("sender"),
//   	smtpReplyCode: jsii.String("smtpReplyCode"),
//
//   	// the properties below are optional
//   	statusCode: jsii.String("statusCode"),
//   	topicArn: jsii.String("topicArn"),
//   }
//
type CfnReceiptRule_BounceActionProperty struct {
	// Human-readable text to include in the bounce message.
	Message *string `field:"required" json:"message" yaml:"message"`
	// The email address of the sender of the bounced email.
	//
	// This is the address from which the bounce message is sent.
	Sender *string `field:"required" json:"sender" yaml:"sender"`
	// The SMTP reply code, as defined by [RFC 5321](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc5321) .
	SmtpReplyCode *string `field:"required" json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// The SMTP enhanced status code, as defined by [RFC 3463](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc3463) .
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the bounce action is taken.
	//
	// You can find the ARN of a topic by using the [ListTopics](https://docs.aws.amazon.com/sns/latest/api/API_ListTopics.html) operation in Amazon SNS.
	//
	// For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html) .
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

