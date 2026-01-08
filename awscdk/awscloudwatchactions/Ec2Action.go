package awscloudwatchactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatchactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use an EC2 action as an Alarm action.
//
// Example:
//   // Alarm must be configured with an EC2 per-instance metric
//   var alarm Alarm
//
//   // Attach a reboot when alarm triggers
//   alarm.AddAlarmAction(
//   actions.NewEc2Action(actions.Ec2InstanceAction_REBOOT))
//
type Ec2Action interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use an EC2 action as an alarm action.
	Bind(scope constructs.Construct, alarm interfacesawscloudwatch.IAlarmRef) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for Ec2Action
type jsiiProxy_Ec2Action struct {
	internal.Type__awscloudwatchIAlarmAction
}

func NewEc2Action(instanceAction Ec2InstanceAction) Ec2Action {
	_init_.Initialize()

	if err := validateNewEc2ActionParameters(instanceAction); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2Action{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.Ec2Action",
		[]interface{}{instanceAction},
		&j,
	)

	return &j
}

func NewEc2Action_Override(e Ec2Action, instanceAction Ec2InstanceAction) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.Ec2Action",
		[]interface{}{instanceAction},
		e,
	)
}

func (e *jsiiProxy_Ec2Action) Bind(scope constructs.Construct, alarm interfacesawscloudwatch.IAlarmRef) *awscloudwatch.AlarmActionConfig {
	if err := e.validateBindParameters(scope, alarm); err != nil {
		panic(err)
	}
	var returns *awscloudwatch.AlarmActionConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{scope, alarm},
		&returns,
	)

	return returns
}

