package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// An Http destination for data from a Kinesis Data Firehose delivery stream.
//
// Example:
//   var endpointConfig HttpEndpointConfig
//
//   httpDestination := firehose.NewHttpEndpoint(&HttpEndpointProps{
//   	EndpointConfig: EndpointConfig,
//   })
//
type HttpEndpoint interface {
	IDestination
	// Binds this destination to the Amazon Data Firehose delivery stream.
	//
	// Implementers should use this method to bind resources to the stack and initialize values using the provided stream.
	Bind(scope constructs.Construct, options *DestinationBindOptions) *DestinationConfig
}

// The jsii proxy struct for HttpEndpoint
type jsiiProxy_HttpEndpoint struct {
	jsiiProxy_IDestination
}

func NewHttpEndpoint(props *HttpEndpointProps) HttpEndpoint {
	_init_.Initialize()

	if err := validateNewHttpEndpointParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.HttpEndpoint",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewHttpEndpoint_Override(h HttpEndpoint, props *HttpEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.HttpEndpoint",
		[]interface{}{props},
		h,
	)
}

func (h *jsiiProxy_HttpEndpoint) Bind(scope constructs.Construct, options *DestinationBindOptions) *DestinationConfig {
	if err := h.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *DestinationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

