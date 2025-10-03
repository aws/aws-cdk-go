package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdocdb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBCluster.
// Experimental.
type IDBClusterRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDBClusterRef
type jsiiProxy_IDBClusterRef struct {
	internal.Type__constructsIConstruct
}

