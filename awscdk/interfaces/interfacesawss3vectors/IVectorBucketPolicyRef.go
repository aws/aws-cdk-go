package interfacesawss3vectors

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3vectors/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VectorBucketPolicy.
// Experimental.
type IVectorBucketPolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VectorBucketPolicy resource.
	// Experimental.
	VectorBucketPolicyRef() *VectorBucketPolicyReference
}

// The jsii proxy for IVectorBucketPolicyRef
type jsiiProxy_IVectorBucketPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IVectorBucketPolicyRef) VectorBucketPolicyRef() *VectorBucketPolicyReference {
	var returns *VectorBucketPolicyReference
	_jsii_.Get(
		j,
		"vectorBucketPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorBucketPolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorBucketPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

