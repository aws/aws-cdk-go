package awscdkiotactionsalpha


// MQTT Quality of Service (QoS) indicates the level of assurance for delivery of an MQTT Message.
//
// Example:
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp, temperature FROM 'device/+/data'")),
//   	Actions: []IAction{
//   		actions.NewIotRepublishMqttAction(jsii.String("${topic()}/republish"), &IotRepublishMqttActionProps{
//   			QualityOfService: actions.MqttQualityOfService_AT_LEAST_ONCE,
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/iot/latest/developerguide/mqtt.html#mqtt-qos
//
// Experimental.
type MqttQualityOfService string

const (
	// QoS level 0.
	//
	// Sent zero or more times.
	// This level should be used for messages that are sent over reliable communication links or that can be missed without a problem.
	// Experimental.
	MqttQualityOfService_ZERO_OR_MORE_TIMES MqttQualityOfService = "ZERO_OR_MORE_TIMES"
	// QoS level 1.
	//
	// Sent at least one time, and then repeatedly until a PUBACK response is received.
	// The message is not considered complete until the sender receives a PUBACK response to indicate successful delivery.
	// Experimental.
	MqttQualityOfService_AT_LEAST_ONCE MqttQualityOfService = "AT_LEAST_ONCE"
)

