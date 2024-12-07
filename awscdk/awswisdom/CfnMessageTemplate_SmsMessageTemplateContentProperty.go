package awswisdom


// The content of message template that applies to SMS channel subtype.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnMessageTemplate_SmsMessageTemplateContentProperty struct {
	// The body to use in SMS messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-smsmessagetemplatecontent.html#cfn-wisdom-messagetemplate-smsmessagetemplatecontent-body
	//
	Body interface{} `field:"required" json:"body" yaml:"body"`
}

