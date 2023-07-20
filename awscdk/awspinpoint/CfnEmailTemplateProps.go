package awspinpoint


// Properties for defining a `CfnEmailTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnEmailTemplateProps := &CfnEmailTemplateProps{
//   	Subject: jsii.String("subject"),
//   	TemplateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	DefaultSubstitutions: jsii.String("defaultSubstitutions"),
//   	HtmlPart: jsii.String("htmlPart"),
//   	Tags: tags,
//   	TemplateDescription: jsii.String("templateDescription"),
//   	TextPart: jsii.String("textPart"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailtemplate.html
//
type CfnEmailTemplateProps struct {
	// The subject line, or title, to use in email messages that are based on the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailtemplate.html#cfn-pinpoint-emailtemplate-subject
	//
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// The name of the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailtemplate.html#cfn-pinpoint-emailtemplate-templatename
	//
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// A JSON object that specifies the default values to use for message variables in the message template.
	//
	// This object is a set of key-value pairs. Each key defines a message variable in the template. The corresponding value defines the default value for that variable. When you create a message that's based on the template, you can override these defaults with message-specific and address-specific variables and values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailtemplate.html#cfn-pinpoint-emailtemplate-defaultsubstitutions
	//
	DefaultSubstitutions *string `field:"optional" json:"defaultSubstitutions" yaml:"defaultSubstitutions"`
	// The message body, in HTML format, to use in email messages that are based on the message template.
	//
	// We recommend using HTML format for email clients that render HTML content. You can include links, formatted text, and more in an HTML message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailtemplate.html#cfn-pinpoint-emailtemplate-htmlpart
	//
	HtmlPart *string `field:"optional" json:"htmlPart" yaml:"htmlPart"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailtemplate.html#cfn-pinpoint-emailtemplate-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A custom description of the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailtemplate.html#cfn-pinpoint-emailtemplate-templatedescription
	//
	TemplateDescription *string `field:"optional" json:"templateDescription" yaml:"templateDescription"`
	// The message body, in plain text format, to use in email messages that are based on the message template.
	//
	// We recommend using plain text format for email clients that don't render HTML content and clients that are connected to high-latency networks, such as mobile devices.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailtemplate.html#cfn-pinpoint-emailtemplate-textpart
	//
	TextPart *string `field:"optional" json:"textPart" yaml:"textPart"`
}

