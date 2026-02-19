package interfacesawss3vectors

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3vectors/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VectorBucket.
// Experimental.
type IVectorBucketRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VectorBucket resource.
	// Experimental.
	VectorBucketRef() *VectorBucketReference
}

// The jsii proxy for IVectorBucketRef
type jsiiProxy_IVectorBucketRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IVectorBucketRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVectorBucketRef) VectorBucketRef() *VectorBucketReference {
	var returns *VectorBucketReference
	_jsii_.Get(
		j,
		"vectorBucketRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorBucketRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorBucketRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

