package awscdkschedulertargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
)

// Schedule an ECS Task on EC2 using AWS EventBridge Scheduler.
//
// Example:
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster iCluster
//   var taskDefinition ec2TaskDefinition
//
//
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(cdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewEcsRunEc2Task(cluster, &Ec2TaskProps{
//   		TaskDefinition: *TaskDefinition,
//   	}),
//   })
//
// Deprecated.
type EcsRunEc2Task interface {
	EcsRunTask
	// Deprecated.
	Cluster() awsecs.ICluster
	// Deprecated.
	Props() *EcsRunTaskBaseProps
	// Deprecated.
	TargetArn() *string
	// Deprecated.
	AddTargetActionToRole(role awsiam.IRole)
	// Create a return a Schedule Target Configuration for the given schedule.
	//
	// Returns: a Schedule Target Configuration.
	// Deprecated.
	Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
	// Deprecated.
	BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
}

// The jsii proxy struct for EcsRunEc2Task
type jsiiProxy_EcsRunEc2Task struct {
	jsiiProxy_EcsRunTask
}

func (j *jsiiProxy_EcsRunEc2Task) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunEc2Task) Props() *EcsRunTaskBaseProps {
	var returns *EcsRunTaskBaseProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunEc2Task) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Deprecated.
func NewEcsRunEc2Task(cluster awsecs.ICluster, props *Ec2TaskProps) EcsRunEc2Task {
	_init_.Initialize()

	if err := validateNewEcsRunEc2TaskParameters(cluster, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsRunEc2Task{}

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.EcsRunEc2Task",
		[]interface{}{cluster, props},
		&j,
	)

	return &j
}

// Deprecated.
func NewEcsRunEc2Task_Override(e EcsRunEc2Task, cluster awsecs.ICluster, props *Ec2TaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.EcsRunEc2Task",
		[]interface{}{cluster, props},
		e,
	)
}

func (e *jsiiProxy_EcsRunEc2Task) AddTargetActionToRole(role awsiam.IRole) {
	if err := e.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (e *jsiiProxy_EcsRunEc2Task) Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := e.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsRunEc2Task) BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := e.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		e,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

