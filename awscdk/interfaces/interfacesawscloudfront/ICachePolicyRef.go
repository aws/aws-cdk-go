package interfacesawscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CachePolicy.
// Experimental.
type ICachePolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CachePolicy resource.
	// Experimental.
	CachePolicyRef() *CachePolicyReference
}

// The jsii proxy for ICachePolicyRef
type jsiiProxy_ICachePolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_ICachePolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICachePolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

