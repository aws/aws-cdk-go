package awscloudwatchactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatchactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use a Lambda action as an Alarm action.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var alarm Alarm
//   var fn Function
//   var alias Alias
//   var version Version
//
//
//   // Attach a Lambda Function when alarm triggers
//   alarm.AddAlarmAction(
//   actions.NewLambdaAction(fn))
//
//   // Attach a Lambda Function Alias when alarm triggers
//   alarm.AddAlarmAction(
//   actions.NewLambdaAction(alias))
//
//   // Attach a Lambda Function version when alarm triggers
//   alarm.AddAlarmAction(
//   actions.NewLambdaAction(version))
//
type LambdaAction interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use a Lambda action as an alarm action.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html
	//
	Bind(scope constructs.Construct, alarm interfacesawscloudwatch.IAlarmRef) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for LambdaAction
type jsiiProxy_LambdaAction struct {
	internal.Type__awscloudwatchIAlarmAction
}

func NewLambdaAction(lambdaFunction interface{}, props *LambdaActionProps) LambdaAction {
	_init_.Initialize()

	if err := validateNewLambdaActionParameters(lambdaFunction, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.LambdaAction",
		[]interface{}{lambdaFunction, props},
		&j,
	)

	return &j
}

func NewLambdaAction_Override(l LambdaAction, lambdaFunction interface{}, props *LambdaActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.LambdaAction",
		[]interface{}{lambdaFunction, props},
		l,
	)
}

func (l *jsiiProxy_LambdaAction) Bind(scope constructs.Construct, alarm interfacesawscloudwatch.IAlarmRef) *awscloudwatch.AlarmActionConfig {
	if err := l.validateBindParameters(scope, alarm); err != nil {
		panic(err)
	}
	var returns *awscloudwatch.AlarmActionConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, alarm},
		&returns,
	)

	return returns
}

