package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HookVersion.
// Experimental.
type IHookVersionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a HookVersion resource.
	// Experimental.
	HookVersionRef() *HookVersionReference
}

// The jsii proxy for IHookVersionRef
type jsiiProxy_IHookVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IHookVersionRef) HookVersionRef() *HookVersionReference {
	var returns *HookVersionReference
	_jsii_.Get(
		j,
		"hookVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHookVersionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHookVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

