package awswisdom


// Properties for defining a `CfnMessageTemplateVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMessageTemplateVersionProps := &CfnMessageTemplateVersionProps{
//   	MessageTemplateArn: jsii.String("messageTemplateArn"),
//
//   	// the properties below are optional
//   	MessageTemplateContentSha256: jsii.String("messageTemplateContentSha256"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplateversion.html
//
type CfnMessageTemplateVersionProps struct {
	// The unqualified Amazon Resource Name (ARN) of the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplateversion.html#cfn-wisdom-messagetemplateversion-messagetemplatearn
	//
	MessageTemplateArn *string `field:"required" json:"messageTemplateArn" yaml:"messageTemplateArn"`
	// The content SHA256 of the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplateversion.html#cfn-wisdom-messagetemplateversion-messagetemplatecontentsha256
	//
	MessageTemplateContentSha256 *string `field:"optional" json:"messageTemplateContentSha256" yaml:"messageTemplateContentSha256"`
}

