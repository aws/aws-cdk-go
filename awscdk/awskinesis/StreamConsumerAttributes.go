package awskinesis


// A reference to a StreamConsumer, which can be imported using `StreamConsumer.fromStreamConsumerAttributes`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamConsumerAttributes := &StreamConsumerAttributes{
//   	StreamConsumerArn: jsii.String("streamConsumerArn"),
//   }
//
type StreamConsumerAttributes struct {
	// The Amazon Resource Name (ARN) of the stream consumer.
	StreamConsumerArn *string `field:"required" json:"streamConsumerArn" yaml:"streamConsumerArn"`
}

