package awscdkschedulertargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
	"github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/internal"
)

// Use an AWS CodeBuild as a target for AWS EventBridge Scheduler.
//
// Example:
//   import codebuild "github.com/aws/aws-cdk-go/awscdk"
//
//   var project project
//
//
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewCodeBuildStartBuild(project),
//   })
//
// Experimental.
type CodeBuildStartBuild interface {
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

// The jsii proxy struct for CodeBuildStartBuild
type jsiiProxy_CodeBuildStartBuild struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awscdkscheduleralphaIScheduleTarget
}

func (j *jsiiProxy_CodeBuildStartBuild) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewCodeBuildStartBuild(project awscodebuild.IProject, props *ScheduleTargetBaseProps) CodeBuildStartBuild {
	_init_.Initialize()

	if err := validateNewCodeBuildStartBuildParameters(project, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodeBuildStartBuild{}

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.CodeBuildStartBuild",
		[]interface{}{project, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCodeBuildStartBuild_Override(c CodeBuildStartBuild, project awscodebuild.IProject, props *ScheduleTargetBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.CodeBuildStartBuild",
		[]interface{}{project, props},
		c,
	)
}

func (c *jsiiProxy_CodeBuildStartBuild) AddTargetActionToRole(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole) {
	if err := c.validateAddTargetActionToRoleParameters(schedule, role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addTargetActionToRole",
		[]interface{}{schedule, role},
	)
}

func (c *jsiiProxy_CodeBuildStartBuild) Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
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

func (c *jsiiProxy_CodeBuildStartBuild) BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
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

