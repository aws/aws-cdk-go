package awsappmesh


// Properties for a VirtualService provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var mesh Mesh
//
//   virtualServiceProviderConfig := &VirtualServiceProviderConfig{
//   	Mesh: mesh,
//
//   	// the properties below are optional
//   	VirtualNodeProvider: &VirtualNodeServiceProviderProperty{
//   		VirtualNodeName: jsii.String("virtualNodeName"),
//   	},
//   	VirtualRouterProvider: &VirtualRouterServiceProviderProperty{
//   		VirtualRouterName: jsii.String("virtualRouterName"),
//   	},
//   }
//
type VirtualServiceProviderConfig struct {
	// Mesh the Provider is using.
	// Default: - none.
	//
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
	// Virtual Node based provider.
	// Default: - none.
	//
	VirtualNodeProvider *CfnVirtualService_VirtualNodeServiceProviderProperty `field:"optional" json:"virtualNodeProvider" yaml:"virtualNodeProvider"`
	// Virtual Router based provider.
	// Default: - none.
	//
	VirtualRouterProvider *CfnVirtualService_VirtualRouterServiceProviderProperty `field:"optional" json:"virtualRouterProvider" yaml:"virtualRouterProvider"`
}

