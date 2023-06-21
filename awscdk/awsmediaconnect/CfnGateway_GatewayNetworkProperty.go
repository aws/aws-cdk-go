package awsmediaconnect


// The network settings for a gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayNetworkProperty := &GatewayNetworkProperty{
//   	CidrBlock: jsii.String("cidrBlock"),
//   	Name: jsii.String("name"),
//   }
//
type CfnGateway_GatewayNetworkProperty struct {
	// A unique IP address range to use for this network.
	//
	// These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// The name of the network.
	//
	// This name is used to reference the network and must be unique among networks in this gateway.
	Name *string `field:"required" json:"name" yaml:"name"`
}

