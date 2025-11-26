package previewawsappmeshmixins


// An object that represents a virtual node service provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualRouterServiceProviderProperty := &VirtualRouterServiceProviderProperty{
//   	VirtualRouterName: jsii.String("virtualRouterName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualservice-virtualrouterserviceprovider.html
//
type CfnVirtualServicePropsMixin_VirtualRouterServiceProviderProperty struct {
	// The name of the virtual router that is acting as a service provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualservice-virtualrouterserviceprovider.html#cfn-appmesh-virtualservice-virtualrouterserviceprovider-virtualroutername
	//
	VirtualRouterName *string `field:"optional" json:"virtualRouterName" yaml:"virtualRouterName"`
}

