package awsdirectconnect


// Information about a BGP peer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bgpPeerProperty := &BgpPeerProperty{
//   	AddressFamily: jsii.String("addressFamily"),
//   	Asn: jsii.String("asn"),
//
//   	// the properties below are optional
//   	AmazonAddress: jsii.String("amazonAddress"),
//   	AuthKey: jsii.String("authKey"),
//   	BgpPeerId: jsii.String("bgpPeerId"),
//   	CustomerAddress: jsii.String("customerAddress"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directconnect-publicvirtualinterface-bgppeer.html
//
type CfnPublicVirtualInterface_BgpPeerProperty struct {
	// The address family for the BGP peer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directconnect-publicvirtualinterface-bgppeer.html#cfn-directconnect-publicvirtualinterface-bgppeer-addressfamily
	//
	AddressFamily *string `field:"required" json:"addressFamily" yaml:"addressFamily"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directconnect-publicvirtualinterface-bgppeer.html#cfn-directconnect-publicvirtualinterface-bgppeer-asn
	//
	Asn *string `field:"required" json:"asn" yaml:"asn"`
	// The IP address assigned to the Amazon interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directconnect-publicvirtualinterface-bgppeer.html#cfn-directconnect-publicvirtualinterface-bgppeer-amazonaddress
	//
	AmazonAddress *string `field:"optional" json:"amazonAddress" yaml:"amazonAddress"`
	// The authentication key for BGP configuration.
	//
	// This string has a minimum length of 6 characters and and a maximum length of 80 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directconnect-publicvirtualinterface-bgppeer.html#cfn-directconnect-publicvirtualinterface-bgppeer-authkey
	//
	AuthKey *string `field:"optional" json:"authKey" yaml:"authKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directconnect-publicvirtualinterface-bgppeer.html#cfn-directconnect-publicvirtualinterface-bgppeer-bgppeerid
	//
	BgpPeerId *string `field:"optional" json:"bgpPeerId" yaml:"bgpPeerId"`
	// The IP address assigned to the customer interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directconnect-publicvirtualinterface-bgppeer.html#cfn-directconnect-publicvirtualinterface-bgppeer-customeraddress
	//
	CustomerAddress *string `field:"optional" json:"customerAddress" yaml:"customerAddress"`
}

