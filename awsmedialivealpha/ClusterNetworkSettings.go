package awsmedialivealpha


// Network settings for a MediaLive Cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   clusterNetworkSettings := &ClusterNetworkSettings{
//   	DefaultRoute: jsii.String("defaultRoute"),
//   	InterfaceMappings: []InterfaceMapping{
//   		&InterfaceMapping{
//   			LogicalInterfaceName: jsii.String("logicalInterfaceName"),
//   			NetworkId: jsii.String("networkId"),
//   		},
//   	},
//   }
//
// Experimental.
type ClusterNetworkSettings struct {
	// The default route for the cluster.
	// Default: - no default route.
	//
	// Experimental.
	DefaultRoute *string `field:"optional" json:"defaultRoute" yaml:"defaultRoute"`
	// The interface mappings for the cluster.
	// Default: - no interface mappings.
	//
	// Experimental.
	InterfaceMappings *[]*InterfaceMapping `field:"optional" json:"interfaceMappings" yaml:"interfaceMappings"`
}

