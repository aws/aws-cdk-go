package awseks


// A reference to a Cluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterReference := &ClusterReference{
//   	ClusterArn: jsii.String("clusterArn"),
//   	ClusterName: jsii.String("clusterName"),
//   }
//
type ClusterReference struct {
	// The ARN of the Cluster resource.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// The Name of the Cluster resource.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
}

