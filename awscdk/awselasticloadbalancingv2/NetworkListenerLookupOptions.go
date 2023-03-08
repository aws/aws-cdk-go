package awselasticloadbalancingv2


// Options for looking up a network listener.
//
// Example:
//   listener := elbv2.NetworkListener_FromLookup(this, jsii.String("ALBListener"), &NetworkListenerLookupOptions{
//   	LoadBalancerTags: map[string]*string{
//   		"Cluster": jsii.String("MyClusterName"),
//   	},
//   	ListenerProtocol: elbv2.Protocol_TCP,
//   	ListenerPort: jsii.Number(12345),
//   })
//
type NetworkListenerLookupOptions struct {
	// Filter listeners by listener port.
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
	// Filter listeners by associated load balancer arn.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Filter listeners by associated load balancer tags.
	LoadBalancerTags *map[string]*string `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
	// Protocol of the listener port.
	ListenerProtocol Protocol `field:"optional" json:"listenerProtocol" yaml:"listenerProtocol"`
}

