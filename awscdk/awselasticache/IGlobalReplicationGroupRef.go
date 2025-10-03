package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GlobalReplicationGroup.
// Experimental.
type IGlobalReplicationGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGlobalReplicationGroupRef
type jsiiProxy_IGlobalReplicationGroupRef struct {
	internal.Type__constructsIConstruct
}

