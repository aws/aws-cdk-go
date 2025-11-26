package previewawsappmeshmixins


// An object that represents the provider for a virtual service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualservice-virtualserviceprovider.html
//
type CfnVirtualServicePropsMixin_VirtualServiceProviderProperty struct {
	// The virtual node associated with a virtual service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualservice-virtualserviceprovider.html#cfn-appmesh-virtualservice-virtualserviceprovider-virtualnode
	//
	VirtualNode interface{} `field:"optional" json:"virtualNode" yaml:"virtualNode"`
	// The virtual router associated with a virtual service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualservice-virtualserviceprovider.html#cfn-appmesh-virtualservice-virtualserviceprovider-virtualrouter
	//
	VirtualRouter interface{} `field:"optional" json:"virtualRouter" yaml:"virtualRouter"`
}

