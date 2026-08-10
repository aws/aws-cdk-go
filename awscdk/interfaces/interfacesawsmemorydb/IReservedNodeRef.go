package interfacesawsmemorydb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmemorydb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReservedNode.
// Experimental.
type IReservedNodeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ReservedNode resource.
	// Experimental.
	ReservedNodeRef() *ReservedNodeReference
}

// The jsii proxy for IReservedNodeRef
type jsiiProxy_IReservedNodeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IReservedNodeRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IReservedNodeRef) ReservedNodeRef() *ReservedNodeReference {
	var returns *ReservedNodeReference
	_jsii_.Get(
		j,
		"reservedNodeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReservedNodeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReservedNodeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

