package previewawswisdommixins


// The content of the message template that applies to the email channel subtype.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   emailMessageTemplateContentProperty := &EmailMessageTemplateContentProperty{
//   	Body: &EmailMessageTemplateContentBodyProperty{
//   		Html: &MessageTemplateBodyContentProviderProperty{
//   			Content: jsii.String("content"),
//   		},
//   		PlainText: &MessageTemplateBodyContentProviderProperty{
//   			Content: jsii.String("content"),
//   		},
//   	},
//   	Headers: []interface{}{
//   		&EmailMessageTemplateHeaderProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Subject: jsii.String("subject"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-emailmessagetemplatecontent.html
//
type CfnMessageTemplatePropsMixin_EmailMessageTemplateContentProperty struct {
	// The body to use in email messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-emailmessagetemplatecontent.html#cfn-wisdom-messagetemplate-emailmessagetemplatecontent-body
	//
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	// The email headers to include in email messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-emailmessagetemplatecontent.html#cfn-wisdom-messagetemplate-emailmessagetemplatecontent-headers
	//
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// The subject line, or title, to use in email messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-emailmessagetemplatecontent.html#cfn-wisdom-messagetemplate-emailmessagetemplatecontent-subject
	//
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

