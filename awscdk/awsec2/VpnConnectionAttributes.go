package awsec2


// Attributes of an imported VpnConnection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpnConnectionAttributes := &VpnConnectionAttributes{
//   	CustomerGatewayAsn: jsii.Number(123),
//   	CustomerGatewayId: jsii.String("customerGatewayId"),
//   	CustomerGatewayIp: jsii.String("customerGatewayIp"),
//   	VpnId: jsii.String("vpnId"),
//   }
//
// Experimental.
type VpnConnectionAttributes struct {
	// The ASN of the customer gateway.
	// Experimental.
	CustomerGatewayAsn *float64 `field:"required" json:"customerGatewayAsn" yaml:"customerGatewayAsn"`
	// The id of the customer gateway.
	// Experimental.
	CustomerGatewayId *string `field:"required" json:"customerGatewayId" yaml:"customerGatewayId"`
	// The ip address of the customer gateway.
	// Experimental.
	CustomerGatewayIp *string `field:"required" json:"customerGatewayIp" yaml:"customerGatewayIp"`
	// The id of the VPN connection.
	// Experimental.
	VpnId *string `field:"required" json:"vpnId" yaml:"vpnId"`
}

