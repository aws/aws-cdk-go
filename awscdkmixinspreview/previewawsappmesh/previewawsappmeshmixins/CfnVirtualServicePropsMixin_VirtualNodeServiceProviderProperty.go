package previewawsappmeshmixins


// An object that represents a virtual node service provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualNodeServiceProviderProperty := &VirtualNodeServiceProviderProperty{
//   	VirtualNodeName: jsii.String("virtualNodeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualservice-virtualnodeserviceprovider.html
//
type CfnVirtualServicePropsMixin_VirtualNodeServiceProviderProperty struct {
	// The name of the virtual node that is acting as a service provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualservice-virtualnodeserviceprovider.html#cfn-appmesh-virtualservice-virtualnodeserviceprovider-virtualnodename
	//
	VirtualNodeName *string `field:"optional" json:"virtualNodeName" yaml:"virtualNodeName"`
}

