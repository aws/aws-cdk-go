package awsiotevents


// Information required to publish the MQTT message through the AWS IoT message broker.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotTopicPublishProperty := &IotTopicPublishProperty{
//   	MqttTopic: jsii.String("mqttTopic"),
//
//   	// the properties below are optional
//   	Payload: &PayloadProperty{
//   		ContentExpression: jsii.String("contentExpression"),
//   		Type: jsii.String("type"),
//   	},
//   }
//
type CfnDetectorModel_IotTopicPublishProperty struct {
	// The MQTT topic of the message.
	//
	// You can use a string expression that includes variables ( `$variable.<variable-name>` ) and input values ( `$input.<input-name>.<path-to-datum>` ) as the topic string.
	MqttTopic *string `field:"required" json:"mqttTopic" yaml:"mqttTopic"`
	// You can configure the action payload when you publish a message to an AWS IoT Core topic.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

