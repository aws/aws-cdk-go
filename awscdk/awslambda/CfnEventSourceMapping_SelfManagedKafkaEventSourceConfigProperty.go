package awslambda


// Specific configuration settings for a Self-Managed Apache Kafka event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selfManagedKafkaEventSourceConfigProperty := &SelfManagedKafkaEventSourceConfigProperty{
//   	ConsumerGroupId: jsii.String("consumerGroupId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-selfmanagedkafkaeventsourceconfig.html
//
type CfnEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty struct {
	// The identifier for the Kafka Consumer Group to join.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-selfmanagedkafkaeventsourceconfig.html#cfn-lambda-eventsourcemapping-selfmanagedkafkaeventsourceconfig-consumergroupid
	//
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
}

