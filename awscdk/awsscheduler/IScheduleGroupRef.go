package awsscheduler

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScheduleGroup.
// Experimental.
type IScheduleGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IScheduleGroupRef
type jsiiProxy_IScheduleGroupRef struct {
	internal.Type__constructsIConstruct
}

