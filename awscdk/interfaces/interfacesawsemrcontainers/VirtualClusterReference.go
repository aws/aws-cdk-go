package interfacesawsemrcontainers


// A reference to a VirtualCluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualClusterReference := &VirtualClusterReference{
//   	VirtualClusterArn: jsii.String("virtualClusterArn"),
//   	VirtualClusterId: jsii.String("virtualClusterId"),
//   }
//
type VirtualClusterReference struct {
	// The ARN of the VirtualCluster resource.
	VirtualClusterArn *string `field:"required" json:"virtualClusterArn" yaml:"virtualClusterArn"`
	// The Id of the VirtualCluster resource.
	VirtualClusterId *string `field:"required" json:"virtualClusterId" yaml:"virtualClusterId"`
}

