package awscognito


// The template for email messages that threat protection sends to a user when your threat protection automated response has a *Notify* action.
//
// This data type is a request parameter of `API_SetRiskConfiguration` and a response parameter of `API_DescribeRiskConfiguration` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notifyEmailTypeProperty := &NotifyEmailTypeProperty{
//   	Subject: jsii.String("subject"),
//
//   	// the properties below are optional
//   	HtmlBody: jsii.String("htmlBody"),
//   	TextBody: jsii.String("textBody"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyemailtype.html
//
type CfnUserPoolRiskConfigurationAttachment_NotifyEmailTypeProperty struct {
	// The subject of the threat protection email notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyemailtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyemailtype-subject
	//
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// The body of an email notification formatted in HTML.
	//
	// Choose an `HtmlBody` or a `TextBody` to send an HTML-formatted or plaintext message, respectively.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyemailtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyemailtype-htmlbody
	//
	HtmlBody *string `field:"optional" json:"htmlBody" yaml:"htmlBody"`
	// The body of an email notification formatted in plaintext.
	//
	// Choose an `HtmlBody` or a `TextBody` to send an HTML-formatted or plaintext message, respectively.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyemailtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyemailtype-textbody
	//
	TextBody *string `field:"optional" json:"textBody" yaml:"textBody"`
}

