package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LaunchNotificationConstraint.
// Experimental.
type ILaunchNotificationConstraintRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LaunchNotificationConstraint resource.
	// Experimental.
	LaunchNotificationConstraintRef() *LaunchNotificationConstraintReference
}

// The jsii proxy for ILaunchNotificationConstraintRef
type jsiiProxy_ILaunchNotificationConstraintRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILaunchNotificationConstraintRef) LaunchNotificationConstraintRef() *LaunchNotificationConstraintReference {
	var returns *LaunchNotificationConstraintReference
	_jsii_.Get(
		j,
		"launchNotificationConstraintRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchNotificationConstraintRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchNotificationConstraintRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

