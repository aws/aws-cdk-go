package awsssmincidents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmincidents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationSet.
// Experimental.
type IReplicationSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IReplicationSetRef
type jsiiProxy_IReplicationSetRef struct {
	internal.Type__constructsIConstruct
}

