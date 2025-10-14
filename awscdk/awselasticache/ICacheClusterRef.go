package awselasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CacheCluster.
// Experimental.
type ICacheClusterRef interface {
	constructs.IConstruct
	// A reference to a CacheCluster resource.
	// Experimental.
	CacheClusterRef() *CacheClusterReference
}

// The jsii proxy for ICacheClusterRef
type jsiiProxy_ICacheClusterRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICacheClusterRef) CacheClusterRef() *CacheClusterReference {
	var returns *CacheClusterReference
	_jsii_.Get(
		j,
		"cacheClusterRef",
		&returns,
	)
	return returns
}

