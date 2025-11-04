package awscloudwatchactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatchactions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use an SSM OpsItem action as an Alarm action.
//
// Example:
//   var alarm Alarm
//
//   // Create an OpsItem with specific severity and category when alarm triggers
//   alarm.AddAlarmAction(
//   actions.NewSsmAction(actions.OpsItemSeverity_CRITICAL, actions.OpsItemCategory_PERFORMANCE))
//
type SsmAction interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use an SSM OpsItem action as an alarm action.
	Bind(scope constructs.Construct, alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for SsmAction
type jsiiProxy_SsmAction struct {
	internal.Type__awscloudwatchIAlarmAction
}

func NewSsmAction(severity OpsItemSeverity, category OpsItemCategory) SsmAction {
	_init_.Initialize()

	if err := validateNewSsmActionParameters(severity); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.SsmAction",
		[]interface{}{severity, category},
		&j,
	)

	return &j
}

func NewSsmAction_Override(s SsmAction, severity OpsItemSeverity, category OpsItemCategory) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.SsmAction",
		[]interface{}{severity, category},
		s,
	)
}

func (s *jsiiProxy_SsmAction) Bind(scope constructs.Construct, alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig {
	if err := s.validateBindParameters(scope, alarm); err != nil {
		panic(err)
	}
	var returns *awscloudwatch.AlarmActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, alarm},
		&returns,
	)

	return returns
}

