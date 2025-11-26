package previewawswisdommixins


// The body to use in email messages.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   emailMessageTemplateContentBodyProperty := &EmailMessageTemplateContentBodyProperty{
//   	Html: &MessageTemplateBodyContentProviderProperty{
//   		Content: jsii.String("content"),
//   	},
//   	PlainText: &MessageTemplateBodyContentProviderProperty{
//   		Content: jsii.String("content"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-emailmessagetemplatecontentbody.html
//
type CfnMessageTemplatePropsMixin_EmailMessageTemplateContentBodyProperty struct {
	// The message body, in HTML format, to use in email messages that are based on the message template.
	//
	// We recommend using HTML format for email clients that render HTML content. You can include links, formatted text, and more in an HTML message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-emailmessagetemplatecontentbody.html#cfn-wisdom-messagetemplate-emailmessagetemplatecontentbody-html
	//
	Html interface{} `field:"optional" json:"html" yaml:"html"`
	// The message body, in plain text format, to use in email messages that are based on the message template.
	//
	// We recommend using plain text format for email clients that don't render HTML content and clients that are connected to high-latency networks, such as mobile devices.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-emailmessagetemplatecontentbody.html#cfn-wisdom-messagetemplate-emailmessagetemplatecontentbody-plaintext
	//
	PlainText interface{} `field:"optional" json:"plainText" yaml:"plainText"`
}

