package awsmsk


// A reference to a BatchScramSecret resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchScramSecretReference := &BatchScramSecretReference{
//   	ClusterArn: jsii.String("clusterArn"),
//   }
//
type BatchScramSecretReference struct {
	// The ClusterArn of the BatchScramSecret resource.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
}

