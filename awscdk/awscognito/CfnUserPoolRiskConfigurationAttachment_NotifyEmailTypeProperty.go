package awscognito


// The notify email type.
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
	// The email subject.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyemailtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyemailtype-subject
	//
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// The email HTML body.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyemailtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyemailtype-htmlbody
	//
	HtmlBody *string `field:"optional" json:"htmlBody" yaml:"htmlBody"`
	// The email text body.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyemailtype.html#cfn-cognito-userpoolriskconfigurationattachment-notifyemailtype-textbody
	//
	TextBody *string `field:"optional" json:"textBody" yaml:"textBody"`
}

