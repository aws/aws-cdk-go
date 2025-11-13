package interfacesawsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LaunchRoleConstraint.
// Experimental.
type ILaunchRoleConstraintRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LaunchRoleConstraint resource.
	// Experimental.
	LaunchRoleConstraintRef() *LaunchRoleConstraintReference
}

// The jsii proxy for ILaunchRoleConstraintRef
type jsiiProxy_ILaunchRoleConstraintRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ILaunchRoleConstraintRef) LaunchRoleConstraintRef() *LaunchRoleConstraintReference {
	var returns *LaunchRoleConstraintReference
	_jsii_.Get(
		j,
		"launchRoleConstraintRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchRoleConstraintRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchRoleConstraintRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

