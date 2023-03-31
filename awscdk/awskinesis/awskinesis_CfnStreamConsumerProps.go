package awskinesis


// Properties for defining a `CfnStreamConsumer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStreamConsumerProps := &cfnStreamConsumerProps{
//   	consumerName: jsii.String("consumerName"),
//   	streamArn: jsii.String("streamArn"),
//   }
//
type CfnStreamConsumerProps struct {
	// The name of the consumer is something you choose when you register the consumer.
	ConsumerName *string `field:"required" json:"consumerName" yaml:"consumerName"`
	// The ARN of the stream with which you registered the consumer.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

