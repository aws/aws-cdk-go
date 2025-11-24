package mixinsawslambda


// Specific configuration settings for a self-managed Apache Kafka event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   selfManagedKafkaEventSourceConfigProperty := &SelfManagedKafkaEventSourceConfigProperty{
//   	ConsumerGroupId: jsii.String("consumerGroupId"),
//   	SchemaRegistryConfig: &SchemaRegistryConfigProperty{
//   		AccessConfigs: []interface{}{
//   			&SchemaRegistryAccessConfigProperty{
//   				Type: jsii.String("type"),
//   				Uri: jsii.String("uri"),
//   			},
//   		},
//   		EventRecordFormat: jsii.String("eventRecordFormat"),
//   		SchemaRegistryUri: jsii.String("schemaRegistryUri"),
//   		SchemaValidationConfigs: []interface{}{
//   			&SchemaValidationConfigProperty{
//   				Attribute: jsii.String("attribute"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-selfmanagedkafkaeventsourceconfig.html
//
type CfnEventSourceMappingPropsMixin_SelfManagedKafkaEventSourceConfigProperty struct {
	// The identifier for the Kafka consumer group to join.
	//
	// The consumer group ID must be unique among all your Kafka event sources. After creating a Kafka event source mapping with the consumer group ID specified, you cannot update this value. For more information, see [Customizable consumer group ID](https://docs.aws.amazon.com/lambda/latest/dg/with-kafka-process.html#services-smaa-topic-add) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-selfmanagedkafkaeventsourceconfig.html#cfn-lambda-eventsourcemapping-selfmanagedkafkaeventsourceconfig-consumergroupid
	//
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// Specific configuration settings for a Kafka schema registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-selfmanagedkafkaeventsourceconfig.html#cfn-lambda-eventsourcemapping-selfmanagedkafkaeventsourceconfig-schemaregistryconfig
	//
	SchemaRegistryConfig interface{} `field:"optional" json:"schemaRegistryConfig" yaml:"schemaRegistryConfig"`
}

