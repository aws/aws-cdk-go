package interfacesawskinesis


// A reference to a StreamConsumer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamConsumerReference := &StreamConsumerReference{
//   	ConsumerArn: jsii.String("consumerArn"),
//   }
//
type StreamConsumerReference struct {
	// The ConsumerARN of the StreamConsumer resource.
	ConsumerArn *string `field:"required" json:"consumerArn" yaml:"consumerArn"`
}

