package awsec2


// Properties for defining a `CfnClientVpnRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClientVpnRouteProps := &cfnClientVpnRouteProps{
//   	clientVpnEndpointId: jsii.String("clientVpnEndpointId"),
//   	destinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	targetVpcSubnetId: jsii.String("targetVpcSubnetId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnClientVpnRouteProps struct {
	// The ID of the Client VPN endpoint to which to add the route.
	ClientVpnEndpointId *string `field:"required" json:"clientVpnEndpointId" yaml:"clientVpnEndpointId"`
	// The IPv4 address range, in CIDR notation, of the route destination. For example:.
	//
	// - To add a route for Internet access, enter `0.0.0.0/0`
	// - To add a route for a peered VPC, enter the peered VPC's IPv4 CIDR range
	// - To add a route for an on-premises network, enter the AWS Site-to-Site VPN connection's IPv4 CIDR range
	// - To add a route for the local network, enter the client CIDR range.
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The ID of the subnet through which you want to route traffic.
	//
	// The specified subnet must be an existing target network of the Client VPN endpoint.
	//
	// Alternatively, if you're adding a route for the local network, specify `local` .
	TargetVpcSubnetId *string `field:"required" json:"targetVpcSubnetId" yaml:"targetVpcSubnetId"`
	// A brief description of the route.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

