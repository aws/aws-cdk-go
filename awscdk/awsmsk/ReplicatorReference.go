package awsmsk


// A reference to a Replicator resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicatorReference := &ReplicatorReference{
//   	ReplicatorArn: jsii.String("replicatorArn"),
//   }
//
type ReplicatorReference struct {
	// The ReplicatorArn of the Replicator resource.
	ReplicatorArn *string `field:"required" json:"replicatorArn" yaml:"replicatorArn"`
}

