package interfacesawsmemorydb


// A reference to a MultiRegionCluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiRegionClusterReference := &MultiRegionClusterReference{
//   	MultiRegionClusterArn: jsii.String("multiRegionClusterArn"),
//   	MultiRegionClusterName: jsii.String("multiRegionClusterName"),
//   }
//
type MultiRegionClusterReference struct {
	// The ARN of the MultiRegionCluster resource.
	MultiRegionClusterArn *string `field:"required" json:"multiRegionClusterArn" yaml:"multiRegionClusterArn"`
	// The MultiRegionClusterName of the MultiRegionCluster resource.
	MultiRegionClusterName *string `field:"required" json:"multiRegionClusterName" yaml:"multiRegionClusterName"`
}

