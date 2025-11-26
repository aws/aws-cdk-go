package previewawswisdommixins


// The body to use in SMS messages.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   smsMessageTemplateContentBodyProperty := &SmsMessageTemplateContentBodyProperty{
//   	PlainText: &MessageTemplateBodyContentProviderProperty{
//   		Content: jsii.String("content"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-smsmessagetemplatecontentbody.html
//
type CfnMessageTemplatePropsMixin_SmsMessageTemplateContentBodyProperty struct {
	// The message body to use in SMS messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-smsmessagetemplatecontentbody.html#cfn-wisdom-messagetemplate-smsmessagetemplatecontentbody-plaintext
	//
	PlainText interface{} `field:"optional" json:"plainText" yaml:"plainText"`
}

