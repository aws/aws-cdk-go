package awsschedulertargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
)

// Base class for Schedule Targets.
type ScheduleTargetBase interface {
	TargetArn() *string
	AddTargetActionToRole(role awsiam.IRole)
	// Create a return a Schedule Target Configuration for the given schedule.
	//
	// Returns: a Schedule Target Configuration.
	Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig
	BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig
}

// The jsii proxy struct for ScheduleTargetBase
type jsiiProxy_ScheduleTargetBase struct {
	_ byte // padding
}

func (j *jsiiProxy_ScheduleTargetBase) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


func NewScheduleTargetBase_Override(s ScheduleTargetBase, baseProps *ScheduleTargetBaseProps, targetArn *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.ScheduleTargetBase",
		[]interface{}{baseProps, targetArn},
		s,
	)
}

func (s *jsiiProxy_ScheduleTargetBase) AddTargetActionToRole(role awsiam.IRole) {
	if err := s.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (s *jsiiProxy_ScheduleTargetBase) Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := s.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleTargetBase) BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := s.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		s,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

