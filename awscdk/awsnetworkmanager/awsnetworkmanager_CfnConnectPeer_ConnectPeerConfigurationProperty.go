package awsnetworkmanager


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectPeerConfigurationProperty := &connectPeerConfigurationProperty{
//   	bgpConfigurations: []interface{}{
//   		&connectPeerBgpConfigurationProperty{
//   			coreNetworkAddress: jsii.String("coreNetworkAddress"),
//   			coreNetworkAsn: jsii.Number(123),
//   			peerAddress: jsii.String("peerAddress"),
//   			peerAsn: jsii.Number(123),
//   		},
//   	},
//   	coreNetworkAddress: jsii.String("coreNetworkAddress"),
//   	insideCidrBlocks: []*string{
//   		jsii.String("insideCidrBlocks"),
//   	},
//   	peerAddress: jsii.String("peerAddress"),
//   	protocol: jsii.String("protocol"),
//   }
//
type CfnConnectPeer_ConnectPeerConfigurationProperty struct {
	// `CfnConnectPeer.ConnectPeerConfigurationProperty.BgpConfigurations`.
	BgpConfigurations interface{} `field:"optional" json:"bgpConfigurations" yaml:"bgpConfigurations"`
	// `CfnConnectPeer.ConnectPeerConfigurationProperty.CoreNetworkAddress`.
	CoreNetworkAddress *string `field:"optional" json:"coreNetworkAddress" yaml:"coreNetworkAddress"`
	// `CfnConnectPeer.ConnectPeerConfigurationProperty.InsideCidrBlocks`.
	InsideCidrBlocks *[]*string `field:"optional" json:"insideCidrBlocks" yaml:"insideCidrBlocks"`
	// `CfnConnectPeer.ConnectPeerConfigurationProperty.PeerAddress`.
	PeerAddress *string `field:"optional" json:"peerAddress" yaml:"peerAddress"`
	// `CfnConnectPeer.ConnectPeerConfigurationProperty.Protocol`.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

