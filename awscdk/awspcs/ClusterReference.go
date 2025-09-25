package awspcs


// A reference to a Cluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterReference := &ClusterReference{
//   	ClusterArn: jsii.String("clusterArn"),
//   }
//
type ClusterReference struct {
	// The Arn of the Cluster resource.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
}

