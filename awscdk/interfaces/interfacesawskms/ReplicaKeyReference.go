package interfacesawskms


// A reference to a ReplicaKey resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicaKeyReference := &ReplicaKeyReference{
//   	KeyId: jsii.String("keyId"),
//   	ReplicaKeyArn: jsii.String("replicaKeyArn"),
//   }
//
type ReplicaKeyReference struct {
	// The KeyId of the ReplicaKey resource.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// The ARN of the ReplicaKey resource.
	ReplicaKeyArn *string `field:"required" json:"replicaKeyArn" yaml:"replicaKeyArn"`
}

