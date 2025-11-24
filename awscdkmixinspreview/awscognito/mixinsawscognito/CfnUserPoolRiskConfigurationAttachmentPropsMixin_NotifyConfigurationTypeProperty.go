package mixinsawscognito


// The configuration for Amazon SES email messages that advanced security features sends to a user when your adaptive authentication automated response has a *Notify* action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   notifyConfigurationTypeProperty := &NotifyConfigurationTypeProperty{
//   	BlockEmail: &NotifyEmailTypeProperty{
//   		HtmlBody: jsii.String("htmlBody"),
//   		Subject: jsii.String("subject"),
//   		TextBody: jsii.String("textBody"),
//   	},
//   	From: jsii.String("from"),
//   	MfaEmail: &NotifyEmailTypeProperty{
//   		HtmlBody: jsii.String("htmlBody"),
//   		Subject: jsii.String("subject"),
//   		TextBody: jsii.String("textBody"),
//   	},
//   	NoActionEmail: &NotifyEmailTypeProperty{
//   		HtmlBody: jsii.String("htmlBody"),
//   		Subject: jsii.String("subject"),
//   		TextBody: jsii.String("textBody"),
//   	},
//   	ReplyTo: jsii.String("replyTo"),
//   	SourceArn: jsii.String("sourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype.html
//
type CfnUserPoolRiskConfigurationAttachmentPropsMixin_NotifyConfigurationTypeProperty struct {
	// The template for the email message that your user pool sends when a detected risk event is blocked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype-blockemail
	//
	BlockEmail interface{} `field:"optional" json:"blockEmail" yaml:"blockEmail"`
	// The email address that sends the email message.
	//
	// The address must be either individually verified with Amazon Simple Email Service, or from a domain that has been verified with Amazon SES.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype-from
	//
	From *string `field:"optional" json:"from" yaml:"from"`
	// The template for the email message that your user pool sends when MFA is challenged in response to a detected risk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype-mfaemail
	//
	MfaEmail interface{} `field:"optional" json:"mfaEmail" yaml:"mfaEmail"`
	// The template for the email message that your user pool sends when no action is taken in response to a detected risk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype-noactionemail
	//
	NoActionEmail interface{} `field:"optional" json:"noActionEmail" yaml:"noActionEmail"`
	// The reply-to email address of an email template.
	//
	// Can be an email address in the format `admin@example.com` or `Administrator <admin@example.com>` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype-replyto
	//
	ReplyTo *string `field:"optional" json:"replyTo" yaml:"replyTo"`
	// The Amazon Resource Name (ARN) of the identity that is associated with the sending authorization policy.
	//
	// This identity permits Amazon Cognito to send for the email address specified in the `From` parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype-sourcearn
	//
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

