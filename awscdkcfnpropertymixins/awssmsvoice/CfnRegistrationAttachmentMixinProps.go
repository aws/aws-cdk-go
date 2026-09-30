package awssmsvoice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRegistrationAttachmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnRegistrationAttachmentMixinProps := &CfnRegistrationAttachmentMixinProps{
//   	AttachmentBody: jsii.String("attachmentBody"),
//   	AttachmentUrl: jsii.String("attachmentUrl"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-registrationattachment.html
//
type CfnRegistrationAttachmentMixinProps struct {
	// The registration file to upload.
	//
	// The maximum file size is 1500KB and valid file extensions are PDF, JPEG and PNG.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-registrationattachment.html#cfn-smsvoice-registrationattachment-attachmentbody
	//
	AttachmentBody *string `field:"optional" json:"attachmentBody" yaml:"attachmentBody"`
	// A URL to the required registration file.
	//
	// For example, the URL to an MMS/shortcode form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-registrationattachment.html#cfn-smsvoice-registrationattachment-attachmenturl
	//
	AttachmentUrl *string `field:"optional" json:"attachmentUrl" yaml:"attachmentUrl"`
	// An array of tags (key and value pairs) to associate with the registration attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-registrationattachment.html#cfn-smsvoice-registrationattachment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

