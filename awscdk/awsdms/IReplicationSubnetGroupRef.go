package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationSubnetGroup.
// Experimental.
type IReplicationSubnetGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IReplicationSubnetGroupRef
type jsiiProxy_IReplicationSubnetGroupRef struct {
	internal.Type__constructsIConstruct
}

