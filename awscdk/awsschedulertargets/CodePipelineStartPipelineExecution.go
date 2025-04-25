package awsschedulertargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsschedulertargets/internal"
)

// Use an AWS CodePipeline pipeline as a target for AWS EventBridge Scheduler.
//
// Example:
//   import codepipeline "github.com/aws/aws-cdk-go/awscdk"
//
//   var pipeline pipeline
//
//
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewCodePipelineStartPipelineExecution(pipeline),
//   })
//
type CodePipelineStartPipelineExecution interface {
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

// The jsii proxy struct for CodePipelineStartPipelineExecution
type jsiiProxy_CodePipelineStartPipelineExecution struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awsschedulerIScheduleTarget
}

func (j *jsiiProxy_CodePipelineStartPipelineExecution) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


func NewCodePipelineStartPipelineExecution(pipeline awscodepipeline.IPipeline, props *ScheduleTargetBaseProps) CodePipelineStartPipelineExecution {
	_init_.Initialize()

	if err := validateNewCodePipelineStartPipelineExecutionParameters(pipeline, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodePipelineStartPipelineExecution{}

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.CodePipelineStartPipelineExecution",
		[]interface{}{pipeline, props},
		&j,
	)

	return &j
}

func NewCodePipelineStartPipelineExecution_Override(c CodePipelineStartPipelineExecution, pipeline awscodepipeline.IPipeline, props *ScheduleTargetBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.CodePipelineStartPipelineExecution",
		[]interface{}{pipeline, props},
		c,
	)
}

func (c *jsiiProxy_CodePipelineStartPipelineExecution) AddTargetActionToRole(role awsiam.IRole) {
	if err := c.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (c *jsiiProxy_CodePipelineStartPipelineExecution) Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := c.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodePipelineStartPipelineExecution) BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := c.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		c,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

