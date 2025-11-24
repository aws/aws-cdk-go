package mixinsawscognito


// The template for email messages that advanced security features sends to a user when your threat protection automated response has a *Notify* action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   notifyEmailTypeProperty := &NotifyEmailTypeProperty{
//   	HtmlBody: jsii.String("htmlBody"),
//   	Subject: jsii.String("subject"),
//   	TextBody: jsii.String("textBody"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyemailtype.html
//
type CfnUserPoolRiskConfigurationAttachmentPropsMixin_NotifyEmailTypeProperty struct {
	// The body of an email notification formatted in HTML.
	//
	// Choose an `HtmlBody` or a `TextBody` to send an HTML-formatted or plaintext message, respectively.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyemailtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyemailtype-htmlbody
	//
	HtmlBody *string `field:"optional" json:"htmlBody" yaml:"htmlBody"`
	// The subject of the threat protection email notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyemailtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyemailtype-subject
	//
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
	// The body of an email notification formatted in plaintext.
	//
	// Choose an `HtmlBody` or a `TextBody` to send an HTML-formatted or plaintext message, respectively.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyemailtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyemailtype-textbody
	//
	TextBody *string `field:"optional" json:"textBody" yaml:"textBody"`
}

