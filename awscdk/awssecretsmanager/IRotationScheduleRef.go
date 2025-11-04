package awssecretsmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RotationSchedule.
// Experimental.
type IRotationScheduleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RotationSchedule resource.
	// Experimental.
	RotationScheduleRef() *RotationScheduleReference
}

// The jsii proxy for IRotationScheduleRef
type jsiiProxy_IRotationScheduleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IRotationScheduleRef) RotationScheduleRef() *RotationScheduleReference {
	var returns *RotationScheduleReference
	_jsii_.Get(
		j,
		"rotationScheduleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRotationScheduleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRotationScheduleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

