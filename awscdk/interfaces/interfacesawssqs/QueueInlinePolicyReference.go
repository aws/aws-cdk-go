package interfacesawssqs


// A reference to a QueueInlinePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queueInlinePolicyReference := &QueueInlinePolicyReference{
//   	Queue: jsii.String("queue"),
//   }
//
type QueueInlinePolicyReference struct {
	// The Queue of the QueueInlinePolicy resource.
	Queue *string `field:"required" json:"queue" yaml:"queue"`
}

