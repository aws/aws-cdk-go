package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LaunchTemplateConstraint.
// Experimental.
type ILaunchTemplateConstraintRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LaunchTemplateConstraint resource.
	// Experimental.
	LaunchTemplateConstraintRef() *LaunchTemplateConstraintReference
}

// The jsii proxy for ILaunchTemplateConstraintRef
type jsiiProxy_ILaunchTemplateConstraintRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILaunchTemplateConstraintRef) LaunchTemplateConstraintRef() *LaunchTemplateConstraintReference {
	var returns *LaunchTemplateConstraintReference
	_jsii_.Get(
		j,
		"launchTemplateConstraintRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchTemplateConstraintRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchTemplateConstraintRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

