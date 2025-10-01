package awsec2


// A reference to a NetworkInterfaceAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfaceAttachmentReference := &NetworkInterfaceAttachmentReference{
//   	AttachmentId: jsii.String("attachmentId"),
//   }
//
type NetworkInterfaceAttachmentReference struct {
	// The AttachmentId of the NetworkInterfaceAttachment resource.
	AttachmentId *string `field:"required" json:"attachmentId" yaml:"attachmentId"`
}

