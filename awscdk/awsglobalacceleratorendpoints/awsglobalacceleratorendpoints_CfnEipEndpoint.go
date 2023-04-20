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
// Experimental.
type CfnEipEndpoint interface {
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

	if err := validateNewCfnEipEndpointParameters(eip, options); err != nil {
		panic(err)
	}
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

