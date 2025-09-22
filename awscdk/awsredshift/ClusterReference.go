package awsredshift


// A reference to a Cluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterReference := &ClusterReference{
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   }
//
type ClusterReference struct {
	// The ClusterIdentifier of the Cluster resource.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
}

