package mixinsawsses


// An object that defines the email template to use for an email message, and the values to use for any message variables in that template.
//
// An *email template* is a type of message template that contains content that you want to reuse in email messages that you send. You can specifiy the email template by providing the name or ARN of an *email template* previously saved in your Amazon SES account or by providing the full template content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   templateProperty := &TemplateProperty{
//   	HtmlPart: jsii.String("htmlPart"),
//   	SubjectPart: jsii.String("subjectPart"),
//   	TemplateName: jsii.String("templateName"),
//   	TextPart: jsii.String("textPart"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html
//
type CfnTemplatePropsMixin_TemplateProperty struct {
	// The HTML body of the email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-htmlpart
	//
	HtmlPart *string `field:"optional" json:"htmlPart" yaml:"htmlPart"`
	// The subject line of the email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-subjectpart
	//
	SubjectPart *string `field:"optional" json:"subjectPart" yaml:"subjectPart"`
	// The name of the template.
	//
	// You will refer to this name when you send email using the `SendEmail` or `SendBulkEmail` operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-templatename
	//
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// The email body that is visible to recipients whose email clients do not display HTML content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-textpart
	//
	TextPart *string `field:"optional" json:"textPart" yaml:"textPart"`
}

