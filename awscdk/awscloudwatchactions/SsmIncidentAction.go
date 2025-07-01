package awscloudwatchactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatchactions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use an SSM Incident Response Plan as an Alarm action.
//
// Example:
//   var alarm alarm
//
//   // Create an Incident Manager incident based on a specific response plan
//   alarm.AddAlarmAction(
//   actions.NewSsmIncidentAction(jsii.String("ResponsePlanName")))
//
type SsmIncidentAction interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use an SSM Incident as an alarm action based on an Incident Manager Response Plan.
	Bind(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for SsmIncidentAction
type jsiiProxy_SsmIncidentAction struct {
	internal.Type__awscloudwatchIAlarmAction
}

func NewSsmIncidentAction(responsePlanName *string) SsmIncidentAction {
	_init_.Initialize()

	if err := validateNewSsmIncidentActionParameters(responsePlanName); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmIncidentAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.SsmIncidentAction",
		[]interface{}{responsePlanName},
		&j,
	)

	return &j
}

func NewSsmIncidentAction_Override(s SsmIncidentAction, responsePlanName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.SsmIncidentAction",
		[]interface{}{responsePlanName},
		s,
	)
}

func (s *jsiiProxy_SsmIncidentAction) Bind(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig {
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

