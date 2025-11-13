package interfacesawsec2


// A reference to a TransitGatewayConnect resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayConnectReference := &TransitGatewayConnectReference{
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   }
//
type TransitGatewayConnectReference struct {
	// The TransitGatewayAttachmentId of the TransitGatewayConnect resource.
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
}

