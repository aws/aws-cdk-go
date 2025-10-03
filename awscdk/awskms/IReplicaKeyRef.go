package awskms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicaKey.
// Experimental.
type IReplicaKeyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IReplicaKeyRef
type jsiiProxy_IReplicaKeyRef struct {
	internal.Type__constructsIConstruct
}

