package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBClusterParameterGroup.
// Experimental.
type IDBClusterParameterGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDBClusterParameterGroupRef
type jsiiProxy_IDBClusterParameterGroupRef struct {
	internal.Type__constructsIConstruct
}

