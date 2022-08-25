package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonManagedKafkaEventSourceConfigProperty := &amazonManagedKafkaEventSourceConfigProperty{
//   	consumerGroupId: jsii.String("consumerGroupId"),
//   }
//
type CfnEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty struct {
	// `CfnEventSourceMapping.AmazonManagedKafkaEventSourceConfigProperty.ConsumerGroupId`.
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
}

