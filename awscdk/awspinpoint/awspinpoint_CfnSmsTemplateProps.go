package awspinpoint


// Properties for defining a `CfnSmsTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnSmsTemplateProps := &cfnSmsTemplateProps{
//   	body: jsii.String("body"),
//   	templateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	defaultSubstitutions: jsii.String("defaultSubstitutions"),
//   	tags: tags,
//   	templateDescription: jsii.String("templateDescription"),
//   }
//
type CfnSmsTemplateProps struct {
	// The message body to use in text messages that are based on the message template.
	Body *string `field:"required" json:"body" yaml:"body"`
	// The name of the message template.
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// A JSON object that specifies the default values to use for message variables in the message template.
	//
	// This object is a set of key-value pairs. Each key defines a message variable in the template. The corresponding value defines the default value for that variable. When you create a message that's based on the template, you can override these defaults with message-specific and address-specific variables and values.
	DefaultSubstitutions *string `field:"optional" json:"defaultSubstitutions" yaml:"defaultSubstitutions"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A custom description of the message template.
	TemplateDescription *string `field:"optional" json:"templateDescription" yaml:"templateDescription"`
}

