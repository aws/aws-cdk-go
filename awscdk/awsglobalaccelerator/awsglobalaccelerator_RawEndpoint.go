package awsglobalaccelerator

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Untyped endpoint implementation.
//
// Prefer using the classes in the `aws-globalaccelerator-endpoints` package instead,
// as they accept typed constructs. You can use this class if you want to use an
// endpoint type that does not have an appropriate class in that package yet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rawEndpoint := awscdk.Aws_globalaccelerator.NewRawEndpoint(&rawEndpointProps{
//   	endpointId: jsii.String("endpointId"),
//
//   	// the properties below are optional
//   	preserveClientIp: jsii.Boolean(false),
//   	region: jsii.String("region"),
//   	weight: jsii.Number(123),
//   })
//
// Experimental.
type RawEndpoint interface {
	IEndpoint
	// The region where the endpoint is located.
	//
	// If the region cannot be determined, `undefined` is returned.
	// Experimental.
	Region() *string
	// Render the endpoint to an endpoint configuration.
	// Experimental.
	RenderEndpointConfiguration() interface{}
}

// The jsii proxy struct for RawEndpoint
type jsiiProxy_RawEndpoint struct {
	jsiiProxy_IEndpoint
}

func (j *jsiiProxy_RawEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}


// Experimental.
func NewRawEndpoint(props *RawEndpointProps) RawEndpoint {
	_init_.Initialize()

	if err := validateNewRawEndpointParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RawEndpoint{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.RawEndpoint",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewRawEndpoint_Override(r RawEndpoint, props *RawEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.RawEndpoint",
		[]interface{}{props},
		r,
	)
}

func (r *jsiiProxy_RawEndpoint) RenderEndpointConfiguration() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"renderEndpointConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

