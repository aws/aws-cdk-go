package awsglobalacceleratorendpoints


// Properties for a NetworkLoadBalancerEndpoint.
//
// Example:
//   var nlb networkLoadBalancer
//   var listener listener
//
//
//   listener.AddEndpointGroup(jsii.String("Group"), &EndpointGroupOptions{
//   	Endpoints: []iEndpoint{
//   		ga_endpoints.NewNetworkLoadBalancerEndpoint(nlb, &NetworkLoadBalancerEndpointProps{
//   			Weight: jsii.Number(128),
//   		}),
//   	},
//   })
//
// Experimental.
type NetworkLoadBalancerEndpointProps struct {
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

