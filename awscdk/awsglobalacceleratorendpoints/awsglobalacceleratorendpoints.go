package awsglobalacceleratorendpoints

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awsglobalaccelerator"
	"github.com/aws/aws-cdk-go/awscdk/awsglobalacceleratorendpoints/internal"
)

// Use an Application Load Balancer as a Global Accelerator Endpoint.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationLoadBalancerEndpoint interface {
	awsglobalaccelerator.IEndpoint
	Region() *string
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


// Experimental.
func NewApplicationLoadBalancerEndpoint(loadBalancer awselasticloadbalancingv2.IApplicationLoadBalancer, options *ApplicationLoadBalancerEndpointOptions) ApplicationLoadBalancerEndpoint {
	_init_.Initialize()

	j := jsiiProxy_ApplicationLoadBalancerEndpoint{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator_endpoints.ApplicationLoadBalancerEndpoint",
		[]interface{}{loadBalancer, options},
		&j,
	)

	return &j
}

// Experimental.
func NewApplicationLoadBalancerEndpoint_Override(a ApplicationLoadBalancerEndpoint, loadBalancer awselasticloadbalancingv2.IApplicationLoadBalancer, options *ApplicationLoadBalancerEndpointOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator_endpoints.ApplicationLoadBalancerEndpoint",
		[]interface{}{loadBalancer, options},
		a,
	)
}

// Render the endpoint to an endpoint configuration.
// Experimental.
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
// TODO: EXAMPLE
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
	PreserveClientIp *bool `json:"preserveClientIp"`
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	// Experimental.
	Weight *float64 `json:"weight"`
}

// Use an EC2 Instance as a Global Accelerator Endpoint.
//
// TODO: EXAMPLE
//
// Experimental.
type CfnEipEndpoint interface {
	awsglobalaccelerator.IEndpoint
	Region() *string
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


// Experimental.
func NewCfnEipEndpoint(eip awsec2.CfnEIP, options *CfnEipEndpointProps) CfnEipEndpoint {
	_init_.Initialize()

	j := jsiiProxy_CfnEipEndpoint{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator_endpoints.CfnEipEndpoint",
		[]interface{}{eip, options},
		&j,
	)

	return &j
}

// Experimental.
func NewCfnEipEndpoint_Override(c CfnEipEndpoint, eip awsec2.CfnEIP, options *CfnEipEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator_endpoints.CfnEipEndpoint",
		[]interface{}{eip, options},
		c,
	)
}

// Render the endpoint to an endpoint configuration.
// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type CfnEipEndpointProps struct {
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	// Experimental.
	Weight *float64 `json:"weight"`
}

// Use an EC2 Instance as a Global Accelerator Endpoint.
//
// TODO: EXAMPLE
//
// Experimental.
type InstanceEndpoint interface {
	awsglobalaccelerator.IEndpoint
	Region() *string
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


// Experimental.
func NewInstanceEndpoint(instance awsec2.IInstance, options *InstanceEndpointProps) InstanceEndpoint {
	_init_.Initialize()

	j := jsiiProxy_InstanceEndpoint{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator_endpoints.InstanceEndpoint",
		[]interface{}{instance, options},
		&j,
	)

	return &j
}

// Experimental.
func NewInstanceEndpoint_Override(i InstanceEndpoint, instance awsec2.IInstance, options *InstanceEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator_endpoints.InstanceEndpoint",
		[]interface{}{instance, options},
		i,
	)
}

// Render the endpoint to an endpoint configuration.
// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type InstanceEndpointProps struct {
	// Forward the client IP address.
	//
	// GlobalAccelerator will create Network Interfaces in your VPC in order
	// to preserve the client IP address.
	//
	// Client IP address preservation is supported only in specific AWS Regions.
	// See the GlobalAccelerator Developer Guide for a list.
	// Experimental.
	PreserveClientIp *bool `json:"preserveClientIp"`
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	// Experimental.
	Weight *float64 `json:"weight"`
}

// Use a Network Load Balancer as a Global Accelerator Endpoint.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkLoadBalancerEndpoint interface {
	awsglobalaccelerator.IEndpoint
	Region() *string
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

// Render the endpoint to an endpoint configuration.
// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type NetworkLoadBalancerEndpointProps struct {
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	// Experimental.
	Weight *float64 `json:"weight"`
}

