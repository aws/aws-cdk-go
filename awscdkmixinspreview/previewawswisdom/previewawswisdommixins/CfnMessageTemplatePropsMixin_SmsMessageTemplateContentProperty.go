package previewawswisdommixins


// The content of the message template that applies to the SMS channel subtype.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   smsMessageTemplateContentProperty := &SmsMessageTemplateContentProperty{
//   	Body: &SmsMessageTemplateContentBodyProperty{
//   		PlainText: &MessageTemplateBodyContentProviderProperty{
//   			Content: jsii.String("content"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-smsmessagetemplatecontent.html
//
type CfnMessageTemplatePropsMixin_SmsMessageTemplateContentProperty struct {
	// The body to use in SMS messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-smsmessagetemplatecontent.html#cfn-wisdom-messagetemplate-smsmessagetemplatecontent-body
	//
	Body interface{} `field:"optional" json:"body" yaml:"body"`
}

