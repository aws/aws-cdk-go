package interfacesawsredshiftserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsredshiftserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Snapshot.
// Experimental.
type ISnapshotRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Snapshot resource.
	// Experimental.
	SnapshotRef() *SnapshotReference
}

// The jsii proxy for ISnapshotRef
type jsiiProxy_ISnapshotRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISnapshotRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISnapshotRef) SnapshotRef() *SnapshotReference {
	var returns *SnapshotReference
	_jsii_.Get(
		j,
		"snapshotRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISnapshotRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISnapshotRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

