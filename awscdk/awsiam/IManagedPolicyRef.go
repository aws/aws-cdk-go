package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ManagedPolicy.
// Experimental.
type IManagedPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ManagedPolicy resource.
	// Experimental.
	ManagedPolicyRef() *ManagedPolicyReference
}

// The jsii proxy for IManagedPolicyRef
type jsiiProxy_IManagedPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IManagedPolicyRef) ManagedPolicyRef() *ManagedPolicyReference {
	var returns *ManagedPolicyReference
	_jsii_.Get(
		j,
		"managedPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

