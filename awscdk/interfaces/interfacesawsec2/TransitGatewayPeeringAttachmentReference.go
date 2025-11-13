package interfacesawsec2


// A reference to a TransitGatewayPeeringAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayPeeringAttachmentReference := &TransitGatewayPeeringAttachmentReference{
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   }
//
type TransitGatewayPeeringAttachmentReference struct {
	// The TransitGatewayAttachmentId of the TransitGatewayPeeringAttachment resource.
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
}

