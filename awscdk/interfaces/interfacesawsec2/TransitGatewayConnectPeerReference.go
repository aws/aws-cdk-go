package interfacesawsec2


// A reference to a TransitGatewayConnectPeer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayConnectPeerReference := &TransitGatewayConnectPeerReference{
//   	TransitGatewayConnectPeerId: jsii.String("transitGatewayConnectPeerId"),
//   }
//
type TransitGatewayConnectPeerReference struct {
	// The TransitGatewayConnectPeerId of the TransitGatewayConnectPeer resource.
	TransitGatewayConnectPeerId *string `field:"required" json:"transitGatewayConnectPeerId" yaml:"transitGatewayConnectPeerId"`
}

