package awsefs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MountTarget.
// Experimental.
type IMountTargetRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a MountTarget resource.
	// Experimental.
	MountTargetRef() *MountTargetReference
}

// The jsii proxy for IMountTargetRef
type jsiiProxy_IMountTargetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMountTargetRef) MountTargetRef() *MountTargetReference {
	var returns *MountTargetReference
	_jsii_.Get(
		j,
		"mountTargetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMountTargetRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMountTargetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

