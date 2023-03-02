package awsses


// The content of the email, composed of a subject line and either an HTML part or a text-only part.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateProperty := &templateProperty{
//   	subjectPart: jsii.String("subjectPart"),
//
//   	// the properties below are optional
//   	htmlPart: jsii.String("htmlPart"),
//   	templateName: jsii.String("templateName"),
//   	textPart: jsii.String("textPart"),
//   }
//
type CfnTemplate_TemplateProperty struct {
	// The subject line of the email.
	SubjectPart *string `field:"required" json:"subjectPart" yaml:"subjectPart"`
	// The HTML body of the email.
	HtmlPart *string `field:"optional" json:"htmlPart" yaml:"htmlPart"`
	// The name of the template.
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// The email body that is visible to recipients whose email clients do not display HTML content.
	TextPart *string `field:"optional" json:"textPart" yaml:"textPart"`
}

