package previewawswisdommixins


// Properties for CfnMessageTemplateVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMessageTemplateVersionMixinProps := &CfnMessageTemplateVersionMixinProps{
//   	MessageTemplateArn: jsii.String("messageTemplateArn"),
//   	MessageTemplateContentSha256: jsii.String("messageTemplateContentSha256"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplateversion.html
//
type CfnMessageTemplateVersionMixinProps struct {
	// The Amazon Resource Name (ARN) of the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplateversion.html#cfn-wisdom-messagetemplateversion-messagetemplatearn
	//
	MessageTemplateArn *string `field:"optional" json:"messageTemplateArn" yaml:"messageTemplateArn"`
	// The content SHA256 of the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplateversion.html#cfn-wisdom-messagetemplateversion-messagetemplatecontentsha256
	//
	MessageTemplateContentSha256 *string `field:"optional" json:"messageTemplateContentSha256" yaml:"messageTemplateContentSha256"`
}

