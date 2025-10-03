package awsmsk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Replicator.
// Experimental.
type IReplicatorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IReplicatorRef
type jsiiProxy_IReplicatorRef struct {
	internal.Type__constructsIConstruct
}

