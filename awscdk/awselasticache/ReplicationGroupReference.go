package awselasticache


// A reference to a ReplicationGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationGroupReference := &ReplicationGroupReference{
//   	ReplicationGroupId: jsii.String("replicationGroupId"),
//   }
//
type ReplicationGroupReference struct {
	// The ReplicationGroupId of the ReplicationGroup resource.
	ReplicationGroupId *string `field:"required" json:"replicationGroupId" yaml:"replicationGroupId"`
}

