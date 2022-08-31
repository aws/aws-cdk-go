package awsdocdb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
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
//   endpoint := awscdk.Aws_docdb.NewEndpoint(jsii.String("address"), jsii.Number(123))
//
// Experimental.
type Endpoint interface {
	// The hostname of the endpoint.
	// Experimental.
	Hostname() *string
	// The port number of the endpoint.
	//
	// This can potentially be a CDK token. If you need to embed the port in a string (e.g. instance user data script),
	// use {@link Endpoint.portAsString}.
	// Experimental.
	Port() *float64
	// The combination of ``HOSTNAME:PORT`` for this endpoint.
	// Experimental.
	SocketAddress() *string
	// Returns the port number as a string representation that can be used for embedding within other strings.
	//
	// This is intended to deal with CDK's token system. Numeric CDK tokens are not expanded when their string
	// representation is embedded in a string. This function returns the port either as an unresolved string token or
	// as a resolved string representation of the port value.
	//
	// Returns: An (un)resolved string representation of the endpoint's port number.
	// Experimental.
	PortAsString() *string
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


// Constructs an Endpoint instance.
// Experimental.
func NewEndpoint(address *string, port *float64) Endpoint {
	_init_.Initialize()

	if err := validateNewEndpointParameters(address, port); err != nil {
		panic(err)
	}
	j := jsiiProxy_Endpoint{}

	_jsii_.Create(
		"monocdk.aws_docdb.Endpoint",
		[]interface{}{address, port},
		&j,
	)

	return &j
}

// Constructs an Endpoint instance.
// Experimental.
func NewEndpoint_Override(e Endpoint, address *string, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_docdb.Endpoint",
		[]interface{}{address, port},
		e,
	)
}

func (e *jsiiProxy_Endpoint) PortAsString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"portAsString",
		nil, // no parameters
		&returns,
	)

	return returns
}

