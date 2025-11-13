package interfacesawsdms


// A reference to a ReplicationInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationInstanceReference := &ReplicationInstanceReference{
//   	ReplicationInstanceId: jsii.String("replicationInstanceId"),
//   }
//
type ReplicationInstanceReference struct {
	// The Id of the ReplicationInstance resource.
	ReplicationInstanceId *string `field:"required" json:"replicationInstanceId" yaml:"replicationInstanceId"`
}

