package awsschedulertargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
)

// Schedule an ECS Task on Fargate using AWS EventBridge Scheduler.
//
// Example:
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster ICluster
//   var taskDefinition FargateTaskDefinition
//
//
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(cdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewEcsRunFargateTask(cluster, &FargateTaskProps{
//   		TaskDefinition: *TaskDefinition,
//   	}),
//   })
//
type EcsRunFargateTask interface {
	EcsRunTask
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

// The jsii proxy struct for EcsRunFargateTask
type jsiiProxy_EcsRunFargateTask struct {
	jsiiProxy_EcsRunTask
}

func (j *jsiiProxy_EcsRunFargateTask) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunFargateTask) Props() *EcsRunTaskBaseProps {
	var returns *EcsRunTaskBaseProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunFargateTask) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


func NewEcsRunFargateTask(cluster awsecs.ICluster, props *FargateTaskProps) EcsRunFargateTask {
	_init_.Initialize()

	if err := validateNewEcsRunFargateTaskParameters(cluster, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsRunFargateTask{}

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.EcsRunFargateTask",
		[]interface{}{cluster, props},
		&j,
	)

	return &j
}

func NewEcsRunFargateTask_Override(e EcsRunFargateTask, cluster awsecs.ICluster, props *FargateTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.EcsRunFargateTask",
		[]interface{}{cluster, props},
		e,
	)
}

func (e *jsiiProxy_EcsRunFargateTask) AddTargetActionToRole(role awsiam.IRole) {
	if err := e.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (e *jsiiProxy_EcsRunFargateTask) Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
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

func (e *jsiiProxy_EcsRunFargateTask) BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
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

