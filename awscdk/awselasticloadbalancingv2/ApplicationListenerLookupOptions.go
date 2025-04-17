package awselasticloadbalancingv2


// Options for ApplicationListener lookup.
//
// Example:
//   listener := elbv2.ApplicationListener_FromLookup(this, jsii.String("ALBListener"), &ApplicationListenerLookupOptions{
//   	LoadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:us-east-2:123456789012:loadbalancer/app/my-load-balancer/1234567890123456"),
//   	ListenerProtocol: elbv2.ApplicationProtocol_HTTPS,
//   	ListenerPort: jsii.Number(443),
//   })
//
type ApplicationListenerLookupOptions struct {
	// Filter listeners by listener port.
	// Default: - does not filter by listener port.
	//
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
	// Filter listeners by associated load balancer arn.
	// Default: - does not filter by load balancer arn.
	//
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Filter listeners by associated load balancer tags.
	// Default: - does not filter by load balancer tags.
	//
	LoadBalancerTags *map[string]*string `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
	// ARN of the listener to look up.
	// Default: - does not filter by listener arn.
	//
	ListenerArn *string `field:"optional" json:"listenerArn" yaml:"listenerArn"`
	// Filter listeners by listener protocol.
	// Default: - does not filter by listener protocol.
	//
	ListenerProtocol ApplicationProtocol `field:"optional" json:"listenerProtocol" yaml:"listenerProtocol"`
}

