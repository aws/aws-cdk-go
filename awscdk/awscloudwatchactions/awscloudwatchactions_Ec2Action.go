package awscloudwatchactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatchactions/internal"
)

// Use an EC2 action as an Alarm action.
//
// Example:
//   // Alarm must be configured with an EC2 per-instance metric
//   var alarm alarm
//
//   // Attach a reboot when alarm triggers
//   alarm.addAlarmAction(
//   actions.NewEc2Action(actions.ec2InstanceAction_REBOOT))
//
// Experimental.
type Ec2Action interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use an EC2 action as an alarm action.
	// Experimental.
	Bind(_scope awscdk.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for Ec2Action
type jsiiProxy_Ec2Action struct {
	internal.Type__awscloudwatchIAlarmAction
}

// Experimental.
func NewEc2Action(instanceAction Ec2InstanceAction) Ec2Action {
	_init_.Initialize()

	if err := validateNewEc2ActionParameters(instanceAction); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2Action{}

	_jsii_.Create(
		"monocdk.aws_cloudwatch_actions.Ec2Action",
		[]interface{}{instanceAction},
		&j,
	)

	return &j
}

// Experimental.
func NewEc2Action_Override(e Ec2Action, instanceAction Ec2InstanceAction) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudwatch_actions.Ec2Action",
		[]interface{}{instanceAction},
		e,
	)
}

func (e *jsiiProxy_Ec2Action) Bind(_scope awscdk.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig {
	if err := e.validateBindParameters(_scope, _alarm); err != nil {
		panic(err)
	}
	var returns *awscloudwatch.AlarmActionConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_scope, _alarm},
		&returns,
	)

	return returns
}

