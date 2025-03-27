package awsschedulertargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsschedulertargets/internal"
)

// Use an AWS Lambda function as a target for AWS EventBridge Scheduler.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String("exports.handler = handler.toString()")),
//   })
//
//   dlq := sqs.NewQueue(this, jsii.String("DLQ"), &QueueProps{
//   	QueueName: jsii.String("MyDLQ"),
//   })
//
//   target := targets.NewLambdaInvoke(fn, &ScheduleTargetBaseProps{
//   	DeadLetterQueue: dlq,
//   	MaxEventAge: awscdk.Duration_Minutes(jsii.Number(1)),
//   	RetryAttempts: jsii.Number(3),
//   	Input: awscdk.ScheduleTargetInput_FromObject(map[string]*string{
//   		"payload": jsii.String("useful"),
//   	}),
//   })
//
//   schedule := awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
//   	Target: Target,
//   })
//
type LambdaInvoke interface {
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

// The jsii proxy struct for LambdaInvoke
type jsiiProxy_LambdaInvoke struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awsschedulerIScheduleTarget
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


func NewLambdaInvoke(func_ awslambda.IFunction, props *ScheduleTargetBaseProps) LambdaInvoke {
	_init_.Initialize()

	if err := validateNewLambdaInvokeParameters(func_, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaInvoke{}

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.LambdaInvoke",
		[]interface{}{func_, props},
		&j,
	)

	return &j
}

func NewLambdaInvoke_Override(l LambdaInvoke, func_ awslambda.IFunction, props *ScheduleTargetBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.LambdaInvoke",
		[]interface{}{func_, props},
		l,
	)
}

func (l *jsiiProxy_LambdaInvoke) AddTargetActionToRole(role awsiam.IRole) {
	if err := l.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (l *jsiiProxy_LambdaInvoke) Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := l.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := l.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		l,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

