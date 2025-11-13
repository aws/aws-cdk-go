package interfacesawsec2


// A reference to a TransitGatewayVpcAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayVpcAttachmentReference := &TransitGatewayVpcAttachmentReference{
//   	TransitGatewayVpcAttachmentId: jsii.String("transitGatewayVpcAttachmentId"),
//   }
//
type TransitGatewayVpcAttachmentReference struct {
	// The Id of the TransitGatewayVpcAttachment resource.
	TransitGatewayVpcAttachmentId *string `field:"required" json:"transitGatewayVpcAttachmentId" yaml:"transitGatewayVpcAttachmentId"`
}

