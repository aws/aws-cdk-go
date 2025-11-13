package interfacesawsnetworkmanager


// A reference to a CustomerGatewayAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerGatewayAssociationReference := &CustomerGatewayAssociationReference{
//   	CustomerGatewayArn: jsii.String("customerGatewayArn"),
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   }
//
type CustomerGatewayAssociationReference struct {
	// The CustomerGatewayArn of the CustomerGatewayAssociation resource.
	CustomerGatewayArn *string `field:"required" json:"customerGatewayArn" yaml:"customerGatewayArn"`
	// The GlobalNetworkId of the CustomerGatewayAssociation resource.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
}

