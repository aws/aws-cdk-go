package awsiot


// Describes an action to republish to another topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   republishActionProperty := &RepublishActionProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	Topic: jsii.String("topic"),
//
//   	// the properties below are optional
//   	Headers: &RepublishActionHeadersProperty{
//   		ContentType: jsii.String("contentType"),
//   		CorrelationData: jsii.String("correlationData"),
//   		MessageExpiry: jsii.String("messageExpiry"),
//   		PayloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   		ResponseTopic: jsii.String("responseTopic"),
//   		UserProperties: []interface{}{
//   			&UserPropertyProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	Qos: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-republishaction.html
//
type CfnTopicRule_RepublishActionProperty struct {
	// The ARN of the IAM role that grants access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-republishaction.html#cfn-iot-topicrule-republishaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the MQTT topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-republishaction.html#cfn-iot-topicrule-republishaction-topic
	//
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// MQTT Version 5.0 headers information. For more information, see [MQTT](https://docs.aws.amazon.com//iot/latest/developerguide/mqtt.html) in the IoT Core Developer Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-republishaction.html#cfn-iot-topicrule-republishaction-headers
	//
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// The Quality of Service (QoS) level to use when republishing messages.
	//
	// The default value is 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-republishaction.html#cfn-iot-topicrule-republishaction-qos
	//
	Qos *float64 `field:"optional" json:"qos" yaml:"qos"`
}

