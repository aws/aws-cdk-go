package previewawssesmixins


// When included in a receipt rule, this action rejects the received email by returning a bounce response to the sender and, optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).
//
// For information about sending a bounce message in response to a received email, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-bounce.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bounceActionProperty := &BounceActionProperty{
//   	Message: jsii.String("message"),
//   	Sender: jsii.String("sender"),
//   	SmtpReplyCode: jsii.String("smtpReplyCode"),
//   	StatusCode: jsii.String("statusCode"),
//   	TopicArn: jsii.String("topicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-bounceaction.html
//
type CfnReceiptRulePropsMixin_BounceActionProperty struct {
	// Human-readable text to include in the bounce message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-bounceaction.html#cfn-ses-receiptrule-bounceaction-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
	// The email address of the sender of the bounced email.
	//
	// This is the address from which the bounce message is sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-bounceaction.html#cfn-ses-receiptrule-bounceaction-sender
	//
	Sender *string `field:"optional" json:"sender" yaml:"sender"`
	// The SMTP reply code, as defined by [RFC 5321](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc5321) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-bounceaction.html#cfn-ses-receiptrule-bounceaction-smtpreplycode
	//
	SmtpReplyCode *string `field:"optional" json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// The SMTP enhanced status code, as defined by [RFC 3463](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc3463) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-bounceaction.html#cfn-ses-receiptrule-bounceaction-statuscode
	//
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the bounce action is taken.
	//
	// You can find the ARN of a topic by using the [ListTopics](https://docs.aws.amazon.com/sns/latest/api/API_ListTopics.html) operation in Amazon SNS.
	//
	// For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-bounceaction.html#cfn-ses-receiptrule-bounceaction-topicarn
	//
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

