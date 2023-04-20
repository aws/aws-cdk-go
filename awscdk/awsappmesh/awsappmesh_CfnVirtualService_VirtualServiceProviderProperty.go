package awsappmesh


// An object that represents the provider for a virtual service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualServiceProviderProperty := &VirtualServiceProviderProperty{
//   	VirtualNode: &VirtualNodeServiceProviderProperty{
//   		VirtualNodeName: jsii.String("virtualNodeName"),
//   	},
//   	VirtualRouter: &VirtualRouterServiceProviderProperty{
//   		VirtualRouterName: jsii.String("virtualRouterName"),
//   	},
//   }
//
type CfnVirtualService_VirtualServiceProviderProperty struct {
	// The virtual node associated with a virtual service.
	VirtualNode interface{} `field:"optional" json:"virtualNode" yaml:"virtualNode"`
	// The virtual router associated with a virtual service.
	VirtualRouter interface{} `field:"optional" json:"virtualRouter" yaml:"virtualRouter"`
}

