package interfacesawslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceSnapshot.
// Experimental.
type IInstanceSnapshotRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a InstanceSnapshot resource.
	// Experimental.
	InstanceSnapshotRef() *InstanceSnapshotReference
}

// The jsii proxy for IInstanceSnapshotRef
type jsiiProxy_IInstanceSnapshotRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IInstanceSnapshotRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IInstanceSnapshotRef) InstanceSnapshotRef() *InstanceSnapshotReference {
	var returns *InstanceSnapshotReference
	_jsii_.Get(
		j,
		"instanceSnapshotRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceSnapshotRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceSnapshotRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

