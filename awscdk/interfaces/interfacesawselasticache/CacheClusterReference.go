package interfacesawselasticache


// A reference to a CacheCluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cacheClusterReference := &CacheClusterReference{
//   	CacheClusterId: jsii.String("cacheClusterId"),
//   }
//
type CacheClusterReference struct {
	// The Id of the CacheCluster resource.
	CacheClusterId *string `field:"required" json:"cacheClusterId" yaml:"cacheClusterId"`
}

