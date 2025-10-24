package awsglobalaccelerator


// Construct options for Listener.
//
// Example:
//   // Create an Accelerator
//   accelerator := globalaccelerator.NewAccelerator(this, jsii.String("Accelerator"))
//
//   // Create a Listener
//   listener := accelerator.AddListener(jsii.String("Listener"), &ListenerOptions{
//   	PortRanges: []PortRange{
//   		&PortRange{
//   			FromPort: jsii.Number(80),
//   		},
//   		&PortRange{
//   			FromPort: jsii.Number(443),
//   		},
//   	},
//   })
//
//   // Import the Load Balancers
//   nlb1 := elbv2.NetworkLoadBalancer_FromNetworkLoadBalancerAttributes(this, jsii.String("NLB1"), &NetworkLoadBalancerAttributes{
//   	LoadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:us-west-2:111111111111:loadbalancer/app/my-load-balancer1/e16bef66805b"),
//   })
//   nlb2 := elbv2.NetworkLoadBalancer_FromNetworkLoadBalancerAttributes(this, jsii.String("NLB2"), &NetworkLoadBalancerAttributes{
//   	LoadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:ap-south-1:111111111111:loadbalancer/app/my-load-balancer2/5513dc2ea8a1"),
//   })
//
//   // Add one EndpointGroup for each Region we are targeting
//   listener.AddEndpointGroup(jsii.String("Group1"), &EndpointGroupOptions{
//   	Endpoints: []IEndpoint{
//   		ga_endpoints.NewNetworkLoadBalancerEndpoint(nlb1),
//   	},
//   })
//   listener.AddEndpointGroup(jsii.String("Group2"), &EndpointGroupOptions{
//   	// Imported load balancers automatically calculate their Region from the ARN.
//   	// If you are load balancing to other resources, you must also pass a `region`
//   	// parameter here.
//   	Endpoints: []IEndpoint{
//   		ga_endpoints.NewNetworkLoadBalancerEndpoint(nlb2),
//   	},
//   })
//
type ListenerOptions struct {
	// The list of port ranges for the connections from clients to the accelerator.
	PortRanges *[]*PortRange `field:"required" json:"portRanges" yaml:"portRanges"`
	// Client affinity to direct all requests from a user to the same endpoint.
	//
	// If you have stateful applications, client affinity lets you direct all
	// requests from a user to the same endpoint.
	//
	// By default, each connection from each client is routed to separate
	// endpoints. Set client affinity to SOURCE_IP to route all connections from
	// a single client to the same endpoint.
	// Default: ClientAffinity.NONE
	//
	ClientAffinity ClientAffinity `field:"optional" json:"clientAffinity" yaml:"clientAffinity"`
	// Name of the listener.
	// Default: - logical ID of the resource.
	//
	ListenerName *string `field:"optional" json:"listenerName" yaml:"listenerName"`
	// The protocol for the connections from clients to the accelerator.
	// Default: ConnectionProtocol.TCP
	//
	Protocol ConnectionProtocol `field:"optional" json:"protocol" yaml:"protocol"`
}

