package interfacesawseks


// A reference to a Addon resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addonReference := &AddonReference{
//   	AddonArn: jsii.String("addonArn"),
//   	AddonName: jsii.String("addonName"),
//   	ClusterName: jsii.String("clusterName"),
//   }
//
type AddonReference struct {
	// The ARN of the Addon resource.
	AddonArn *string `field:"required" json:"addonArn" yaml:"addonArn"`
	// The AddonName of the Addon resource.
	AddonName *string `field:"required" json:"addonName" yaml:"addonName"`
	// The ClusterName of the Addon resource.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
}

