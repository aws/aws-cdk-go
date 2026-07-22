package awsmedialive


// Node interface mapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   nodeInterfaceMappingProperty := &NodeInterfaceMappingProperty{
//   	LogicalInterfaceName: jsii.String("logicalInterfaceName"),
//   	NetworkInterfaceMode: jsii.String("networkInterfaceMode"),
//   	PhysicalInterfaceName: jsii.String("physicalInterfaceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-node-nodeinterfacemapping.html
//
type CfnNodePropsMixin_NodeInterfaceMappingProperty struct {
	// The logical name for this interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-node-nodeinterfacemapping.html#cfn-medialive-node-nodeinterfacemapping-logicalinterfacename
	//
	LogicalInterfaceName *string `field:"optional" json:"logicalInterfaceName" yaml:"logicalInterfaceName"`
	// The network interface mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-node-nodeinterfacemapping.html#cfn-medialive-node-nodeinterfacemapping-networkinterfacemode
	//
	NetworkInterfaceMode *string `field:"optional" json:"networkInterfaceMode" yaml:"networkInterfaceMode"`
	// The physical interface name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-node-nodeinterfacemapping.html#cfn-medialive-node-nodeinterfacemapping-physicalinterfacename
	//
	PhysicalInterfaceName *string `field:"optional" json:"physicalInterfaceName" yaml:"physicalInterfaceName"`
}

