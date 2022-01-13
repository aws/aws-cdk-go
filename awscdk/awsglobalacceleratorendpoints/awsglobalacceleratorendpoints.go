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
// TODO: EXAMPLE
//
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

// Render the endpoint to an endpoint configuration.
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
type ApplicationLoadBalancerEndpointOptions struct {
	// Forward the client IP address in an `X-Forwarded-For` header.
	//
	// GlobalAccelerator will create Network Interfaces in your VPC in order
	// to preserve the client IP address.
	//
	// Client IP address preservation is supported only in specific AWS Regions.
	// See the GlobalAccelerator Developer Guide for a list.
	PreserveClientIp *bool `json:"preserveClientIp" yaml:"preserveClientIp"`
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	Weight *float64 `json:"weight" yaml:"weight"`
}

// Use an EC2 Instance as a Global Accelerator Endpoint.
//
// TODO: EXAMPLE
//
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

// Render the endpoint to an endpoint configuration.
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
type CfnEipEndpointProps struct {
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	Weight *float64 `json:"weight" yaml:"weight"`
}

// Use an EC2 Instance as a Global Accelerator Endpoint.
//
// TODO: EXAMPLE
//
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

// Render the endpoint to an endpoint configuration.
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
type InstanceEndpointProps struct {
	// Forward the client IP address.
	//
	// GlobalAccelerator will create Network Interfaces in your VPC in order
	// to preserve the client IP address.
	//
	// Client IP address preservation is supported only in specific AWS Regions.
	// See the GlobalAccelerator Developer Guide for a list.
	PreserveClientIp *bool `json:"preserveClientIp" yaml:"preserveClientIp"`
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	Weight *float64 `json:"weight" yaml:"weight"`
}

// Use a Network Load Balancer as a Global Accelerator Endpoint.
//
// TODO: EXAMPLE
//
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

// Render the endpoint to an endpoint configuration.
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
type NetworkLoadBalancerEndpointProps struct {
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	Weight *float64 `json:"weight" yaml:"weight"`
}

