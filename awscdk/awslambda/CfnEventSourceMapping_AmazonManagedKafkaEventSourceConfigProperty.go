package awslambda


// Specific configuration settings for an MSK event source.
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
	// The identifier for the Kafka Consumer Group to join.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-amazonmanagedkafkaeventsourceconfig.html#cfn-lambda-eventsourcemapping-amazonmanagedkafkaeventsourceconfig-consumergroupid
	//
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
}

