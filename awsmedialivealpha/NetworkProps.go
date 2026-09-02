package awsmedialivealpha


// Properties for creating a MediaLive Network.
//
// Example:
//   var stack Stack
//
//   network := medialive.NewNetwork(stack, jsii.String("Network"), &NetworkProps{
//   	NetworkName: jsii.String("on-prem-network"),
//   	IpPools: []*string{
//   		jsii.String("10.0.0.0/24"),
//   	},
//   	Routes: []NetworkRoute{
//   		&NetworkRoute{
//   			Cidr: jsii.String("0.0.0.0/0"),
//   			Gateway: jsii.String("10.0.0.1"),
//   		},
//   	},
//   })
//
// Experimental.
type NetworkProps struct {
	// The list of IP address CIDR pools for the network.
	// Experimental.
	IpPools *[]*string `field:"required" json:"ipPools" yaml:"ipPools"`
	// The name of the network.
	// Default: - auto-generated name.
	//
	// Experimental.
	NetworkName *string `field:"optional" json:"networkName" yaml:"networkName"`
	// The routes for the network.
	// Default: - no routes.
	//
	// Experimental.
	Routes *[]*NetworkRoute `field:"optional" json:"routes" yaml:"routes"`
	// Tags to add to the network.
	// Default: - no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

