package awsschedulertargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspector"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsschedulertargets/internal"
)

// Use an Amazon Inspector as a target for AWS EventBridge Scheduler.
//
// Example:
//   import inspector "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnAssessmentTemplate cfnAssessmentTemplate
//
//
//   assessmentTemplate := inspector.AssessmentTemplate_FromCfnAssessmentTemplate(this, jsii.String("MyAssessmentTemplate"), cfnAssessmentTemplate)
//
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewInspectorStartAssessmentRun(assessmentTemplate),
//   })
//
type InspectorStartAssessmentRun interface {
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

// The jsii proxy struct for InspectorStartAssessmentRun
type jsiiProxy_InspectorStartAssessmentRun struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awsschedulerIScheduleTarget
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


func NewInspectorStartAssessmentRun(template awsinspector.IAssessmentTemplate, props *ScheduleTargetBaseProps) InspectorStartAssessmentRun {
	_init_.Initialize()

	if err := validateNewInspectorStartAssessmentRunParameters(template, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_InspectorStartAssessmentRun{}

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.InspectorStartAssessmentRun",
		[]interface{}{template, props},
		&j,
	)

	return &j
}

func NewInspectorStartAssessmentRun_Override(i InspectorStartAssessmentRun, template awsinspector.IAssessmentTemplate, props *ScheduleTargetBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.InspectorStartAssessmentRun",
		[]interface{}{template, props},
		i,
	)
}

func (i *jsiiProxy_InspectorStartAssessmentRun) AddTargetActionToRole(role awsiam.IRole) {
	if err := i.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (i *jsiiProxy_InspectorStartAssessmentRun) Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := i.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InspectorStartAssessmentRun) BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := i.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		i,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

