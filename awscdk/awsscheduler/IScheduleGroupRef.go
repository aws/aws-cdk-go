package awsscheduler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScheduleGroup.
// Experimental.
type IScheduleGroupRef interface {
	constructs.IConstruct
	// A reference to a ScheduleGroup resource.
	// Experimental.
	ScheduleGroupRef() *ScheduleGroupReference
}

// The jsii proxy for IScheduleGroupRef
type jsiiProxy_IScheduleGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IScheduleGroupRef) ScheduleGroupRef() *ScheduleGroupReference {
	var returns *ScheduleGroupReference
	_jsii_.Get(
		j,
		"scheduleGroupRef",
		&returns,
	)
	return returns
}

