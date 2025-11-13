package interfacesawslocation


// A reference to a TrackerConsumer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trackerConsumerReference := &TrackerConsumerReference{
//   	ConsumerArn: jsii.String("consumerArn"),
//   	TrackerName: jsii.String("trackerName"),
//   }
//
type TrackerConsumerReference struct {
	// The ConsumerArn of the TrackerConsumer resource.
	ConsumerArn *string `field:"required" json:"consumerArn" yaml:"consumerArn"`
	// The TrackerName of the TrackerConsumer resource.
	TrackerName *string `field:"required" json:"trackerName" yaml:"trackerName"`
}

