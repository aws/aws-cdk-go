package awsnetworkmanager


// A reference to a DirectConnectGatewayAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   directConnectGatewayAttachmentReference := &DirectConnectGatewayAttachmentReference{
//   	AttachmentId: jsii.String("attachmentId"),
//   }
//
type DirectConnectGatewayAttachmentReference struct {
	// The AttachmentId of the DirectConnectGatewayAttachment resource.
	AttachmentId *string `field:"required" json:"attachmentId" yaml:"attachmentId"`
}

