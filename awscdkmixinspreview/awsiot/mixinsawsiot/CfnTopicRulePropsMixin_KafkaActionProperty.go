package mixinsawsiot


// Send messages to an Amazon Managed Streaming for Apache Kafka (Amazon MSK) or self-managed Apache Kafka cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kafkaActionProperty := &KafkaActionProperty{
//   	ClientProperties: map[string]*string{
//   		"clientPropertiesKey": jsii.String("clientProperties"),
//   	},
//   	DestinationArn: jsii.String("destinationArn"),
//   	Headers: []interface{}{
//   		&KafkaActionHeaderProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Key: jsii.String("key"),
//   	Partition: jsii.String("partition"),
//   	Topic: jsii.String("topic"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaaction.html
//
type CfnTopicRulePropsMixin_KafkaActionProperty struct {
	// Properties of the Apache Kafka producer client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaaction.html#cfn-iot-topicrule-kafkaaction-clientproperties
	//
	ClientProperties interface{} `field:"optional" json:"clientProperties" yaml:"clientProperties"`
	// The ARN of Kafka action's VPC `TopicRuleDestination` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaaction.html#cfn-iot-topicrule-kafkaaction-destinationarn
	//
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// The list of Kafka headers that you specify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaaction.html#cfn-iot-topicrule-kafkaaction-headers
	//
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// The Kafka message key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaaction.html#cfn-iot-topicrule-kafkaaction-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The Kafka message partition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaaction.html#cfn-iot-topicrule-kafkaaction-partition
	//
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// The Kafka topic for messages to be sent to the Kafka broker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaaction.html#cfn-iot-topicrule-kafkaaction-topic
	//
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

