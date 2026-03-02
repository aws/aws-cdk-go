package previewawsdirectconnectmixins


// Information about a BGP peer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bgpPeerProperty := &BgpPeerProperty{
//   	AddressFamily: jsii.String("addressFamily"),
//   	AmazonAddress: jsii.String("amazonAddress"),
//   	Asn: jsii.String("asn"),
//   	AuthKey: jsii.String("authKey"),
//   	BgpPeerId: jsii.String("bgpPeerId"),
//   	CustomerAddress: jsii.String("customerAddress"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directconnect-privatevirtualinterface-bgppeer.html
//
type CfnPrivateVirtualInterfacePropsMixin_BgpPeerProperty struct {
	// The address family for the BGP peer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directconnect-privatevirtualinterface-bgppeer.html#cfn-directconnect-privatevirtualinterface-bgppeer-addressfamily
	//
	AddressFamily *string `field:"optional" json:"addressFamily" yaml:"addressFamily"`
	// The IP address assigned to the Amazon interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directconnect-privatevirtualinterface-bgppeer.html#cfn-directconnect-privatevirtualinterface-bgppeer-amazonaddress
	//
	AmazonAddress *string `field:"optional" json:"amazonAddress" yaml:"amazonAddress"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directconnect-privatevirtualinterface-bgppeer.html#cfn-directconnect-privatevirtualinterface-bgppeer-asn
	//
	Asn *string `field:"optional" json:"asn" yaml:"asn"`
	// The authentication key for BGP configuration.
	//
	// This string has a minimum length of 6 characters and and a maximum length of 80 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directconnect-privatevirtualinterface-bgppeer.html#cfn-directconnect-privatevirtualinterface-bgppeer-authkey
	//
	AuthKey *string `field:"optional" json:"authKey" yaml:"authKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directconnect-privatevirtualinterface-bgppeer.html#cfn-directconnect-privatevirtualinterface-bgppeer-bgppeerid
	//
	BgpPeerId *string `field:"optional" json:"bgpPeerId" yaml:"bgpPeerId"`
	// The IP address assigned to the customer interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directconnect-privatevirtualinterface-bgppeer.html#cfn-directconnect-privatevirtualinterface-bgppeer-customeraddress
	//
	CustomerAddress *string `field:"optional" json:"customerAddress" yaml:"customerAddress"`
}

