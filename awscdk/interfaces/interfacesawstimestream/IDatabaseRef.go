package interfacesawstimestream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawstimestream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Database.
// Experimental.
type IDatabaseRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Database resource.
	// Experimental.
	DatabaseRef() *DatabaseReference
}

// The jsii proxy for IDatabaseRef
type jsiiProxy_IDatabaseRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDatabaseRef) DatabaseRef() *DatabaseReference {
	var returns *DatabaseReference
	_jsii_.Get(
		j,
		"databaseRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

