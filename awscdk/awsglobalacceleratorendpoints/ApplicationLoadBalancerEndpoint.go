package awsglobalacceleratorendpoints

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

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
//   listener.AddEndpointGroup(jsii.String("Group"), &EndpointGroupOptions{
//   	Endpoints: []iEndpoint{
//   		ga_endpoints.NewApplicationLoadBalancerEndpoint(alb, &ApplicationLoadBalancerEndpointOptions{
//   			Weight: jsii.Number(128),
//   			PreserveClientIp: jsii.Boolean(true),
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

	if err := validateNewApplicationLoadBalancerEndpointParameters(loadBalancer, options); err != nil {
		panic(err)
	}
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

