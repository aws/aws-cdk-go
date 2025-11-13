package interfacesawsdynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdynamodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GlobalTable.
// Experimental.
type IGlobalTableRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a GlobalTable resource.
	// Experimental.
	GlobalTableRef() *GlobalTableReference
}

// The jsii proxy for IGlobalTableRef
type jsiiProxy_IGlobalTableRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IGlobalTableRef) GlobalTableRef() *GlobalTableReference {
	var returns *GlobalTableReference
	_jsii_.Get(
		j,
		"globalTableRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGlobalTableRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGlobalTableRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

