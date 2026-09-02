package interfacesawsdms


// A reference to a ReplicationSubnetGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationSubnetGroupReference := &ReplicationSubnetGroupReference{
//   	ReplicationSubnetGroupIdentifier: jsii.String("replicationSubnetGroupIdentifier"),
//   }
//
type ReplicationSubnetGroupReference struct {
	// The ReplicationSubnetGroupIdentifier of the ReplicationSubnetGroup resource.
	ReplicationSubnetGroupIdentifier *string `field:"required" json:"replicationSubnetGroupIdentifier" yaml:"replicationSubnetGroupIdentifier"`
}

