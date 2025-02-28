package awscdkschedulertargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
	"github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/internal"
)

// Use a wider set of AWS API as a target for AWS EventBridge Scheduler.
//
// Example:
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Cron(&CronOptionsWithTimezone{
//   		Minute: jsii.String("0"),
//   		Hour: jsii.String("0"),
//   	}),
//   	Target: targets.NewUniversal(&UniversalTargetProps{
//   		Service: jsii.String("rds"),
//   		Action: jsii.String("stopDBCluster"),
//   		Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
//   			"DbClusterIdentifier": jsii.String("my-db"),
//   		}),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-targets-universal.html
//
// Experimental.
type Universal interface {
	ScheduleTargetBase
	awscdkscheduleralpha.IScheduleTarget
	// Experimental.
	TargetArn() *string
	// Experimental.
	AddTargetActionToRole(role awsiam.IRole)
	// Create a return a Schedule Target Configuration for the given schedule.
	//
	// Returns: a Schedule Target Configuration.
	// Experimental.
	Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
	// Experimental.
	BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
}

// The jsii proxy struct for Universal
type jsiiProxy_Universal struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awscdkscheduleralphaIScheduleTarget
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


// Experimental.
func NewUniversal(props *UniversalTargetProps) Universal {
	_init_.Initialize()

	if err := validateNewUniversalParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Universal{}

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.Universal",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewUniversal_Override(u Universal, props *UniversalTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.Universal",
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

func (u *jsiiProxy_Universal) Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := u.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		u,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_Universal) BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := u.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		u,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

