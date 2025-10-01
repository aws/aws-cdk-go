package awsmedialive


// On premises settings which will have the interface network mappings and default Output logical interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterNetworkSettingsProperty := &ClusterNetworkSettingsProperty{
//   	DefaultRoute: jsii.String("defaultRoute"),
//   	InterfaceMappings: []interface{}{
//   		&InterfaceMappingProperty{
//   			LogicalInterfaceName: jsii.String("logicalInterfaceName"),
//   			NetworkId: jsii.String("networkId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-cluster-clusternetworksettings.html
//
type CfnCluster_ClusterNetworkSettingsProperty struct {
	// Default value if the customer does not define it in channel Output API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-cluster-clusternetworksettings.html#cfn-medialive-cluster-clusternetworksettings-defaultroute
	//
	DefaultRoute *string `field:"optional" json:"defaultRoute" yaml:"defaultRoute"`
	// Network mappings for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-cluster-clusternetworksettings.html#cfn-medialive-cluster-clusternetworksettings-interfacemappings
	//
	InterfaceMappings interface{} `field:"optional" json:"interfaceMappings" yaml:"interfaceMappings"`
}

