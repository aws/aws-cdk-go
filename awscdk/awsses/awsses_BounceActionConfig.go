package awsses


// BoundAction configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bounceActionConfig := &bounceActionConfig{
//   	message: jsii.String("message"),
//   	sender: jsii.String("sender"),
//   	smtpReplyCode: jsii.String("smtpReplyCode"),
//
//   	// the properties below are optional
//   	statusCode: jsii.String("statusCode"),
//   	topicArn: jsii.String("topicArn"),
//   }
//
// Experimental.
type BounceActionConfig struct {
	// Human-readable text to include in the bounce message.
	// Experimental.
	Message *string `field:"required" json:"message" yaml:"message"`
	// The email address of the sender of the bounced email.
	//
	// This is the address that the bounce message is sent from.
	// Experimental.
	Sender *string `field:"required" json:"sender" yaml:"sender"`
	// The SMTP reply code, as defined by RFC 5321.
	// Experimental.
	SmtpReplyCode *string `field:"required" json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// The SMTP enhanced status code, as defined by RFC 3463.
	// Experimental.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the bounce action is taken.
	// Experimental.
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

