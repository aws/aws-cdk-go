package awscdkschedulertargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
	"github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/internal"
)

// Use a SageMaker pipeline as a target for AWS EventBridge Scheduler.
//
// Example:
//   import sagemaker "github.com/aws/aws-cdk-go/awscdk"
//
//   var pipeline iPipeline
//
//
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewSageMakerStartPipelineExecution(pipeline, &SageMakerStartPipelineExecutionProps{
//   		PipelineParameterList: []sageMakerPipelineParameter{
//   			&sageMakerPipelineParameter{
//   				Name: jsii.String("parameter-name"),
//   				Value: jsii.String("parameter-value"),
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type SageMakerStartPipelineExecution interface {
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
	BindBaseTargetConfig(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
}

// The jsii proxy struct for SageMakerStartPipelineExecution
type jsiiProxy_SageMakerStartPipelineExecution struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awscdkscheduleralphaIScheduleTarget
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


// Experimental.
func NewSageMakerStartPipelineExecution(pipeline awssagemaker.IPipeline, props *SageMakerStartPipelineExecutionProps) SageMakerStartPipelineExecution {
	_init_.Initialize()

	if err := validateNewSageMakerStartPipelineExecutionParameters(pipeline, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SageMakerStartPipelineExecution{}

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.SageMakerStartPipelineExecution",
		[]interface{}{pipeline, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerStartPipelineExecution_Override(s SageMakerStartPipelineExecution, pipeline awssagemaker.IPipeline, props *SageMakerStartPipelineExecutionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.SageMakerStartPipelineExecution",
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

func (s *jsiiProxy_SageMakerStartPipelineExecution) Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
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

func (s *jsiiProxy_SageMakerStartPipelineExecution) BindBaseTargetConfig(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := s.validateBindBaseTargetConfigParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		s,
		"bindBaseTargetConfig",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

