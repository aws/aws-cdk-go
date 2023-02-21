package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Connection endpoint of a database cluster or instance.
//
// Consists of a combination of hostname and port.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpoint := awscdk.Aws_rds.NewEndpoint(jsii.String("address"), jsii.Number(123))
//
type Endpoint interface {
	// The hostname of the endpoint.
	Hostname() *string
	// The port of the endpoint.
	Port() *float64
	// The combination of "HOSTNAME:PORT" for this endpoint.
	SocketAddress() *string
}

// The jsii proxy struct for Endpoint
type jsiiProxy_Endpoint struct {
	_ byte // padding
}

func (j *jsiiProxy_Endpoint) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Endpoint) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Endpoint) SocketAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketAddress",
		&returns,
	)
	return returns
}


func NewEndpoint(address *string, port *float64) Endpoint {
	_init_.Initialize()

	if err := validateNewEndpointParameters(address, port); err != nil {
		panic(err)
	}
	j := jsiiProxy_Endpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.Endpoint",
		[]interface{}{address, port},
		&j,
	)

	return &j
}

func NewEndpoint_Override(e Endpoint, address *string, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.Endpoint",
		[]interface{}{address, port},
		e,
	)
}

