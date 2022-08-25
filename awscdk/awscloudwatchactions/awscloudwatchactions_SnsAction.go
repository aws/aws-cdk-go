package awscloudwatchactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatchactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use an SNS topic as an alarm action.
//
// Example:
//   import cw_actions "github.com/aws/aws-cdk-go/awscdk"
//   var alarm alarm
//
//
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//   alarm.addAlarmAction(cw_actions.NewSnsAction(topic))
//
type SnsAction interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use an SNS topic as an alarm action.
	Bind(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for SnsAction
type jsiiProxy_SnsAction struct {
	internal.Type__awscloudwatchIAlarmAction
}

func NewSnsAction(topic awssns.ITopic) SnsAction {
	_init_.Initialize()

	j := jsiiProxy_SnsAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.SnsAction",
		[]interface{}{topic},
		&j,
	)

	return &j
}

func NewSnsAction_Override(s SnsAction, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.SnsAction",
		[]interface{}{topic},
		s,
	)
}

func (s *jsiiProxy_SnsAction) Bind(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig {
	var returns *awscloudwatch.AlarmActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _alarm},
		&returns,
	)

	return returns
}

