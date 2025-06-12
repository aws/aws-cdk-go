package awsschedulertargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsschedulertargets/internal"
)

// Schedule an ECS Task using AWS EventBridge Scheduler.
type EcsRunTask interface {
	ScheduleTargetBase
	awsscheduler.IScheduleTarget
	Cluster() awsecs.ICluster
	Props() *EcsRunTaskBaseProps
	TargetArn() *string
	AddTargetActionToRole(role awsiam.IRole)
	// Create a return a Schedule Target Configuration for the given schedule.
	//
	// Returns: a Schedule Target Configuration.
	Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig
	BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig
}

// The jsii proxy struct for EcsRunTask
type jsiiProxy_EcsRunTask struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awsschedulerIScheduleTarget
}

func (j *jsiiProxy_EcsRunTask) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) Props() *EcsRunTaskBaseProps {
	var returns *EcsRunTaskBaseProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


func NewEcsRunTask_Override(e EcsRunTask, cluster awsecs.ICluster, props *EcsRunTaskBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.EcsRunTask",
		[]interface{}{cluster, props},
		e,
	)
}

func (e *jsiiProxy_EcsRunTask) AddTargetActionToRole(role awsiam.IRole) {
	if err := e.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (e *jsiiProxy_EcsRunTask) Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := e.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsRunTask) BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := e.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		e,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

