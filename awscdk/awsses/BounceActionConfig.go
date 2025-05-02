package awsses


// BoundAction configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bounceActionConfig := &BounceActionConfig{
//   	Message: jsii.String("message"),
//   	Sender: jsii.String("sender"),
//   	SmtpReplyCode: jsii.String("smtpReplyCode"),
//
//   	// the properties below are optional
//   	StatusCode: jsii.String("statusCode"),
//   	TopicArn: jsii.String("topicArn"),
//   }
//
type BounceActionConfig struct {
	// Human-readable text to include in the bounce message.
	Message *string `field:"required" json:"message" yaml:"message"`
	// The email address of the sender of the bounced email.
	//
	// This is the address that the bounce message is sent from.
	Sender *string `field:"required" json:"sender" yaml:"sender"`
	// The SMTP reply code, as defined by RFC 5321.
	SmtpReplyCode *string `field:"required" json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// The SMTP enhanced status code, as defined by RFC 3463.
	// Default: - No status code.
	//
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the bounce action is taken.
	// Default: - No notification is sent to SNS.
	//
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

