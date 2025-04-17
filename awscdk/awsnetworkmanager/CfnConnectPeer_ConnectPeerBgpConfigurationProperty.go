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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-connectpeer-connectpeerbgpconfiguration.html
//
type CfnConnectPeer_ConnectPeerBgpConfigurationProperty struct {
	// The address of a core network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-connectpeer-connectpeerbgpconfiguration.html#cfn-networkmanager-connectpeer-connectpeerbgpconfiguration-corenetworkaddress
	//
	CoreNetworkAddress *string `field:"optional" json:"coreNetworkAddress" yaml:"coreNetworkAddress"`
	// The ASN of the Coret Network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-connectpeer-connectpeerbgpconfiguration.html#cfn-networkmanager-connectpeer-connectpeerbgpconfiguration-corenetworkasn
	//
	CoreNetworkAsn *float64 `field:"optional" json:"coreNetworkAsn" yaml:"coreNetworkAsn"`
	// The address of a core network Connect peer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-connectpeer-connectpeerbgpconfiguration.html#cfn-networkmanager-connectpeer-connectpeerbgpconfiguration-peeraddress
	//
	PeerAddress *string `field:"optional" json:"peerAddress" yaml:"peerAddress"`
	// The ASN of the Connect peer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-connectpeer-connectpeerbgpconfiguration.html#cfn-networkmanager-connectpeer-connectpeerbgpconfiguration-peerasn
	//
	PeerAsn *float64 `field:"optional" json:"peerAsn" yaml:"peerAsn"`
}

