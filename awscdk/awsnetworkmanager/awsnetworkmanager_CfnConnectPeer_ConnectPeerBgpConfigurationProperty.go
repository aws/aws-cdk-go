package awsnetworkmanager


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectPeerBgpConfigurationProperty := &connectPeerBgpConfigurationProperty{
//   	coreNetworkAddress: jsii.String("coreNetworkAddress"),
//   	coreNetworkAsn: jsii.Number(123),
//   	peerAddress: jsii.String("peerAddress"),
//   	peerAsn: jsii.Number(123),
//   }
//
type CfnConnectPeer_ConnectPeerBgpConfigurationProperty struct {
	// `CfnConnectPeer.ConnectPeerBgpConfigurationProperty.CoreNetworkAddress`.
	CoreNetworkAddress *string `field:"optional" json:"coreNetworkAddress" yaml:"coreNetworkAddress"`
	// `CfnConnectPeer.ConnectPeerBgpConfigurationProperty.CoreNetworkAsn`.
	CoreNetworkAsn *float64 `field:"optional" json:"coreNetworkAsn" yaml:"coreNetworkAsn"`
	// `CfnConnectPeer.ConnectPeerBgpConfigurationProperty.PeerAddress`.
	PeerAddress *string `field:"optional" json:"peerAddress" yaml:"peerAddress"`
	// `CfnConnectPeer.ConnectPeerBgpConfigurationProperty.PeerAsn`.
	PeerAsn *float64 `field:"optional" json:"peerAsn" yaml:"peerAsn"`
}

