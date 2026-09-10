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
//   }
//
type ReplicationTaskReference struct {
	// The ReplicationTaskArn of the ReplicationTask resource.
	ReplicationTaskArn *string `field:"required" json:"replicationTaskArn" yaml:"replicationTaskArn"`
}

