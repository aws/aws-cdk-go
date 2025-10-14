package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LaunchTemplateConstraint.
// Experimental.
type ILaunchTemplateConstraintRef interface {
	constructs.IConstruct
	// A reference to a LaunchTemplateConstraint resource.
	// Experimental.
	LaunchTemplateConstraintRef() *LaunchTemplateConstraintReference
}

// The jsii proxy for ILaunchTemplateConstraintRef
type jsiiProxy_ILaunchTemplateConstraintRef struct {
	internal.Type__constructsIConstruct
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

