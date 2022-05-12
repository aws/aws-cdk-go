package awsglobalacceleratorendpoints

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalacceleratorendpoints/internal"
)

// Use an Application Load Balancer as a Global Accelerator Endpoint.
//
// Example:
//   var alb applicationLoadBalancer
//   var listener listener
//
//
//   listener.addEndpointGroup(jsii.String("Group"), &endpointGroupOptions{
//   	endpoints: []iEndpoint{
//   		ga_endpoints.NewApplicationLoadBalancerEndpoint(alb, &applicationLoadBalancerEndpointOptions{
//   			weight: jsii.Number(128),
//   			preserveClientIp: jsii.Boolean(true),
//   		}),
//   	},
//   })
//
type ApplicationLoadBalancerEndpoint interface {
	awsglobalaccelerator.IEndpoint
	// The region where the endpoint is located.
	//
	// If the region cannot be determined, `undefined` is returned.
	Region() *string
	// Render the endpoint to an endpoint configuration.
	RenderEndpointConfiguration() interface{}
}

// The jsii proxy struct for ApplicationLoadBalancerEndpoint
type jsiiProxy_ApplicationLoadBalancerEndpoint struct {
	internal.Type__awsglobalacceleratorIEndpoint
}

func (j *jsiiProxy_ApplicationLoadBalancerEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}


func NewApplicationLoadBalancerEndpoint(loadBalancer awselasticloadbalancingv2.IApplicationLoadBalancer, options *ApplicationLoadBalancerEndpointOptions) ApplicationLoadBalancerEndpoint {
	_init_.Initialize()

	j := jsiiProxy_ApplicationLoadBalancerEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.ApplicationLoadBalancerEndpoint",
		[]interface{}{loadBalancer, options},
		&j,
	)

	return &j
}

func NewApplicationLoadBalancerEndpoint_Override(a ApplicationLoadBalancerEndpoint, loadBalancer awselasticloadbalancingv2.IApplicationLoadBalancer, options *ApplicationLoadBalancerEndpointOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.ApplicationLoadBalancerEndpoint",
		[]interface{}{loadBalancer, options},
		a,
	)
}

func (a *jsiiProxy_ApplicationLoadBalancerEndpoint) RenderEndpointConfiguration() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderEndpointConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a ApplicationLoadBalancerEndpoint.
//
// Example:
//   var alb applicationLoadBalancer
//   var listener listener
//
//
//   listener.addEndpointGroup(jsii.String("Group"), &endpointGroupOptions{
//   	endpoints: []iEndpoint{
//   		ga_endpoints.NewApplicationLoadBalancerEndpoint(alb, &applicationLoadBalancerEndpointOptions{
//   			weight: jsii.Number(128),
//   			preserveClientIp: jsii.Boolean(true),
//   		}),
//   	},
//   })
//
type ApplicationLoadBalancerEndpointOptions struct {
	// Forward the client IP address in an `X-Forwarded-For` header.
	//
	// GlobalAccelerator will create Network Interfaces in your VPC in order
	// to preserve the client IP address.
	//
	// Client IP address preservation is supported only in specific AWS Regions.
	// See the GlobalAccelerator Developer Guide for a list.
	PreserveClientIp *bool `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

// Use an EC2 Instance as a Global Accelerator Endpoint.
//
// Example:
//   var listener listener
//   var eip cfnEIP
//
//
//   listener.addEndpointGroup(jsii.String("Group"), &endpointGroupOptions{
//   	endpoints: []iEndpoint{
//   		ga_endpoints.NewCfnEipEndpoint(eip, &cfnEipEndpointProps{
//   			weight: jsii.Number(128),
//   		}),
//   	},
//   })
//
type CfnEipEndpoint interface {
	awsglobalaccelerator.IEndpoint
	// The region where the endpoint is located.
	//
	// If the region cannot be determined, `undefined` is returned.
	Region() *string
	// Render the endpoint to an endpoint configuration.
	RenderEndpointConfiguration() interface{}
}

// The jsii proxy struct for CfnEipEndpoint
type jsiiProxy_CfnEipEndpoint struct {
	internal.Type__awsglobalacceleratorIEndpoint
}

func (j *jsiiProxy_CfnEipEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}


func NewCfnEipEndpoint(eip awsec2.CfnEIP, options *CfnEipEndpointProps) CfnEipEndpoint {
	_init_.Initialize()

	j := jsiiProxy_CfnEipEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.CfnEipEndpoint",
		[]interface{}{eip, options},
		&j,
	)

	return &j
}

func NewCfnEipEndpoint_Override(c CfnEipEndpoint, eip awsec2.CfnEIP, options *CfnEipEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.CfnEipEndpoint",
		[]interface{}{eip, options},
		c,
	)
}

func (c *jsiiProxy_CfnEipEndpoint) RenderEndpointConfiguration() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderEndpointConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a NetworkLoadBalancerEndpoint.
//
// Example:
//   var listener listener
//   var eip cfnEIP
//
//
//   listener.addEndpointGroup(jsii.String("Group"), &endpointGroupOptions{
//   	endpoints: []iEndpoint{
//   		ga_endpoints.NewCfnEipEndpoint(eip, &cfnEipEndpointProps{
//   			weight: jsii.Number(128),
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

// Use an EC2 Instance as a Global Accelerator Endpoint.
//
// Example:
//   var listener listener
//   var instance instance
//
//
//   listener.addEndpointGroup(jsii.String("Group"), &endpointGroupOptions{
//   	endpoints: []iEndpoint{
//   		ga_endpoints.NewInstanceEndpoint(instance, &instanceEndpointProps{
//   			weight: jsii.Number(128),
//   			preserveClientIp: jsii.Boolean(true),
//   		}),
//   	},
//   })
//
type InstanceEndpoint interface {
	awsglobalaccelerator.IEndpoint
	// The region where the endpoint is located.
	//
	// If the region cannot be determined, `undefined` is returned.
	Region() *string
	// Render the endpoint to an endpoint configuration.
	RenderEndpointConfiguration() interface{}
}

// The jsii proxy struct for InstanceEndpoint
type jsiiProxy_InstanceEndpoint struct {
	internal.Type__awsglobalacceleratorIEndpoint
}

func (j *jsiiProxy_InstanceEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}


func NewInstanceEndpoint(instance awsec2.IInstance, options *InstanceEndpointProps) InstanceEndpoint {
	_init_.Initialize()

	j := jsiiProxy_InstanceEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.InstanceEndpoint",
		[]interface{}{instance, options},
		&j,
	)

	return &j
}

func NewInstanceEndpoint_Override(i InstanceEndpoint, instance awsec2.IInstance, options *InstanceEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.InstanceEndpoint",
		[]interface{}{instance, options},
		i,
	)
}

func (i *jsiiProxy_InstanceEndpoint) RenderEndpointConfiguration() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"renderEndpointConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a NetworkLoadBalancerEndpoint.
//
// Example:
//   var listener listener
//   var instance instance
//
//
//   listener.addEndpointGroup(jsii.String("Group"), &endpointGroupOptions{
//   	endpoints: []iEndpoint{
//   		ga_endpoints.NewInstanceEndpoint(instance, &instanceEndpointProps{
//   			weight: jsii.Number(128),
//   			preserveClientIp: jsii.Boolean(true),
//   		}),
//   	},
//   })
//
type InstanceEndpointProps struct {
	// Forward the client IP address.
	//
	// GlobalAccelerator will create Network Interfaces in your VPC in order
	// to preserve the client IP address.
	//
	// Client IP address preservation is supported only in specific AWS Regions.
	// See the GlobalAccelerator Developer Guide for a list.
	PreserveClientIp *bool `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

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

