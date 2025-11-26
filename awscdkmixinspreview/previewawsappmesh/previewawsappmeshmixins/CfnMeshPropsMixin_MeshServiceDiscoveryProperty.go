package previewawsappmeshmixins


// An object that represents the service discovery information for a service mesh.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   meshServiceDiscoveryProperty := &MeshServiceDiscoveryProperty{
//   	IpPreference: jsii.String("ipPreference"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-mesh-meshservicediscovery.html
//
type CfnMeshPropsMixin_MeshServiceDiscoveryProperty struct {
	// The IP version to use to control traffic within the mesh.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-mesh-meshservicediscovery.html#cfn-appmesh-mesh-meshservicediscovery-ippreference
	//
	IpPreference *string `field:"optional" json:"ipPreference" yaml:"ipPreference"`
}

