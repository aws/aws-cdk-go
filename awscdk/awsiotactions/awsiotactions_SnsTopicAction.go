package awsiotactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/awsiotactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

// The action to write the data from an MQTT message to an Amazon SNS topic.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewSnsTopicAction(topic, &snsTopicActionProps{
//   			messageFormat: actions.snsActionMessageFormat_JSON,
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/iot/latest/developerguide/sns-rule-action.html
//
// Experimental.
type SnsTopicAction interface {
	awsiot.IAction
	// Returns the topic rule action specification.
	// Experimental.
	Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for SnsTopicAction
type jsiiProxy_SnsTopicAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewSnsTopicAction(topic awssns.ITopic, props *SnsTopicActionProps) SnsTopicAction {
	_init_.Initialize()

	j := jsiiProxy_SnsTopicAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.SnsTopicAction",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsTopicAction_Override(s SnsTopicAction, topic awssns.ITopic, props *SnsTopicActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.SnsTopicAction",
		[]interface{}{topic, props},
		s,
	)
}

func (s *jsiiProxy_SnsTopicAction) Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig {
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

