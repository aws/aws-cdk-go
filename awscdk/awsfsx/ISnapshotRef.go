package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfsx/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Snapshot.
// Experimental.
type ISnapshotRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Snapshot resource.
	// Experimental.
	SnapshotRef() *SnapshotReference
}

// The jsii proxy for ISnapshotRef
type jsiiProxy_ISnapshotRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ISnapshotRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

