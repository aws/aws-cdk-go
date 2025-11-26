package previewawswisdommixins


// The container of the message template body.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   messageTemplateBodyContentProviderProperty := &MessageTemplateBodyContentProviderProperty{
//   	Content: jsii.String("content"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-messagetemplatebodycontentprovider.html
//
type CfnMessageTemplatePropsMixin_MessageTemplateBodyContentProviderProperty struct {
	// The content of the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-messagetemplatebodycontentprovider.html#cfn-wisdom-messagetemplate-messagetemplatebodycontentprovider-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
}

