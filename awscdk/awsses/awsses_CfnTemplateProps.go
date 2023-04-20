package awsses


// Properties for defining a `CfnTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTemplateProps := &CfnTemplateProps{
//   	Template: &TemplateProperty{
//   		SubjectPart: jsii.String("subjectPart"),
//
//   		// the properties below are optional
//   		HtmlPart: jsii.String("htmlPart"),
//   		TemplateName: jsii.String("templateName"),
//   		TextPart: jsii.String("textPart"),
//   	},
//   }
//
type CfnTemplateProps struct {
	// The content of the email, composed of a subject line and either an HTML part or a text-only part.
	Template interface{} `field:"optional" json:"template" yaml:"template"`
}

