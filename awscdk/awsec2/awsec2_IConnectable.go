package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An object that has a Connections object.
type IConnectable interface {
	// The network connections associated with this resource.
	Connections() Connections
}

// The jsii proxy for IConnectable
type jsiiProxy_IConnectable struct {
	_ byte // padding
}

func (j *jsiiProxy_IConnectable) Connections() Connections {
	var returns Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

