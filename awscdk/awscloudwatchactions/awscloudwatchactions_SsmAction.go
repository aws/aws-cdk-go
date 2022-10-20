package awscloudwatchactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatchactions/internal"
)

// Use an SSM OpsItem action as an Alarm action.
//
// Example:
//   var alarm alarm
//
//   // Create an OpsItem with specific severity and category when alarm triggers
//   alarm.addAlarmAction(
//   actions.NewSsmAction(actions.opsItemSeverity_CRITICAL, actions.opsItemCategory_PERFORMANCE))
//
// Experimental.
type SsmAction interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use an SSM OpsItem action as an alarm action.
	// Experimental.
	Bind(_scope awscdk.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for SsmAction
type jsiiProxy_SsmAction struct {
	internal.Type__awscloudwatchIAlarmAction
}

// Experimental.
func NewSsmAction(severity OpsItemSeverity, category OpsItemCategory) SsmAction {
	_init_.Initialize()

	if err := validateNewSsmActionParameters(severity); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmAction{}

	_jsii_.Create(
		"monocdk.aws_cloudwatch_actions.SsmAction",
		[]interface{}{severity, category},
		&j,
	)

	return &j
}

// Experimental.
func NewSsmAction_Override(s SsmAction, severity OpsItemSeverity, category OpsItemCategory) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudwatch_actions.SsmAction",
		[]interface{}{severity, category},
		s,
	)
}

func (s *jsiiProxy_SsmAction) Bind(_scope awscdk.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig {
	if err := s.validateBindParameters(_scope, _alarm); err != nil {
		panic(err)
	}
	var returns *awscloudwatch.AlarmActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _alarm},
		&returns,
	)

	return returns
}

