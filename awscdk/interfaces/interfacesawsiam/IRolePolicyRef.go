package interfacesawsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RolePolicy.
// Experimental.
type IRolePolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RolePolicy resource.
	// Experimental.
	RolePolicyRef() *RolePolicyReference
}

// The jsii proxy for IRolePolicyRef
type jsiiProxy_IRolePolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IRolePolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

