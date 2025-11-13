package interfacesawsnetworkmanager


// A reference to a TransitGatewayPeering resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayPeeringReference := &TransitGatewayPeeringReference{
//   	PeeringId: jsii.String("peeringId"),
//   }
//
type TransitGatewayPeeringReference struct {
	// The PeeringId of the TransitGatewayPeering resource.
	PeeringId *string `field:"required" json:"peeringId" yaml:"peeringId"`
}

