package awsglobalacceleratorendpoints

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalacceleratorendpoints/internal"
)

// Use a Network Load Balancer as a Global Accelerator Endpoint.
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
type NetworkLoadBalancerEndpoint interface {
	awsglobalaccelerator.IEndpoint
	// The region where the endpoint is located.
	//
	// If the region cannot be determined, `undefined` is returned.
	Region() *string
	// Render the endpoint to an endpoint configuration.
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


func NewNetworkLoadBalancerEndpoint(loadBalancer awselasticloadbalancingv2.INetworkLoadBalancer, options *NetworkLoadBalancerEndpointProps) NetworkLoadBalancerEndpoint {
	_init_.Initialize()

	if err := validateNewNetworkLoadBalancerEndpointParameters(loadBalancer, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkLoadBalancerEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.NetworkLoadBalancerEndpoint",
		[]interface{}{loadBalancer, options},
		&j,
	)

	return &j
}

func NewNetworkLoadBalancerEndpoint_Override(n NetworkLoadBalancerEndpoint, loadBalancer awselasticloadbalancingv2.INetworkLoadBalancer, options *NetworkLoadBalancerEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.NetworkLoadBalancerEndpoint",
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

