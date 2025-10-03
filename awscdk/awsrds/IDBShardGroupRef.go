package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBShardGroup.
// Experimental.
type IDBShardGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDBShardGroupRef
type jsiiProxy_IDBShardGroupRef struct {
	internal.Type__constructsIConstruct
}

