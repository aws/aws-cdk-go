package awselasticloadbalancingv2


// Properties for a redirection config.
//
// Example:
//   var lb applicationLoadBalancer
//
//
//   lb.addRedirect(&applicationLoadBalancerRedirectConfig{
//   	sourceProtocol: elbv2.applicationProtocol_HTTPS,
//   	sourcePort: jsii.Number(8443),
//   	targetProtocol: elbv2.*applicationProtocol_HTTP,
//   	targetPort: jsii.Number(8080),
//   })
//
type ApplicationLoadBalancerRedirectConfig struct {
	// Allow anyone to connect to this listener.
	//
	// If this is specified, the listener will be opened up to anyone who can reach it.
	// For internal load balancers this is anyone in the same VPC. For public load
	// balancers, this is anyone on the internet.
	//
	// If you want to be more selective about who can access this load
	// balancer, set this to `false` and use the listener's `connections`
	// object to selectively grant access to the listener.
	Open *bool `field:"optional" json:"open" yaml:"open"`
	// The port number to listen to.
	SourcePort *float64 `field:"optional" json:"sourcePort" yaml:"sourcePort"`
	// The protocol of the listener being created.
	SourceProtocol ApplicationProtocol `field:"optional" json:"sourceProtocol" yaml:"sourceProtocol"`
	// The port number to redirect to.
	TargetPort *float64 `field:"optional" json:"targetPort" yaml:"targetPort"`
	// The protocol of the redirection target.
	TargetProtocol ApplicationProtocol `field:"optional" json:"targetProtocol" yaml:"targetProtocol"`
}

