package awsglobalacceleratorendpoints

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awsglobalaccelerator"
	"github.com/aws/aws-cdk-go/awscdk/awsglobalacceleratorendpoints/internal"
)

// Use a Network Load Balancer as a Global Accelerator Endpoint.
//
// Example:
//   // Create an Accelerator
//   accelerator := globalaccelerator.NewAccelerator(this, jsii.String("Accelerator"))
//
//   // Create a Listener
//   listener := accelerator.addListener(jsii.String("Listener"), &listenerOptions{
//   	portRanges: []portRange{
//   		&portRange{
//   			fromPort: jsii.Number(80),
//   		},
//   		&portRange{
//   			fromPort: jsii.Number(443),
//   		},
//   	},
//   })
//
//   // Import the Load Balancers
//   nlb1 := elbv2.networkLoadBalancer.fromNetworkLoadBalancerAttributes(this, jsii.String("NLB1"), &networkLoadBalancerAttributes{
//   	loadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:us-west-2:111111111111:loadbalancer/app/my-load-balancer1/e16bef66805b"),
//   })
//   nlb2 := elbv2.networkLoadBalancer.fromNetworkLoadBalancerAttributes(this, jsii.String("NLB2"), &networkLoadBalancerAttributes{
//   	loadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:ap-south-1:111111111111:loadbalancer/app/my-load-balancer2/5513dc2ea8a1"),
//   })
//
//   // Add one EndpointGroup for each Region we are targeting
//   listener.addEndpointGroup(jsii.String("Group1"), &endpointGroupOptions{
//   	endpoints: []iEndpoint{
//   		ga_endpoints.NewNetworkLoadBalancerEndpoint(nlb1),
//   	},
//   })
//   listener.addEndpointGroup(jsii.String("Group2"), &endpointGroupOptions{
//   	// Imported load balancers automatically calculate their Region from the ARN.
//   	// If you are load balancing to other resources, you must also pass a `region`
//   	// parameter here.
//   	endpoints: []*iEndpoint{
//   		ga_endpoints.NewNetworkLoadBalancerEndpoint(nlb2),
//   	},
//   })
//
// Experimental.
type NetworkLoadBalancerEndpoint interface {
	awsglobalaccelerator.IEndpoint
	// The region where the endpoint is located.
	//
	// If the region cannot be determined, `undefined` is returned.
	// Experimental.
	Region() *string
	// Render the endpoint to an endpoint configuration.
	// Experimental.
	RenderEndpointConfiguration() interface{}
}

// The jsii proxy struct for NetworkLoadBalancerEndpoint
type jsiiProxy_NetworkLoadBalancerEndpoint struct {
	internal.Type__awsglobalacceleratorIEndpoint
}

func (j *jsiiProxy_NetworkLoadBalancerEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}


// Experimental.
func NewNetworkLoadBalancerEndpoint(loadBalancer awselasticloadbalancingv2.INetworkLoadBalancer, options *NetworkLoadBalancerEndpointProps) NetworkLoadBalancerEndpoint {
	_init_.Initialize()

	if err := validateNewNetworkLoadBalancerEndpointParameters(loadBalancer, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkLoadBalancerEndpoint{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator_endpoints.NetworkLoadBalancerEndpoint",
		[]interface{}{loadBalancer, options},
		&j,
	)

	return &j
}

// Experimental.
func NewNetworkLoadBalancerEndpoint_Override(n NetworkLoadBalancerEndpoint, loadBalancer awselasticloadbalancingv2.INetworkLoadBalancer, options *NetworkLoadBalancerEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator_endpoints.NetworkLoadBalancerEndpoint",
		[]interface{}{loadBalancer, options},
		n,
	)
}

func (n *jsiiProxy_NetworkLoadBalancerEndpoint) RenderEndpointConfiguration() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"renderEndpointConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

