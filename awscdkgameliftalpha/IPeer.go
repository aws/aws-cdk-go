package awscdkgameliftalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for classes that provide the peer-specification parts of an inbound permission.
// Experimental.
type IPeer interface {
	// Produce the ingress rule JSON for the given connection.
	// Experimental.
	ToJson() interface{}
	// A unique identifier for this connection peer.
	// Experimental.
	UniqueId() *string
}

// The jsii proxy for IPeer
type jsiiProxy_IPeer struct {
	_ byte // padding
}

func (i *jsiiProxy_IPeer) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IPeer) UniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueId",
		&returns,
	)
	return returns
}

