package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationInstance.
// Experimental.
type IReplicationInstanceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IReplicationInstanceRef
type jsiiProxy_IReplicationInstanceRef struct {
	internal.Type__constructsIConstruct
}

