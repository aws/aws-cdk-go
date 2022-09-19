package awsglobalacceleratorendpoints


// Properties for a NetworkLoadBalancerEndpoint.
//
// Example:
//   var nlb networkLoadBalancer
//   var listener listener
//
//
//   listener.addEndpointGroup(jsii.String("Group"), &endpointGroupOptions{
//   	endpoints: []iEndpoint{
//   		ga_endpoints.NewNetworkLoadBalancerEndpoint(nlb, &networkLoadBalancerEndpointProps{
//   			weight: jsii.Number(128),
//   		}),
//   	},
//   })
//
type NetworkLoadBalancerEndpointProps struct {
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

