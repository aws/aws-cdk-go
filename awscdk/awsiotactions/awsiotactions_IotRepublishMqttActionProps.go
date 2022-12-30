package awsiotactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Configuration properties of an action to republish MQTT messages.
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
type IotRepublishMqttActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The Quality of Service (QoS) level to use when republishing messages.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/mqtt.html#mqtt-qos
	//
	// Experimental.
	QualityOfService MqttQualityOfService `field:"optional" json:"qualityOfService" yaml:"qualityOfService"`
}

