package interfacesawsneptunegraph

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsneptunegraph/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GraphSnapshot.
// Experimental.
type IGraphSnapshotRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a GraphSnapshot resource.
	// Experimental.
	GraphSnapshotRef() *GraphSnapshotReference
}

// The jsii proxy for IGraphSnapshotRef
type jsiiProxy_IGraphSnapshotRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IGraphSnapshotRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IGraphSnapshotRef) GraphSnapshotRef() *GraphSnapshotReference {
	var returns *GraphSnapshotReference
	_jsii_.Get(
		j,
		"graphSnapshotRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphSnapshotRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphSnapshotRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

