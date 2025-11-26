package previewawspinpointmixins


// Properties for CfnSmsTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   cfnSmsTemplateMixinProps := &CfnSmsTemplateMixinProps{
//   	Body: jsii.String("body"),
//   	DefaultSubstitutions: jsii.String("defaultSubstitutions"),
//   	Tags: tags,
//   	TemplateDescription: jsii.String("templateDescription"),
//   	TemplateName: jsii.String("templateName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-smstemplate.html
//
type CfnSmsTemplateMixinProps struct {
	// The message body to use in text messages that are based on the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-smstemplate.html#cfn-pinpoint-smstemplate-body
	//
	Body *string `field:"optional" json:"body" yaml:"body"`
	// A JSON object that specifies the default values to use for message variables in the message template.
	//
	// This object is a set of key-value pairs. Each key defines a message variable in the template. The corresponding value defines the default value for that variable. When you create a message that's based on the template, you can override these defaults with message-specific and address-specific variables and values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-smstemplate.html#cfn-pinpoint-smstemplate-defaultsubstitutions
	//
	DefaultSubstitutions *string `field:"optional" json:"defaultSubstitutions" yaml:"defaultSubstitutions"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-smstemplate.html#cfn-pinpoint-smstemplate-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A custom description of the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-smstemplate.html#cfn-pinpoint-smstemplate-templatedescription
	//
	TemplateDescription *string `field:"optional" json:"templateDescription" yaml:"templateDescription"`
	// The name of the message template to use for the message.
	//
	// If specified, this value must match the name of an existing message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-smstemplate.html#cfn-pinpoint-smstemplate-templatename
	//
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
}

