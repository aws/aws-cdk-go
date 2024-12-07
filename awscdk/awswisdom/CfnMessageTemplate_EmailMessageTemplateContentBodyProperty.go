package awswisdom


// The body to use in email messages.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnMessageTemplate_EmailMessageTemplateContentBodyProperty struct {
	// The container of message template body.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-emailmessagetemplatecontentbody.html#cfn-wisdom-messagetemplate-emailmessagetemplatecontentbody-html
	//
	Html interface{} `field:"optional" json:"html" yaml:"html"`
	// The container of message template body.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-emailmessagetemplatecontentbody.html#cfn-wisdom-messagetemplate-emailmessagetemplatecontentbody-plaintext
	//
	PlainText interface{} `field:"optional" json:"plainText" yaml:"plainText"`
}

