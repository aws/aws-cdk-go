package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LaunchRoleConstraint.
// Experimental.
type ILaunchRoleConstraintRef interface {
	constructs.IConstruct
	// A reference to a LaunchRoleConstraint resource.
	// Experimental.
	LaunchRoleConstraintRef() *LaunchRoleConstraintReference
}

// The jsii proxy for ILaunchRoleConstraintRef
type jsiiProxy_ILaunchRoleConstraintRef struct {
	internal.Type__constructsIConstruct
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

