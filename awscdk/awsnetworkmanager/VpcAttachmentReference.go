package awsnetworkmanager


// A reference to a VpcAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcAttachmentReference := &VpcAttachmentReference{
//   	AttachmentId: jsii.String("attachmentId"),
//   }
//
type VpcAttachmentReference struct {
	// The AttachmentId of the VpcAttachment resource.
	AttachmentId *string `field:"required" json:"attachmentId" yaml:"attachmentId"`
}

