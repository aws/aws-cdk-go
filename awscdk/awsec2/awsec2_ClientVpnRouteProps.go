package awsec2


// Properties for a ClientVpnRoute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var clientVpnEndpoint clientVpnEndpoint
//   var clientVpnRouteTarget clientVpnRouteTarget
//
//   clientVpnRouteProps := &ClientVpnRouteProps{
//   	Cidr: jsii.String("cidr"),
//   	Target: clientVpnRouteTarget,
//
//   	// the properties below are optional
//   	ClientVpnEndoint: clientVpnEndpoint,
//   	ClientVpnEndpoint: clientVpnEndpoint,
//   	Description: jsii.String("description"),
//   }
//
// Experimental.
type ClientVpnRouteProps struct {
	// The IPv4 address range, in CIDR notation, of the route destination.
	//
	// For example:
	//    - To add a route for Internet access, enter 0.0.0.0/0
	//    - To add a route for a peered VPC, enter the peered VPC's IPv4 CIDR range
	//    - To add a route for an on-premises network, enter the AWS Site-to-Site VPN
	//      connection's IPv4 CIDR range
	// - To add a route for the local network, enter the client CIDR range.
	// Experimental.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// The target for the route.
	// Experimental.
	Target ClientVpnRouteTarget `field:"required" json:"target" yaml:"target"`
	// A brief description of the authorization rule.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The client VPN endpoint to which to add the route.
	// Deprecated: Use `clientVpnEndpoint` instead.
	ClientVpnEndoint IClientVpnEndpoint `field:"optional" json:"clientVpnEndoint" yaml:"clientVpnEndoint"`
	// The client VPN endpoint to which to add the route.
	// Experimental.
	ClientVpnEndpoint IClientVpnEndpoint `field:"optional" json:"clientVpnEndpoint" yaml:"clientVpnEndpoint"`
}

