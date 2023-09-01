package awscdkschedulertargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
	"github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/internal"
)

// Example:
//   var fn function
//
//
//   target := targets.NewLambdaInvoke(fn, &ScheduleTargetBaseProps{
//   	Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
//   		"payload": jsii.String("useful"),
//   	}),
//   })
//
//   schedule := awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(10))),
//   	Target: Target,
//   	Description: jsii.String("This is a test schedule that invokes lambda function every 10 minutes."),
//   })
//
// Experimental.
type LambdaInvoke interface {
	ScheduleTargetBase
	awscdkscheduleralpha.IScheduleTarget
	// Experimental.
	TargetArn() *string
	// Experimental.
	AddTargetActionToRole(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole)
	// Experimental.
	Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
	// Experimental.
	BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
}

// The jsii proxy struct for LambdaInvoke
type jsiiProxy_LambdaInvoke struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awscdkscheduleralphaIScheduleTarget
}

func (j *jsiiProxy_LambdaInvoke) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaInvoke(func_ awslambda.IFunction, props *ScheduleTargetBaseProps) LambdaInvoke {
	_init_.Initialize()

	if err := validateNewLambdaInvokeParameters(func_, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaInvoke{}

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.LambdaInvoke",
		[]interface{}{func_, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaInvoke_Override(l LambdaInvoke, func_ awslambda.IFunction, props *ScheduleTargetBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.LambdaInvoke",
		[]interface{}{func_, props},
		l,
	)
}

func (l *jsiiProxy_LambdaInvoke) AddTargetActionToRole(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole) {
	if err := l.validateAddTargetActionToRoleParameters(schedule, role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addTargetActionToRole",
		[]interface{}{schedule, role},
	)
}

func (l *jsiiProxy_LambdaInvoke) Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := l.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := l.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		l,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

