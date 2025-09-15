package awscodeconnections

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeconnections/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Connection.
// Experimental.
type IConnectionRef interface {
	constructs.IConstruct
	// A reference to a Connection resource.
	// Experimental.
	ConnectionRef() *ConnectionReference
}

// The jsii proxy for IConnectionRef
type jsiiProxy_IConnectionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConnectionRef) ConnectionRef() *ConnectionReference {
	var returns *ConnectionReference
	_jsii_.Get(
		j,
		"connectionRef",
		&returns,
	)
	return returns
}

