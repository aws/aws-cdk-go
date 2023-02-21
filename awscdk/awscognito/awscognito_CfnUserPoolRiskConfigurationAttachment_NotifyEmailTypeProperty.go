package awscognito


// The notify email type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notifyEmailTypeProperty := &notifyEmailTypeProperty{
//   	subject: jsii.String("subject"),
//
//   	// the properties below are optional
//   	htmlBody: jsii.String("htmlBody"),
//   	textBody: jsii.String("textBody"),
//   }
//
type CfnUserPoolRiskConfigurationAttachment_NotifyEmailTypeProperty struct {
	// The email subject.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// The email HTML body.
	HtmlBody *string `field:"optional" json:"htmlBody" yaml:"htmlBody"`
	// The email text body.
	TextBody *string `field:"optional" json:"textBody" yaml:"textBody"`
}

