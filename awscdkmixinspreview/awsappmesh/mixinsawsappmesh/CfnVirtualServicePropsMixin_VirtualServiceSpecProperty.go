package mixinsawsappmesh


// An object that represents the specification of a virtual service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualServiceSpecProperty := &VirtualServiceSpecProperty{
//   	Provider: &VirtualServiceProviderProperty{
//   		VirtualNode: &VirtualNodeServiceProviderProperty{
//   			VirtualNodeName: jsii.String("virtualNodeName"),
//   		},
//   		VirtualRouter: &VirtualRouterServiceProviderProperty{
//   			VirtualRouterName: jsii.String("virtualRouterName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualservice-virtualservicespec.html
//
type CfnVirtualServicePropsMixin_VirtualServiceSpecProperty struct {
	// The App Mesh object that is acting as the provider for a virtual service.
	//
	// You can specify a single virtual node or virtual router.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualservice-virtualservicespec.html#cfn-appmesh-virtualservice-virtualservicespec-provider
	//
	Provider interface{} `field:"optional" json:"provider" yaml:"provider"`
}

