package awsec2


// Options for a ClientVpnRoute.
//
// Example:
//   endpoint := vpc.addClientVpnEndpoint(jsii.String("Endpoint"), &ClientVpnEndpointOptions{
//   	Cidr: jsii.String("10.100.0.0/16"),
//   	ServerCertificateArn: jsii.String("arn:aws:acm:us-east-1:123456789012:certificate/server-certificate-id"),
//   	UserBasedAuthentication: ec2.ClientVpnUserBasedAuthentication_Federated(samlProvider),
//   })
//
//   // Client-to-client access
//   endpoint.AddRoute(jsii.String("Route"), &ClientVpnRouteOptions{
//   	Cidr: jsii.String("10.100.0.0/16"),
//   	Target: ec2.ClientVpnRouteTarget_Local(),
//   })
//
type ClientVpnRouteOptions struct {
	// The IPv4 address range, in CIDR notation, of the route destination.
	//
	// For example:
	//   - To add a route for Internet access, enter 0.0.0.0/0
	//   - To add a route for a peered VPC, enter the peered VPC's IPv4 CIDR range
	//   - To add a route for an on-premises network, enter the AWS Site-to-Site VPN
	//     connection's IPv4 CIDR range
	// - To add a route for the local network, enter the client CIDR range.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// The target for the route.
	Target ClientVpnRouteTarget `field:"required" json:"target" yaml:"target"`
	// A brief description of the authorization rule.
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

