package interfacesawsdms


// A reference to a ReplicationInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationInstanceReference := &ReplicationInstanceReference{
//   	ReplicationInstanceArn: jsii.String("replicationInstanceArn"),
//   }
//
type ReplicationInstanceReference struct {
	// The ReplicationInstanceArn of the ReplicationInstance resource.
	ReplicationInstanceArn *string `field:"required" json:"replicationInstanceArn" yaml:"replicationInstanceArn"`
}

