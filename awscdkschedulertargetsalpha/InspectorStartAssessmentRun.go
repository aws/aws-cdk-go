package awscdkschedulertargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspector"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
	"github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/internal"
)

// Use an Amazon Inspector as a target for AWS EventBridge Scheduler.
//
// Example:
//   import inspector "github.com/aws/aws-cdk-go/awscdk"
//
//   var assessmentTemplate cfnAssessmentTemplate
//
//
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewInspectorStartAssessmentRun(assessmentTemplate),
//   })
//
// Experimental.
type InspectorStartAssessmentRun interface {
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

// The jsii proxy struct for InspectorStartAssessmentRun
type jsiiProxy_InspectorStartAssessmentRun struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awscdkscheduleralphaIScheduleTarget
}

func (j *jsiiProxy_InspectorStartAssessmentRun) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewInspectorStartAssessmentRun(template awsinspector.CfnAssessmentTemplate, props *ScheduleTargetBaseProps) InspectorStartAssessmentRun {
	_init_.Initialize()

	if err := validateNewInspectorStartAssessmentRunParameters(template, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_InspectorStartAssessmentRun{}

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.InspectorStartAssessmentRun",
		[]interface{}{template, props},
		&j,
	)

	return &j
}

// Experimental.
func NewInspectorStartAssessmentRun_Override(i InspectorStartAssessmentRun, template awsinspector.CfnAssessmentTemplate, props *ScheduleTargetBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.InspectorStartAssessmentRun",
		[]interface{}{template, props},
		i,
	)
}

func (i *jsiiProxy_InspectorStartAssessmentRun) AddTargetActionToRole(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole) {
	if err := i.validateAddTargetActionToRoleParameters(schedule, role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addTargetActionToRole",
		[]interface{}{schedule, role},
	)
}

func (i *jsiiProxy_InspectorStartAssessmentRun) Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := i.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InspectorStartAssessmentRun) BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := i.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		i,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

