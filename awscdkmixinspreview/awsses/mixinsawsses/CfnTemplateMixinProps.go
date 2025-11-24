package mixinsawsses


// Properties for CfnTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTemplateMixinProps := &CfnTemplateMixinProps{
//   	Template: &TemplateProperty{
//   		HtmlPart: jsii.String("htmlPart"),
//   		SubjectPart: jsii.String("subjectPart"),
//   		TemplateName: jsii.String("templateName"),
//   		TextPart: jsii.String("textPart"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-template.html
//
type CfnTemplateMixinProps struct {
	// The content of the email, composed of a subject line and either an HTML part or a text-only part.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-template.html#cfn-ses-template-template
	//
	Template interface{} `field:"optional" json:"template" yaml:"template"`
}

