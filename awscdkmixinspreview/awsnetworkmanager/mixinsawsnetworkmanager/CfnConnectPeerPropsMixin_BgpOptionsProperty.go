package mixinsawsnetworkmanager


// Describes the BGP options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bgpOptionsProperty := &BgpOptionsProperty{
//   	PeerAsn: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-connectpeer-bgpoptions.html
//
type CfnConnectPeerPropsMixin_BgpOptionsProperty struct {
	// The Peer ASN of the BGP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-connectpeer-bgpoptions.html#cfn-networkmanager-connectpeer-bgpoptions-peerasn
	//
	PeerAsn *float64 `field:"optional" json:"peerAsn" yaml:"peerAsn"`
}

