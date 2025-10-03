package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RotationSchedule.
// Experimental.
type IRotationScheduleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRotationScheduleRef
type jsiiProxy_IRotationScheduleRef struct {
	internal.Type__constructsIConstruct
}

