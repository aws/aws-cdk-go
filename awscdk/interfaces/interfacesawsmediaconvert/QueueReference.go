package interfacesawsmediaconvert


// A reference to a Queue resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queueReference := &QueueReference{
//   	QueueArn: jsii.String("queueArn"),
//   	QueueName: jsii.String("queueName"),
//   }
//
type QueueReference struct {
	// The ARN of the Queue resource.
	QueueArn *string `field:"required" json:"queueArn" yaml:"queueArn"`
	// The Name of the Queue resource.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
}

