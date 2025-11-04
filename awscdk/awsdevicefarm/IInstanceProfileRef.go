package awsdevicefarm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceProfile.
// Experimental.
type IInstanceProfileRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a InstanceProfile resource.
	// Experimental.
	InstanceProfileRef() *InstanceProfileReference
}

// The jsii proxy for IInstanceProfileRef
type jsiiProxy_IInstanceProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IInstanceProfileRef) InstanceProfileRef() *InstanceProfileReference {
	var returns *InstanceProfileReference
	_jsii_.Get(
		j,
		"instanceProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceProfileRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

