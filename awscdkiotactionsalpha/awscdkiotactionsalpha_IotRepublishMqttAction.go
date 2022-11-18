// Receipt rule actions for AWS IoT
package awscdkiotactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2"
)

// The action to put the record from an MQTT message to republish another MQTT topic.
//
// Example:
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp, temperature FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewIotRepublishMqttAction(jsii.String("${topic()}/republish"), &iotRepublishMqttActionProps{
//   			qualityOfService: actions.mqttQualityOfService_AT_LEAST_ONCE,
//   		}),
//   	},
//   })
//
// Experimental.
type IotRepublishMqttAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for IotRepublishMqttAction
type jsiiProxy_IotRepublishMqttAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewIotRepublishMqttAction(topic *string, props *IotRepublishMqttActionProps) IotRepublishMqttAction {
	_init_.Initialize()

	if err := validateNewIotRepublishMqttActionParameters(topic, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotRepublishMqttAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.IotRepublishMqttAction",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIotRepublishMqttAction_Override(i IotRepublishMqttAction, topic *string, props *IotRepublishMqttActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.IotRepublishMqttAction",
		[]interface{}{topic, props},
		i,
	)
}

