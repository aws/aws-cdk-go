package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CacheCluster.
// Experimental.
type ICacheClusterRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICacheClusterRef
type jsiiProxy_ICacheClusterRef struct {
	internal.Type__constructsIConstruct
}

