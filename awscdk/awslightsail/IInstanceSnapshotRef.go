package awslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceSnapshot.
// Experimental.
type IInstanceSnapshotRef interface {
	constructs.IConstruct
	// A reference to a InstanceSnapshot resource.
	// Experimental.
	InstanceSnapshotRef() *InstanceSnapshotReference
}

// The jsii proxy for IInstanceSnapshotRef
type jsiiProxy_IInstanceSnapshotRef struct {
	internal.Type__constructsIConstruct
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

