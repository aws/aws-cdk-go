package interfacesawselasticache


// A reference to a CacheCluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cacheClusterReference := &CacheClusterReference{
//   	ClusterName: jsii.String("clusterName"),
//   }
//
type CacheClusterReference struct {
	// The ClusterName of the CacheCluster resource.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
}

