package awsschedulertargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsschedulertargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
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
//   	Input: awscdk.ScheduleTargetInput_FromObject(payload),
//   })
//
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
//   	Target: Target,
//   })
//
type SnsPublish interface {
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

// The jsii proxy struct for SnsPublish
type jsiiProxy_SnsPublish struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awsschedulerIScheduleTarget
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


func NewSnsPublish(topic awssns.ITopic, props *ScheduleTargetBaseProps) SnsPublish {
	_init_.Initialize()

	if err := validateNewSnsPublishParameters(topic, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsPublish{}

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.SnsPublish",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

func NewSnsPublish_Override(s SnsPublish, topic awssns.ITopic, props *ScheduleTargetBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.SnsPublish",
		[]interface{}{topic, props},
		s,
	)
}

func (s *jsiiProxy_SnsPublish) AddTargetActionToRole(role awsiam.IRole) {
	if err := s.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (s *jsiiProxy_SnsPublish) Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
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

func (s *jsiiProxy_SnsPublish) BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
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

