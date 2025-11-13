package interfacesawsathena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsathena/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PreparedStatement.
// Experimental.
type IPreparedStatementRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PreparedStatement resource.
	// Experimental.
	PreparedStatementRef() *PreparedStatementReference
}

// The jsii proxy for IPreparedStatementRef
type jsiiProxy_IPreparedStatementRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IPreparedStatementRef) PreparedStatementRef() *PreparedStatementReference {
	var returns *PreparedStatementReference
	_jsii_.Get(
		j,
		"preparedStatementRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPreparedStatementRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPreparedStatementRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

