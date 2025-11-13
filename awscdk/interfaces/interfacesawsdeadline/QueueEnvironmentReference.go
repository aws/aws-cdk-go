package interfacesawsdeadline


// A reference to a QueueEnvironment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queueEnvironmentReference := &QueueEnvironmentReference{
//   	FarmId: jsii.String("farmId"),
//   	QueueEnvironmentId: jsii.String("queueEnvironmentId"),
//   	QueueId: jsii.String("queueId"),
//   }
//
type QueueEnvironmentReference struct {
	// The FarmId of the QueueEnvironment resource.
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// The QueueEnvironmentId of the QueueEnvironment resource.
	QueueEnvironmentId *string `field:"required" json:"queueEnvironmentId" yaml:"queueEnvironmentId"`
	// The QueueId of the QueueEnvironment resource.
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
}

