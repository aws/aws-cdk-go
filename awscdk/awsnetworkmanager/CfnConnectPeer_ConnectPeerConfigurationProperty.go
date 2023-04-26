package awsnetworkmanager


// Describes a core network Connect peer configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectPeerConfigurationProperty := &ConnectPeerConfigurationProperty{
//   	BgpConfigurations: []interface{}{
//   		&ConnectPeerBgpConfigurationProperty{
//   			CoreNetworkAddress: jsii.String("coreNetworkAddress"),
//   			CoreNetworkAsn: jsii.Number(123),
//   			PeerAddress: jsii.String("peerAddress"),
//   			PeerAsn: jsii.Number(123),
//   		},
//   	},
//   	CoreNetworkAddress: jsii.String("coreNetworkAddress"),
//   	InsideCidrBlocks: []*string{
//   		jsii.String("insideCidrBlocks"),
//   	},
//   	PeerAddress: jsii.String("peerAddress"),
//   	Protocol: jsii.String("protocol"),
//   }
//
type CfnConnectPeer_ConnectPeerConfigurationProperty struct {
	// The Connect peer BGP configurations.
	BgpConfigurations interface{} `field:"optional" json:"bgpConfigurations" yaml:"bgpConfigurations"`
	// The IP address of a core network.
	CoreNetworkAddress *string `field:"optional" json:"coreNetworkAddress" yaml:"coreNetworkAddress"`
	// The inside IP addresses used for a Connect peer configuration.
	InsideCidrBlocks *[]*string `field:"optional" json:"insideCidrBlocks" yaml:"insideCidrBlocks"`
	// The IP address of the Connect peer.
	PeerAddress *string `field:"optional" json:"peerAddress" yaml:"peerAddress"`
	// The protocol used for a Connect peer configuration.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

