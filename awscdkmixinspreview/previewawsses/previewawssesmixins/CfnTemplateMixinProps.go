package previewawssesmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTemplateMixinProps := &CfnTemplateMixinProps{
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
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
	// The tags (keys and values) associated with the email template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-template.html#cfn-ses-template-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The content of the email, composed of a subject line and either an HTML part or a text-only part.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-template.html#cfn-ses-template-template
	//
	Template interface{} `field:"optional" json:"template" yaml:"template"`
}

