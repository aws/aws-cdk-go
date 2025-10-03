package awsmemorydb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmemorydb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MultiRegionCluster.
// Experimental.
type IMultiRegionClusterRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMultiRegionClusterRef
type jsiiProxy_IMultiRegionClusterRef struct {
	internal.Type__constructsIConstruct
}

