package interfacesawsec2


// A reference to a VPCGatewayAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCGatewayAttachmentReference := &VPCGatewayAttachmentReference{
//   	AttachmentType: jsii.String("attachmentType"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
type VPCGatewayAttachmentReference struct {
	// The AttachmentType of the VPCGatewayAttachment resource.
	AttachmentType *string `field:"required" json:"attachmentType" yaml:"attachmentType"`
	// The VpcId of the VPCGatewayAttachment resource.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

