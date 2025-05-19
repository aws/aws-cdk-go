package awscognito


// Result of binding email settings with a user pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolEmailConfig := &UserPoolEmailConfig{
//   	ConfigurationSet: jsii.String("configurationSet"),
//   	EmailSendingAccount: jsii.String("emailSendingAccount"),
//   	From: jsii.String("from"),
//   	ReplyToEmailAddress: jsii.String("replyToEmailAddress"),
//   	SourceArn: jsii.String("sourceArn"),
//   }
//
type UserPoolEmailConfig struct {
	// The name of the configuration set in SES.
	// Default: - none.
	//
	ConfigurationSet *string `field:"optional" json:"configurationSet" yaml:"configurationSet"`
	// Specifies whether to use Cognito's built in email functionality or SES.
	// Default: - Cognito built in email functionality.
	//
	EmailSendingAccount *string `field:"optional" json:"emailSendingAccount" yaml:"emailSendingAccount"`
	// Identifies either the sender's email address or the sender's name with their email address.
	//
	// If emailSendingAccount is DEVELOPER then this cannot be specified.
	// Default: 'no-reply@verificationemail.com'
	//
	From *string `field:"optional" json:"from" yaml:"from"`
	// The destination to which the receiver of the email should reply to.
	// Default: - same as `from`.
	//
	ReplyToEmailAddress *string `field:"optional" json:"replyToEmailAddress" yaml:"replyToEmailAddress"`
	// The ARN of a verified email address in Amazon SES.
	//
	// required if emailSendingAccount is DEVELOPER or if
	// 'from' is provided.
	// Default: - none.
	//
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

