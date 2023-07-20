package awsiot


// Send messages to an Amazon Managed Streaming for Apache Kafka (Amazon MSK) or self-managed Apache Kafka cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaActionProperty := &KafkaActionProperty{
//   	ClientProperties: map[string]*string{
//   		"clientPropertiesKey": jsii.String("clientProperties"),
//   	},
//   	DestinationArn: jsii.String("destinationArn"),
//   	Topic: jsii.String("topic"),
//
//   	// the properties below are optional
//   	Key: jsii.String("key"),
//   	Partition: jsii.String("partition"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaaction.html
//
type CfnTopicRule_KafkaActionProperty struct {
	// Properties of the Apache Kafka producer client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaaction.html#cfn-iot-topicrule-kafkaaction-clientproperties
	//
	ClientProperties interface{} `field:"required" json:"clientProperties" yaml:"clientProperties"`
	// The ARN of Kafka action's VPC `TopicRuleDestination` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaaction.html#cfn-iot-topicrule-kafkaaction-destinationarn
	//
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// The Kafka topic for messages to be sent to the Kafka broker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaaction.html#cfn-iot-topicrule-kafkaaction-topic
	//
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The Kafka message key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaaction.html#cfn-iot-topicrule-kafkaaction-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The Kafka message partition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaaction.html#cfn-iot-topicrule-kafkaaction-partition
	//
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
}

