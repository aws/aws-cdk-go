package awsschedulertargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsschedulertargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Use an AWS Step function as a target for AWS EventBridge Scheduler.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import tasks "github.com/aws/aws-cdk-go/awscdk"
//
//
//   payload := map[string]*string{
//   	"Name": jsii.String("MyParameter"),
//   	"Value": jsii.String("🌥️"),
//   }
//
//   putParameterStep := tasks.NewCallAwsService(this, jsii.String("PutParameter"), &CallAwsServiceProps{
//   	Service: jsii.String("ssm"),
//   	Action: jsii.String("putParameter"),
//   	IamResources: []*string{
//   		jsii.String("*"),
//   	},
//   	Parameters: map[string]interface{}{
//   		"Name.$": jsii.String("$.Name"),
//   		"Value.$": jsii.String("$.Value"),
//   		"Type": jsii.String("String"),
//   		"Overwrite": jsii.Boolean(true),
//   	},
//   })
//
//   stateMachine := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
//   	DefinitionBody: sfn.DefinitionBody_FromChainable(putParameterStep),
//   })
//
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
//   	Target: targets.NewStepFunctionsStartExecution(stateMachine, &ScheduleTargetBaseProps{
//   		Input: awscdk.ScheduleTargetInput_FromObject(payload),
//   	}),
//   })
//
type StepFunctionsStartExecution interface {
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

// The jsii proxy struct for StepFunctionsStartExecution
type jsiiProxy_StepFunctionsStartExecution struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awsschedulerIScheduleTarget
}

func (j *jsiiProxy_StepFunctionsStartExecution) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


func NewStepFunctionsStartExecution(stateMachine awsstepfunctions.IStateMachine, props *ScheduleTargetBaseProps) StepFunctionsStartExecution {
	_init_.Initialize()

	if err := validateNewStepFunctionsStartExecutionParameters(stateMachine, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StepFunctionsStartExecution{}

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.StepFunctionsStartExecution",
		[]interface{}{stateMachine, props},
		&j,
	)

	return &j
}

func NewStepFunctionsStartExecution_Override(s StepFunctionsStartExecution, stateMachine awsstepfunctions.IStateMachine, props *ScheduleTargetBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.StepFunctionsStartExecution",
		[]interface{}{stateMachine, props},
		s,
	)
}

func (s *jsiiProxy_StepFunctionsStartExecution) AddTargetActionToRole(role awsiam.IRole) {
	if err := s.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (s *jsiiProxy_StepFunctionsStartExecution) Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
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

func (s *jsiiProxy_StepFunctionsStartExecution) BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
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

