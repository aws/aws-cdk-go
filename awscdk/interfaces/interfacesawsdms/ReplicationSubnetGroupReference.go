package interfacesawsdms


// A reference to a ReplicationSubnetGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationSubnetGroupReference := &ReplicationSubnetGroupReference{
//   	ReplicationSubnetGroupId: jsii.String("replicationSubnetGroupId"),
//   }
//
type ReplicationSubnetGroupReference struct {
	// The Id of the ReplicationSubnetGroup resource.
	ReplicationSubnetGroupId *string `field:"required" json:"replicationSubnetGroupId" yaml:"replicationSubnetGroupId"`
}

