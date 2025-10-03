package awsemrcontainers

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsemrcontainers/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualCluster.
// Experimental.
type IVirtualClusterRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVirtualClusterRef
type jsiiProxy_IVirtualClusterRef struct {
	internal.Type__constructsIConstruct
}

