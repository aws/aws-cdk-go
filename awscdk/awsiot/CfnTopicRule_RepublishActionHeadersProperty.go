package awsiot


// Specifies MQTT Version 5.0 headers information. For more information, see [MQTT](https://docs.aws.amazon.com//iot/latest/developerguide/mqtt.html) in the IoT Core Developer Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   republishActionHeadersProperty := &RepublishActionHeadersProperty{
//   	ContentType: jsii.String("contentType"),
//   	CorrelationData: jsii.String("correlationData"),
//   	MessageExpiry: jsii.String("messageExpiry"),
//   	PayloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   	ResponseTopic: jsii.String("responseTopic"),
//   	UserProperties: []interface{}{
//   		&UserPropertyProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTopicRule_RepublishActionHeadersProperty struct {
	// A UTF-8 encoded string that describes the content of the publishing message.
	//
	// For more information, see [Content Type](https://docs.aws.amazon.com/https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901118) in the MQTT Version 5.0 specification.
	//
	// Supports [substitution templates](https://docs.aws.amazon.com//iot/latest/developerguide/iot-substitution-templates.html) .
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The base64-encoded binary data used by the sender of the request message to identify which request the response message is for.
	//
	// For more information, see [Correlation Data](https://docs.aws.amazon.com/https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901115) in the MQTT Version 5.0 specification.
	//
	// Supports [substitution templates](https://docs.aws.amazon.com//iot/latest/developerguide/iot-substitution-templates.html) .
	//
	// > This binary data must be base64-encoded.
	CorrelationData *string `field:"optional" json:"correlationData" yaml:"correlationData"`
	// A user-defined integer value that represents the message expiry interval at the broker.
	//
	// If the messages haven't been sent to the subscribers within that interval, the message expires and is removed. The value of `messageExpiry` represents the number of seconds before it expires. For more information about the limits of `messageExpiry` , see [Message broker and protocol limits and quotas](https://docs.aws.amazon.com//general/latest/gr/iot-core.html#limits_iot) in the IoT Core Reference Guide.
	//
	// Supports [substitution templates](https://docs.aws.amazon.com//iot/latest/developerguide/iot-substitution-templates.html) .
	MessageExpiry *string `field:"optional" json:"messageExpiry" yaml:"messageExpiry"`
	// An `Enum` string value that indicates whether the payload is formatted as UTF-8.
	//
	// Valid values are `UNSPECIFIED_BYTES` and `UTF8_DATA` .
	//
	// For more information, see [Payload Format Indicator](https://docs.aws.amazon.com/https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901111) from the MQTT Version 5.0 specification.
	//
	// Supports [substitution templates](https://docs.aws.amazon.com//iot/latest/developerguide/iot-substitution-templates.html) .
	PayloadFormatIndicator *string `field:"optional" json:"payloadFormatIndicator" yaml:"payloadFormatIndicator"`
	// A UTF-8 encoded string that's used as the topic name for a response message.
	//
	// The response topic is used to describe the topic to which the receiver should publish as part of the request-response flow. The topic must not contain wildcard characters.
	//
	// For more information, see [Response Topic](https://docs.aws.amazon.com/https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901114) in the MQTT Version 5.0 specification.
	//
	// Supports [substitution templates](https://docs.aws.amazon.com//iot/latest/developerguide/iot-substitution-templates.html) .
	ResponseTopic *string `field:"optional" json:"responseTopic" yaml:"responseTopic"`
	// An array of key-value pairs that you define in the MQTT5 header.
	UserProperties interface{} `field:"optional" json:"userProperties" yaml:"userProperties"`
}

