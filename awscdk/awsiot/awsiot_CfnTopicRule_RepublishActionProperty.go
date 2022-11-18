package awsiot


// Describes an action to republish to another topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   republishActionProperty := &republishActionProperty{
//   	roleArn: jsii.String("roleArn"),
//   	topic: jsii.String("topic"),
//
//   	// the properties below are optional
//   	headers: &republishActionHeadersProperty{
//   		contentType: jsii.String("contentType"),
//   		correlationData: jsii.String("correlationData"),
//   		messageExpiry: jsii.String("messageExpiry"),
//   		payloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   		responseTopic: jsii.String("responseTopic"),
//   		userProperties: []interface{}{
//   			&userPropertyProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	qos: jsii.Number(123),
//   }
//
type CfnTopicRule_RepublishActionProperty struct {
	// The ARN of the IAM role that grants access.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the MQTT topic.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// `CfnTopicRule.RepublishActionProperty.Headers`.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// The Quality of Service (QoS) level to use when republishing messages.
	//
	// The default value is 0.
	Qos *float64 `field:"optional" json:"qos" yaml:"qos"`
}

