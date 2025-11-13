package interfacesawsssmincidents


// A reference to a ReplicationSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationSetReference := &ReplicationSetReference{
//   	ReplicationSetArn: jsii.String("replicationSetArn"),
//   }
//
type ReplicationSetReference struct {
	// The Arn of the ReplicationSet resource.
	ReplicationSetArn *string `field:"required" json:"replicationSetArn" yaml:"replicationSetArn"`
}

