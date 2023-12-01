package awscdkschedulertargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
	"github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/internal"
)

// Use an Amazon SNS topic as a target for AWS EventBridge Scheduler.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   payload := map[string]*string{
//   	"message": jsii.String("Hello scheduler!"),
//   }
//
//   target := targets.NewSnsPublish(topic, &ScheduleTargetBaseProps{
//   	Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(payload),
//   })
//
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
//   	Target: Target,
//   })
//
// Experimental.
type SnsPublish interface {
	ScheduleTargetBase
	awscdkscheduleralpha.IScheduleTarget
	// Experimental.
	TargetArn() *string
	// Experimental.
	AddTargetActionToRole(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole)
	// Create a return a Schedule Target Configuration for the given schedule.
	//
	// Returns: a Schedule Target Configuration.
	// Experimental.
	Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
	// Experimental.
	BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
}

// The jsii proxy struct for SnsPublish
type jsiiProxy_SnsPublish struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awscdkscheduleralphaIScheduleTarget
}

func (j *jsiiProxy_SnsPublish) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewSnsPublish(topic awssns.ITopic, props *ScheduleTargetBaseProps) SnsPublish {
	_init_.Initialize()

	if err := validateNewSnsPublishParameters(topic, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsPublish{}

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.SnsPublish",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsPublish_Override(s SnsPublish, topic awssns.ITopic, props *ScheduleTargetBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.SnsPublish",
		[]interface{}{topic, props},
		s,
	)
}

func (s *jsiiProxy_SnsPublish) AddTargetActionToRole(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole) {
	if err := s.validateAddTargetActionToRoleParameters(schedule, role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addTargetActionToRole",
		[]interface{}{schedule, role},
	)
}

func (s *jsiiProxy_SnsPublish) Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := s.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsPublish) BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := s.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		s,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

