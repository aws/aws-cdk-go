package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationConfig.
// Experimental.
type IReplicationConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for IReplicationConfigRef
type jsiiProxy_IReplicationConfigRef struct {
	internal.Type__constructsIConstruct
}

