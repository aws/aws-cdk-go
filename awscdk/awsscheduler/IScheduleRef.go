package awsscheduler

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Schedule.
// Experimental.
type IScheduleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IScheduleRef
type jsiiProxy_IScheduleRef struct {
	internal.Type__constructsIConstruct
}

