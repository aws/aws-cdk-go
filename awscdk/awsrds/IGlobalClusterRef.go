package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GlobalCluster.
// Experimental.
type IGlobalClusterRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGlobalClusterRef
type jsiiProxy_IGlobalClusterRef struct {
	internal.Type__constructsIConstruct
}

