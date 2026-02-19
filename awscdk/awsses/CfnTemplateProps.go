package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTemplateProps := &CfnTemplateProps{
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-template.html
//
type CfnTemplateProps struct {
	// The tags (keys and values) associated with the email template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-template.html#cfn-ses-template-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The content of the email, composed of a subject line and either an HTML part or a text-only part.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-template.html#cfn-ses-template-template
	//
	Template interface{} `field:"optional" json:"template" yaml:"template"`
}

