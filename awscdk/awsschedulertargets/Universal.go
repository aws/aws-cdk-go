package awsschedulertargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsschedulertargets/internal"
)

// Use a wider set of AWS API as a target for AWS EventBridge Scheduler.
//
// Example:
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Cron(&CronOptionsWithTimezone{
//   		Minute: jsii.String("0"),
//   		Hour: jsii.String("0"),
//   	}),
//   	Target: targets.NewUniversal(&UniversalTargetProps{
//   		Service: jsii.String("rds"),
//   		Action: jsii.String("stopDBCluster"),
//   		Input: awscdk.ScheduleTargetInput_FromObject(map[string]*string{
//   			"DbClusterIdentifier": jsii.String("my-db"),
//   		}),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-targets-universal.html
//
type Universal interface {
	ScheduleTargetBase
	awsscheduler.IScheduleTarget
	TargetArn() *string
	AddTargetActionToRole(role awsiam.IRole)
	// Create a return a Schedule Target Configuration for the given schedule.
	//
	// Returns: a Schedule Target Configuration.
	Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig
	BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig
}

// The jsii proxy struct for Universal
type jsiiProxy_Universal struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awsschedulerIScheduleTarget
}

func (j *jsiiProxy_Universal) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


func NewUniversal(props *UniversalTargetProps) Universal {
	_init_.Initialize()

	if err := validateNewUniversalParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Universal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.Universal",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewUniversal_Override(u Universal, props *UniversalTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.Universal",
		[]interface{}{props},
		u,
	)
}

func (u *jsiiProxy_Universal) AddTargetActionToRole(role awsiam.IRole) {
	if err := u.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (u *jsiiProxy_Universal) Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := u.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		u,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_Universal) BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := u.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		u,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

