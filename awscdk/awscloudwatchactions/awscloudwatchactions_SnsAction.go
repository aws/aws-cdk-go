package awscloudwatchactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatchactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
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
// Experimental.
type SnsAction interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use an SNS topic as an alarm action.
	// Experimental.
	Bind(_scope awscdk.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for SnsAction
type jsiiProxy_SnsAction struct {
	internal.Type__awscloudwatchIAlarmAction
}

// Experimental.
func NewSnsAction(topic awssns.ITopic) SnsAction {
	_init_.Initialize()

	if err := validateNewSnsActionParameters(topic); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsAction{}

	_jsii_.Create(
		"monocdk.aws_cloudwatch_actions.SnsAction",
		[]interface{}{topic},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsAction_Override(s SnsAction, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudwatch_actions.SnsAction",
		[]interface{}{topic},
		s,
	)
}

func (s *jsiiProxy_SnsAction) Bind(_scope awscdk.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig {
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

