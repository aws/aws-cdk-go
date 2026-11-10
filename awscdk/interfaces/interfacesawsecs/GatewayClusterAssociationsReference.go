package interfacesawsecs


// A reference to a GatewayClusterAssociations resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayClusterAssociationsReference := &GatewayClusterAssociationsReference{
//   	Cluster: jsii.String("cluster"),
//   	Gateway: jsii.String("gateway"),
//   }
//
type GatewayClusterAssociationsReference struct {
	// The Cluster of the GatewayClusterAssociations resource.
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The Gateway of the GatewayClusterAssociations resource.
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
}

