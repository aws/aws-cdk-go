package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfsx/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Snapshot.
// Experimental.
type ISnapshotRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISnapshotRef
type jsiiProxy_ISnapshotRef struct {
	internal.Type__constructsIConstruct
}

