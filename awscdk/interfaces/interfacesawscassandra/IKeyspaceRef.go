package interfacesawscassandra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscassandra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Keyspace.
// Experimental.
type IKeyspaceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Keyspace resource.
	// Experimental.
	KeyspaceRef() *KeyspaceReference
}

// The jsii proxy for IKeyspaceRef
type jsiiProxy_IKeyspaceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IKeyspaceRef) KeyspaceRef() *KeyspaceReference {
	var returns *KeyspaceReference
	_jsii_.Get(
		j,
		"keyspaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyspaceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyspaceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

