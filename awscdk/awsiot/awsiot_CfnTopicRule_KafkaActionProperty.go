package awsiot


// Send messages to an Amazon Managed Streaming for Apache Kafka (Amazon MSK) or self-managed Apache Kafka cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaActionProperty := &kafkaActionProperty{
//   	clientProperties: map[string]*string{
//   		"clientPropertiesKey": jsii.String("clientProperties"),
//   	},
//   	destinationArn: jsii.String("destinationArn"),
//   	topic: jsii.String("topic"),
//
//   	// the properties below are optional
//   	key: jsii.String("key"),
//   	partition: jsii.String("partition"),
//   }
//
type CfnTopicRule_KafkaActionProperty struct {
	// Properties of the Apache Kafka producer client.
	ClientProperties interface{} `field:"required" json:"clientProperties" yaml:"clientProperties"`
	// The ARN of Kafka action's VPC `TopicRuleDestination` .
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// The Kafka topic for messages to be sent to the Kafka broker.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The Kafka message key.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The Kafka message partition.
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
}

