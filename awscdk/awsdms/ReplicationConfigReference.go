package awsdms


// A reference to a ReplicationConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationConfigReference := &ReplicationConfigReference{
//   	ReplicationConfigArn: jsii.String("replicationConfigArn"),
//   }
//
type ReplicationConfigReference struct {
	// The ReplicationConfigArn of the ReplicationConfig resource.
	ReplicationConfigArn *string `field:"required" json:"replicationConfigArn" yaml:"replicationConfigArn"`
}

