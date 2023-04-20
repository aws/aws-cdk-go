package awsiotactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/awsiotactions/internal"
)

// The action to put the record from an MQTT message to republish another MQTT topic.
//
// Example:
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp, temperature FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewIotRepublishMqttAction(jsii.String("${topic()}/republish"), &IotRepublishMqttActionProps{
//   			QualityOfService: actions.MqttQualityOfService_AT_LEAST_ONCE,
//   		}),
//   	},
//   })
//
// Experimental.
type IotRepublishMqttAction interface {
	awsiot.IAction
	// Returns the topic rule action specification.
	// Experimental.
	Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for IotRepublishMqttAction
type jsiiProxy_IotRepublishMqttAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewIotRepublishMqttAction(topic *string, props *IotRepublishMqttActionProps) IotRepublishMqttAction {
	_init_.Initialize()

	if err := validateNewIotRepublishMqttActionParameters(topic, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotRepublishMqttAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.IotRepublishMqttAction",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIotRepublishMqttAction_Override(i IotRepublishMqttAction, topic *string, props *IotRepublishMqttActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.IotRepublishMqttAction",
		[]interface{}{topic, props},
		i,
	)
}

func (i *jsiiProxy_IotRepublishMqttAction) Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig {
	if err := i.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

