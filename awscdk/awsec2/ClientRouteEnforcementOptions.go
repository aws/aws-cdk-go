package awsec2


// Options for Client Route Enforcement.
//
// Example:
//   endpoint := vpc.addClientVpnEndpoint(jsii.String("Endpoint"), &ClientVpnEndpointOptions{
//   	Cidr: jsii.String("10.100.0.0/16"),
//   	ServerCertificateArn: jsii.String("arn:aws:acm:us-east-1:123456789012:certificate/server-certificate-id"),
//   	ClientCertificateArn: jsii.String("arn:aws:acm:us-east-1:123456789012:certificate/client-certificate-id"),
//   	ClientRouteEnforcementOptions: &ClientRouteEnforcementOptions{
//   		Enforced: jsii.Boolean(true),
//   	},
//   })
//
type ClientRouteEnforcementOptions struct {
	// Enable or disable Client Route Enforcement.
	//
	// The state can either be true (enabled) or false (disabled).
	Enforced *bool `field:"required" json:"enforced" yaml:"enforced"`
}

