package awswisdom


// Information about the message template attachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageTemplateAttachmentProperty := &MessageTemplateAttachmentProperty{
//   	AttachmentName: jsii.String("attachmentName"),
//   	S3PresignedUrl: jsii.String("s3PresignedUrl"),
//
//   	// the properties below are optional
//   	AttachmentId: jsii.String("attachmentId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-messagetemplateattachment.html
//
type CfnMessageTemplate_MessageTemplateAttachmentProperty struct {
	// The name of the attachment file being uploaded.
	//
	// The name should include the file extension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-messagetemplateattachment.html#cfn-wisdom-messagetemplate-messagetemplateattachment-attachmentname
	//
	AttachmentName *string `field:"required" json:"attachmentName" yaml:"attachmentName"`
	// The S3 Presigned URL for the attachment file.
	//
	// When generating the PreSignedUrl, please ensure that the expires-in time is set to 30 minutes. The URL can be generated through the AWS Console or through the AWS CLI (https://docs.aws.amazon.com/AmazonS3/latest/userguide/ShareObjectPreSignedURL.html).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-messagetemplateattachment.html#cfn-wisdom-messagetemplate-messagetemplateattachment-s3presignedurl
	//
	S3PresignedUrl *string `field:"required" json:"s3PresignedUrl" yaml:"s3PresignedUrl"`
	// The identifier of the attachment file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-messagetemplateattachment.html#cfn-wisdom-messagetemplate-messagetemplateattachment-attachmentid
	//
	AttachmentId *string `field:"optional" json:"attachmentId" yaml:"attachmentId"`
}

