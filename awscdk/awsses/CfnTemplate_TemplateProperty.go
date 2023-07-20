package awsses


// The content of the email, composed of a subject line and either an HTML part or a text-only part.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateProperty := &TemplateProperty{
//   	SubjectPart: jsii.String("subjectPart"),
//
//   	// the properties below are optional
//   	HtmlPart: jsii.String("htmlPart"),
//   	TemplateName: jsii.String("templateName"),
//   	TextPart: jsii.String("textPart"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html
//
type CfnTemplate_TemplateProperty struct {
	// The subject line of the email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-subjectpart
	//
	SubjectPart *string `field:"required" json:"subjectPart" yaml:"subjectPart"`
	// The HTML body of the email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-htmlpart
	//
	HtmlPart *string `field:"optional" json:"htmlPart" yaml:"htmlPart"`
	// The name of the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-templatename
	//
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// The email body that is visible to recipients whose email clients do not display HTML content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-textpart
	//
	TextPart *string `field:"optional" json:"textPart" yaml:"textPart"`
}

