package awswisdom


// The content of the message template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentProperty := &ContentProperty{
//   	EmailMessageTemplateContent: &EmailMessageTemplateContentProperty{
//   		Body: &EmailMessageTemplateContentBodyProperty{
//   			Html: &MessageTemplateBodyContentProviderProperty{
//   				Content: jsii.String("content"),
//   			},
//   			PlainText: &MessageTemplateBodyContentProviderProperty{
//   				Content: jsii.String("content"),
//   			},
//   		},
//   		Headers: []interface{}{
//   			&EmailMessageTemplateHeaderProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Subject: jsii.String("subject"),
//   	},
//   	SmsMessageTemplateContent: &SmsMessageTemplateContentProperty{
//   		Body: &SmsMessageTemplateContentBodyProperty{
//   			PlainText: &MessageTemplateBodyContentProviderProperty{
//   				Content: jsii.String("content"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-content.html
//
type CfnMessageTemplate_ContentProperty struct {
	// The content of the message template that applies to the email channel subtype.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-content.html#cfn-wisdom-messagetemplate-content-emailmessagetemplatecontent
	//
	EmailMessageTemplateContent interface{} `field:"optional" json:"emailMessageTemplateContent" yaml:"emailMessageTemplateContent"`
	// The content of message template that applies to SMS channel subtype.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-content.html#cfn-wisdom-messagetemplate-content-smsmessagetemplatecontent
	//
	SmsMessageTemplateContent interface{} `field:"optional" json:"smsMessageTemplateContent" yaml:"smsMessageTemplateContent"`
}

