package interfacesawsmsk


// A reference to a ClusterPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterPolicyReference := &ClusterPolicyReference{
//   	ClusterArn: jsii.String("clusterArn"),
//   }
//
type ClusterPolicyReference struct {
	// The ClusterArn of the ClusterPolicy resource.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
}

