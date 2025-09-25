package awsecs


// A reference to a ClusterCapacityProviderAssociations resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterCapacityProviderAssociationsReference := &ClusterCapacityProviderAssociationsReference{
//   	Cluster: jsii.String("cluster"),
//   }
//
type ClusterCapacityProviderAssociationsReference struct {
	// The Cluster of the ClusterCapacityProviderAssociations resource.
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
}

