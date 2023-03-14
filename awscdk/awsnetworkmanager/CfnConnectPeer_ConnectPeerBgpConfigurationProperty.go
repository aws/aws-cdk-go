package awsnetworkmanager


// Describes a core network BGP configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectPeerBgpConfigurationProperty := &ConnectPeerBgpConfigurationProperty{
//   	CoreNetworkAddress: jsii.String("coreNetworkAddress"),
//   	CoreNetworkAsn: jsii.Number(123),
//   	PeerAddress: jsii.String("peerAddress"),
//   	PeerAsn: jsii.Number(123),
//   }
//
type CfnConnectPeer_ConnectPeerBgpConfigurationProperty struct {
	// The address of a core network.
	CoreNetworkAddress *string `field:"optional" json:"coreNetworkAddress" yaml:"coreNetworkAddress"`
	// The ASN of the Coret Network.
	CoreNetworkAsn *float64 `field:"optional" json:"coreNetworkAsn" yaml:"coreNetworkAsn"`
	// The address of a core network Connect peer.
	PeerAddress *string `field:"optional" json:"peerAddress" yaml:"peerAddress"`
	// The ASN of the Connect peer.
	PeerAsn *float64 `field:"optional" json:"peerAsn" yaml:"peerAsn"`
}

