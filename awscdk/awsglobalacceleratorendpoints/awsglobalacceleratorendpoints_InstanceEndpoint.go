package awsglobalacceleratorendpoints

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsglobalaccelerator"
	"github.com/aws/aws-cdk-go/awscdk/awsglobalacceleratorendpoints/internal"
)

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
// Experimental.
type InstanceEndpoint interface {
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

