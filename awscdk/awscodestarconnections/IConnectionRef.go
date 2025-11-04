package awscodestarconnections

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Connection.
// Experimental.
type IConnectionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Connection resource.
	// Experimental.
	ConnectionRef() *ConnectionReference
}

// The jsii proxy for IConnectionRef
type jsiiProxy_IConnectionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IConnectionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

