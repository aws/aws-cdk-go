package awsglobalacceleratorendpoints


// Properties for a ApplicationLoadBalancerEndpoint.
//
// Example:
//   var alb applicationLoadBalancer
//   var listener listener
//
//
//   listener.AddEndpointGroup(jsii.String("Group"), &EndpointGroupOptions{
//   	Endpoints: []iEndpoint{
//   		ga_endpoints.NewApplicationLoadBalancerEndpoint(alb, &ApplicationLoadBalancerEndpointOptions{
//   			Weight: jsii.Number(128),
//   			PreserveClientIp: jsii.Boolean(true),
//   		}),
//   	},
//   })
//
// Experimental.
type ApplicationLoadBalancerEndpointOptions struct {
	// Forward the client IP address in an `X-Forwarded-For` header.
	//
	// GlobalAccelerator will create Network Interfaces in your VPC in order
	// to preserve the client IP address.
	//
	// Client IP address preservation is supported only in specific AWS Regions.
	// See the GlobalAccelerator Developer Guide for a list.
	// Experimental.
	PreserveClientIp *bool `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

