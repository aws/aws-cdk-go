package awscdkeksv2alpha


// Represents the attributes of an addon for an Amazon EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeksv2alpha"
//
//   addonAttributes := &AddonAttributes{
//   	AddonName: jsii.String("addonName"),
//   	ClusterName: jsii.String("clusterName"),
//   }
//
// Deprecated.
type AddonAttributes struct {
	// The name of the addon.
	// Deprecated.
	AddonName *string `field:"required" json:"addonName" yaml:"addonName"`
	// The name of the Amazon EKS cluster the addon is associated with.
	// Deprecated.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
}

