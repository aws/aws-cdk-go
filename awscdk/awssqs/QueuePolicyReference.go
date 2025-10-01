package awssqs


// A reference to a QueuePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queuePolicyReference := &QueuePolicyReference{
//   	QueuePolicyId: jsii.String("queuePolicyId"),
//   }
//
type QueuePolicyReference struct {
	// The Id of the QueuePolicy resource.
	QueuePolicyId *string `field:"required" json:"queuePolicyId" yaml:"queuePolicyId"`
}

