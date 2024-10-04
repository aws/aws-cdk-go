package awslambda


// Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonManagedKafkaEventSourceConfigProperty := &AmazonManagedKafkaEventSourceConfigProperty{
//   	ConsumerGroupId: jsii.String("consumerGroupId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-amazonmanagedkafkaeventsourceconfig.html
//
type CfnEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty struct {
	// The identifier for the Kafka consumer group to join.
	//
	// The consumer group ID must be unique among all your Kafka event sources. After creating a Kafka event source mapping with the consumer group ID specified, you cannot update this value. For more information, see [Customizable consumer group ID](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-consumer-group-id) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-amazonmanagedkafkaeventsourceconfig.html#cfn-lambda-eventsourcemapping-amazonmanagedkafkaeventsourceconfig-consumergroupid
	//
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
}

