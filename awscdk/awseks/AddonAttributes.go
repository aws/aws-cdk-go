package awseks


// Represents the attributes of an addon for an Amazon EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addonAttributes := &AddonAttributes{
//   	AddonName: jsii.String("addonName"),
//   	ClusterName: jsii.String("clusterName"),
//   }
//
type AddonAttributes struct {
	// The name of the addon.
	AddonName *string `field:"required" json:"addonName" yaml:"addonName"`
	// The name of the Amazon EKS cluster the addon is associated with.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
}

