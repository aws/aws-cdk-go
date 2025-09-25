package awssecretsmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RotationSchedule.
// Experimental.
type IRotationScheduleRef interface {
	constructs.IConstruct
	// A reference to a RotationSchedule resource.
	// Experimental.
	RotationScheduleRef() *RotationScheduleReference
}

// The jsii proxy for IRotationScheduleRef
type jsiiProxy_IRotationScheduleRef struct {
	internal.Type__constructsIConstruct
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

