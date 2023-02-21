package awsnetworkmanager


// Describes the BGP options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bgpOptionsProperty := &bgpOptionsProperty{
//   	peerAsn: jsii.Number(123),
//   }
//
type CfnConnectPeer_BgpOptionsProperty struct {
	// The Peer ASN of the BGP.
	PeerAsn *float64 `field:"optional" json:"peerAsn" yaml:"peerAsn"`
}

