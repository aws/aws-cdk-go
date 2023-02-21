package awsglobalacceleratorendpoints


// Properties for a NetworkLoadBalancerEndpoint.
//
// Example:
//   var listener listener
//   var eip cfnEIP
//
//
//   listener.AddEndpointGroup(jsii.String("Group"), &EndpointGroupOptions{
//   	Endpoints: []iEndpoint{
//   		ga_endpoints.NewCfnEipEndpoint(eip, &CfnEipEndpointProps{
//   			Weight: jsii.Number(128),
//   		}),
//   	},
//   })
//
type CfnEipEndpointProps struct {
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

