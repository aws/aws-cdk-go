package awsglobalaccelerator

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
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
//   rawEndpoint := awscdk.Aws_globalaccelerator.NewRawEndpoint(&RawEndpointProps{
//   	EndpointId: jsii.String("endpointId"),
//
//   	// the properties below are optional
//   	PreserveClientIp: jsii.Boolean(false),
//   	Region: jsii.String("region"),
//   	Weight: jsii.Number(123),
//   })
//
type RawEndpoint interface {
	IEndpoint
	// The region where the endpoint is located.
	//
	// If the region cannot be determined, `undefined` is returned.
	Region() *string
	// Render the endpoint to an endpoint configuration.
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


func NewRawEndpoint(props *RawEndpointProps) RawEndpoint {
	_init_.Initialize()

	if err := validateNewRawEndpointParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RawEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.RawEndpoint",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewRawEndpoint_Override(r RawEndpoint, props *RawEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.RawEndpoint",
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

