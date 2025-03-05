package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &CfnClusterProps{
//   	ClusterType: jsii.String("clusterType"),
//   	InstanceRoleArn: jsii.String("instanceRoleArn"),
//   	Name: jsii.String("name"),
//   	NetworkSettings: &ClusterNetworkSettingsProperty{
//   		DefaultRoute: jsii.String("defaultRoute"),
//   		InterfaceMappings: []interface{}{
//   			&InterfaceMappingProperty{
//   				LogicalInterfaceName: jsii.String("logicalInterfaceName"),
//   				NetworkId: jsii.String("networkId"),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cluster.html
//
type CfnClusterProps struct {
	// The hardware type for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cluster.html#cfn-medialive-cluster-clustertype
	//
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	// The IAM role your nodes will use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cluster.html#cfn-medialive-cluster-instancerolearn
	//
	InstanceRoleArn *string `field:"optional" json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// The user-specified name of the Cluster to be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cluster.html#cfn-medialive-cluster-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// On premises settings which will have the interface network mappings and default Output logical interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cluster.html#cfn-medialive-cluster-networksettings
	//
	NetworkSettings interface{} `field:"optional" json:"networkSettings" yaml:"networkSettings"`
	// A collection of key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cluster.html#cfn-medialive-cluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

