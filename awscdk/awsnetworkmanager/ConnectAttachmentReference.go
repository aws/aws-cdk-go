package awsnetworkmanager


// A reference to a ConnectAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectAttachmentReference := &ConnectAttachmentReference{
//   	AttachmentId: jsii.String("attachmentId"),
//   }
//
type ConnectAttachmentReference struct {
	// The AttachmentId of the ConnectAttachment resource.
	AttachmentId *string `field:"required" json:"attachmentId" yaml:"attachmentId"`
}

