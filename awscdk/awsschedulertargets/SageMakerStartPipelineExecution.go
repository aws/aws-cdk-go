package awsschedulertargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsschedulertargets/internal"
)

// Use a SageMaker pipeline as a target for AWS EventBridge Scheduler.
//
// Example:
//   import sagemaker "github.com/aws/aws-cdk-go/awscdk"
//
//   var pipeline IPipeline
//
//
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewSageMakerStartPipelineExecution(pipeline, &SageMakerStartPipelineExecutionProps{
//   		PipelineParameterList: []SageMakerPipelineParameter{
//   			&SageMakerPipelineParameter{
//   				Name: jsii.String("parameter-name"),
//   				Value: jsii.String("parameter-value"),
//   			},
//   		},
//   	}),
//   })
//
type SageMakerStartPipelineExecution interface {
	ScheduleTargetBase
	awsscheduler.IScheduleTarget
	TargetArn() *string
	AddTargetActionToRole(role awsiam.IRole)
	// Create a return a Schedule Target Configuration for the given schedule.
	//
	// Returns: a Schedule Target Configuration.
	Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig
	BindBaseTargetConfig(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig
}

// The jsii proxy struct for SageMakerStartPipelineExecution
type jsiiProxy_SageMakerStartPipelineExecution struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awsschedulerIScheduleTarget
}

func (j *jsiiProxy_SageMakerStartPipelineExecution) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


func NewSageMakerStartPipelineExecution(pipeline awssagemaker.IPipeline, props *SageMakerStartPipelineExecutionProps) SageMakerStartPipelineExecution {
	_init_.Initialize()

	if err := validateNewSageMakerStartPipelineExecutionParameters(pipeline, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SageMakerStartPipelineExecution{}

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.SageMakerStartPipelineExecution",
		[]interface{}{pipeline, props},
		&j,
	)

	return &j
}

func NewSageMakerStartPipelineExecution_Override(s SageMakerStartPipelineExecution, pipeline awssagemaker.IPipeline, props *SageMakerStartPipelineExecutionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.SageMakerStartPipelineExecution",
		[]interface{}{pipeline, props},
		s,
	)
}

func (s *jsiiProxy_SageMakerStartPipelineExecution) AddTargetActionToRole(role awsiam.IRole) {
	if err := s.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (s *jsiiProxy_SageMakerStartPipelineExecution) Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
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

func (s *jsiiProxy_SageMakerStartPipelineExecution) BindBaseTargetConfig(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := s.validateBindBaseTargetConfigParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		s,
		"bindBaseTargetConfig",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

