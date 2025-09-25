package awsdms


// A reference to a ReplicationTask resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationTaskReference := &ReplicationTaskReference{
//   	ReplicationTaskId: jsii.String("replicationTaskId"),
//   }
//
type ReplicationTaskReference struct {
	// The Id of the ReplicationTask resource.
	ReplicationTaskId *string `field:"required" json:"replicationTaskId" yaml:"replicationTaskId"`
}

