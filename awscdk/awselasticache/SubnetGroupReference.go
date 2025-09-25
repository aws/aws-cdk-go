package awselasticache


// A reference to a SubnetGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetGroupReference := &SubnetGroupReference{
//   	CacheSubnetGroupName: jsii.String("cacheSubnetGroupName"),
//   }
//
type SubnetGroupReference struct {
	// The CacheSubnetGroupName of the SubnetGroup resource.
	CacheSubnetGroupName *string `field:"required" json:"cacheSubnetGroupName" yaml:"cacheSubnetGroupName"`
}

