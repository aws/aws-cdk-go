package awsscheduler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsscheduler"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface representing a created or an imported `Schedule`.
type ISchedule interface {
	awscdk.IResource
	interfacesawsscheduler.IScheduleRef
	// The arn of the schedule.
	ScheduleArn() *string
	// The schedule group associated with this schedule.
	ScheduleGroup() IScheduleGroup
	// The name of the schedule.
	ScheduleName() *string
}

// The jsii proxy for ISchedule
type jsiiProxy_ISchedule struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsschedulerIScheduleRef
}

func (i *jsiiProxy_ISchedule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ISchedule) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ISchedule) ScheduleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchedule) ScheduleGroup() IScheduleGroup {
	var returns IScheduleGroup
	_jsii_.Get(
		j,
		"scheduleGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchedule) ScheduleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchedule) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchedule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchedule) ScheduleRef() *interfacesawsscheduler.ScheduleReference {
	var returns *interfacesawsscheduler.ScheduleReference
	_jsii_.Get(
		j,
		"scheduleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchedule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

