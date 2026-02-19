package interfacesawss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BucketPolicy.
// Experimental.
type IBucketPolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a BucketPolicy resource.
	// Experimental.
	BucketPolicyRef() *BucketPolicyReference
}

// The jsii proxy for IBucketPolicyRef
type jsiiProxy_IBucketPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IBucketPolicyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IBucketPolicyRef) BucketPolicyRef() *BucketPolicyReference {
	var returns *BucketPolicyReference
	_jsii_.Get(
		j,
		"bucketPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBucketPolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBucketPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

