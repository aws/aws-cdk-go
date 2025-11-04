package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RolePolicy.
// Experimental.
type IRolePolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RolePolicy resource.
	// Experimental.
	RolePolicyRef() *RolePolicyReference
}

// The jsii proxy for IRolePolicyRef
type jsiiProxy_IRolePolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IRolePolicyRef) RolePolicyRef() *RolePolicyReference {
	var returns *RolePolicyReference
	_jsii_.Get(
		j,
		"rolePolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRolePolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRolePolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

