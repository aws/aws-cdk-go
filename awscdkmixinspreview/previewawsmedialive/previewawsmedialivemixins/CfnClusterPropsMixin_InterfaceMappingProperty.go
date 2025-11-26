package previewawsmedialivemixins


// Network mappings for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   interfaceMappingProperty := &InterfaceMappingProperty{
//   	LogicalInterfaceName: jsii.String("logicalInterfaceName"),
//   	NetworkId: jsii.String("networkId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-cluster-interfacemapping.html
//
type CfnClusterPropsMixin_InterfaceMappingProperty struct {
	// logical interface name, unique in the list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-cluster-interfacemapping.html#cfn-medialive-cluster-interfacemapping-logicalinterfacename
	//
	LogicalInterfaceName *string `field:"optional" json:"logicalInterfaceName" yaml:"logicalInterfaceName"`
	// Network Id to be associated with the logical interface name, can be duplicated in list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-cluster-interfacemapping.html#cfn-medialive-cluster-interfacemapping-networkid
	//
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
}

