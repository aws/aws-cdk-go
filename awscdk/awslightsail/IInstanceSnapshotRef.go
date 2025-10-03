package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceSnapshot.
// Experimental.
type IInstanceSnapshotRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInstanceSnapshotRef
type jsiiProxy_IInstanceSnapshotRef struct {
	internal.Type__constructsIConstruct
}

