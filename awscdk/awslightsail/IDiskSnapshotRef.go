package awslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DiskSnapshot.
// Experimental.
type IDiskSnapshotRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DiskSnapshot resource.
	// Experimental.
	DiskSnapshotRef() *DiskSnapshotReference
}

// The jsii proxy for IDiskSnapshotRef
type jsiiProxy_IDiskSnapshotRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDiskSnapshotRef) DiskSnapshotRef() *DiskSnapshotReference {
	var returns *DiskSnapshotReference
	_jsii_.Get(
		j,
		"diskSnapshotRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDiskSnapshotRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDiskSnapshotRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

