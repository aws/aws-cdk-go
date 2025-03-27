package awscdkschedulertargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
	"github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/internal"
)

// Use an AWS CodePipeline pipeline as a target for AWS EventBridge Scheduler.
//
// Example:
//   import codepipeline "github.com/aws/aws-cdk-go/awscdk"
//
//   var pipeline pipeline
//
//
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewCodePipelineStartPipelineExecution(pipeline),
//   })
//
// Deprecated.
type CodePipelineStartPipelineExecution interface {
	ScheduleTargetBase
	awscdkscheduleralpha.IScheduleTarget
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

// The jsii proxy struct for CodePipelineStartPipelineExecution
type jsiiProxy_CodePipelineStartPipelineExecution struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awscdkscheduleralphaIScheduleTarget
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


// Deprecated.
func NewCodePipelineStartPipelineExecution(pipeline awscodepipeline.IPipeline, props *ScheduleTargetBaseProps) CodePipelineStartPipelineExecution {
	_init_.Initialize()

	if err := validateNewCodePipelineStartPipelineExecutionParameters(pipeline, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodePipelineStartPipelineExecution{}

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.CodePipelineStartPipelineExecution",
		[]interface{}{pipeline, props},
		&j,
	)

	return &j
}

// Deprecated.
func NewCodePipelineStartPipelineExecution_Override(c CodePipelineStartPipelineExecution, pipeline awscodepipeline.IPipeline, props *ScheduleTargetBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.CodePipelineStartPipelineExecution",
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

func (c *jsiiProxy_CodePipelineStartPipelineExecution) Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := c.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodePipelineStartPipelineExecution) BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := c.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		c,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

