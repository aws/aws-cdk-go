package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CachePolicy.
// Experimental.
type ICachePolicyRef interface {
	constructs.IConstruct
	// A reference to a CachePolicy resource.
	// Experimental.
	CachePolicyRef() *CachePolicyReference
}

// The jsii proxy for ICachePolicyRef
type jsiiProxy_ICachePolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICachePolicyRef) CachePolicyRef() *CachePolicyReference {
	var returns *CachePolicyReference
	_jsii_.Get(
		j,
		"cachePolicyRef",
		&returns,
	)
	return returns
}

