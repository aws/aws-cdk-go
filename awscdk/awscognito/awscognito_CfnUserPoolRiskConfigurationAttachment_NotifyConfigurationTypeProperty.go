package awscognito


// The notify configuration type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notifyConfigurationTypeProperty := &notifyConfigurationTypeProperty{
//   	sourceArn: jsii.String("sourceArn"),
//
//   	// the properties below are optional
//   	blockEmail: &notifyEmailTypeProperty{
//   		subject: jsii.String("subject"),
//
//   		// the properties below are optional
//   		htmlBody: jsii.String("htmlBody"),
//   		textBody: jsii.String("textBody"),
//   	},
//   	from: jsii.String("from"),
//   	mfaEmail: &notifyEmailTypeProperty{
//   		subject: jsii.String("subject"),
//
//   		// the properties below are optional
//   		htmlBody: jsii.String("htmlBody"),
//   		textBody: jsii.String("textBody"),
//   	},
//   	noActionEmail: &notifyEmailTypeProperty{
//   		subject: jsii.String("subject"),
//
//   		// the properties below are optional
//   		htmlBody: jsii.String("htmlBody"),
//   		textBody: jsii.String("textBody"),
//   	},
//   	replyTo: jsii.String("replyTo"),
//   }
//
type CfnUserPoolRiskConfigurationAttachment_NotifyConfigurationTypeProperty struct {
	// The Amazon Resource Name (ARN) of the identity that is associated with the sending authorization policy.
	//
	// This identity permits Amazon Cognito to send for the email address specified in the `From` parameter.
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// Email template used when a detected risk event is blocked.
	BlockEmail interface{} `field:"optional" json:"blockEmail" yaml:"blockEmail"`
	// The email address that is sending the email.
	//
	// The address must be either individually verified with Amazon Simple Email Service, or from a domain that has been verified with Amazon SES.
	From *string `field:"optional" json:"from" yaml:"from"`
	// The multi-factor authentication (MFA) email template used when MFA is challenged as part of a detected risk.
	MfaEmail interface{} `field:"optional" json:"mfaEmail" yaml:"mfaEmail"`
	// The email template used when a detected risk event is allowed.
	NoActionEmail interface{} `field:"optional" json:"noActionEmail" yaml:"noActionEmail"`
	// The destination to which the receiver of an email should reply to.
	ReplyTo *string `field:"optional" json:"replyTo" yaml:"replyTo"`
}

