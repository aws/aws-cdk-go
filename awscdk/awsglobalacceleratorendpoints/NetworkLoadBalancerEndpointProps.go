package awsglobalacceleratorendpoints


// Properties for a NetworkLoadBalancerEndpoint.
//
// Example:
//   var nlb NetworkLoadBalancer
//   var listener Listener
//
//
//   listener.AddEndpointGroup(jsii.String("Group"), &EndpointGroupOptions{
//   	Endpoints: []IEndpoint{
//   		ga_endpoints.NewNetworkLoadBalancerEndpoint(nlb, &NetworkLoadBalancerEndpointProps{
//   			Weight: jsii.Number(128),
//   			PreserveClientIp: jsii.Boolean(true),
//   		}),
//   	},
//   })
//
type NetworkLoadBalancerEndpointProps struct {
	// Forward the client IP address in an `X-Forwarded-For` header.
	//
	// GlobalAccelerator will create Network Interfaces in your VPC in order
	// to preserve the client IP address.
	//
	// Client IP address preservation is supported only in specific AWS Regions.
	// See the GlobalAccelerator Developer Guide for a list.
	// Default: false.
	//
	PreserveClientIp *bool `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	// Default: 128.
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

