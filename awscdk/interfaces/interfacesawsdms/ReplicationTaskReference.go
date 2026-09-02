package interfacesawsdms


// A reference to a ReplicationTask resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationTaskReference := &ReplicationTaskReference{
//   	ReplicationTaskArn: jsii.String("replicationTaskArn"),
//   	ReplicationTaskId: jsii.String("replicationTaskId"),
//   }
//
type ReplicationTaskReference struct {
	// The ARN of the ReplicationTask resource.
	ReplicationTaskArn *string `field:"required" json:"replicationTaskArn" yaml:"replicationTaskArn"`
	// The Id of the ReplicationTask resource.
	ReplicationTaskId *string `field:"required" json:"replicationTaskId" yaml:"replicationTaskId"`
}

