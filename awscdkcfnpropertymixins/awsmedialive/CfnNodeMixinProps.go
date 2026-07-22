package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnNodePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnNodeMixinProps := &CfnNodeMixinProps{
//   	ClusterId: jsii.String("clusterId"),
//   	Name: jsii.String("name"),
//   	NodeInterfaceMappings: []interface{}{
//   		&NodeInterfaceMappingProperty{
//   			LogicalInterfaceName: jsii.String("logicalInterfaceName"),
//   			NetworkInterfaceMode: jsii.String("networkInterfaceMode"),
//   			PhysicalInterfaceName: jsii.String("physicalInterfaceName"),
//   		},
//   	},
//   	Role: jsii.String("role"),
//   	SdiSourceMappings: []interface{}{
//   		&SdiSourceMappingProperty{
//   			CardNumber: jsii.Number(123),
//   			ChannelNumber: jsii.Number(123),
//   			SdiSource: jsii.String("sdiSource"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-node.html
//
type CfnNodeMixinProps struct {
	// The ID of the Cluster that the Node belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-node.html#cfn-medialive-node-clusterid
	//
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// The user-specified name of the Node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-node.html#cfn-medialive-node-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of interface mappings for the Node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-node.html#cfn-medialive-node-nodeinterfacemappings
	//
	NodeInterfaceMappings interface{} `field:"optional" json:"nodeInterfaceMappings" yaml:"nodeInterfaceMappings"`
	// The role of the Node in the Cluster.
	//
	// ACTIVE means the Node is available for encoding. BACKUP means the Node is a redundant Node and might get used if an ACTIVE Node fails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-node.html#cfn-medialive-node-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
	// An array of SDI source mappings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-node.html#cfn-medialive-node-sdisourcemappings
	//
	SdiSourceMappings interface{} `field:"optional" json:"sdiSourceMappings" yaml:"sdiSourceMappings"`
	// A collection of key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-node.html#cfn-medialive-node-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

