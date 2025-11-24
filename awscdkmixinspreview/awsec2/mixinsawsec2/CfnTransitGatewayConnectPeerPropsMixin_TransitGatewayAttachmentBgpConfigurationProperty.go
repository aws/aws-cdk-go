package mixinsawsec2


// The BGP configuration information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transitGatewayAttachmentBgpConfigurationProperty := &TransitGatewayAttachmentBgpConfigurationProperty{
//   	BgpStatus: jsii.String("bgpStatus"),
//   	PeerAddress: jsii.String("peerAddress"),
//   	PeerAsn: jsii.Number(123),
//   	TransitGatewayAddress: jsii.String("transitGatewayAddress"),
//   	TransitGatewayAsn: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayconnectpeer-transitgatewayattachmentbgpconfiguration.html
//
type CfnTransitGatewayConnectPeerPropsMixin_TransitGatewayAttachmentBgpConfigurationProperty struct {
	// The BGP status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayconnectpeer-transitgatewayattachmentbgpconfiguration.html#cfn-ec2-transitgatewayconnectpeer-transitgatewayattachmentbgpconfiguration-bgpstatus
	//
	BgpStatus *string `field:"optional" json:"bgpStatus" yaml:"bgpStatus"`
	// The interior BGP peer IP address for the appliance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayconnectpeer-transitgatewayattachmentbgpconfiguration.html#cfn-ec2-transitgatewayconnectpeer-transitgatewayattachmentbgpconfiguration-peeraddress
	//
	PeerAddress *string `field:"optional" json:"peerAddress" yaml:"peerAddress"`
	// The peer Autonomous System Number (ASN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayconnectpeer-transitgatewayattachmentbgpconfiguration.html#cfn-ec2-transitgatewayconnectpeer-transitgatewayattachmentbgpconfiguration-peerasn
	//
	PeerAsn *float64 `field:"optional" json:"peerAsn" yaml:"peerAsn"`
	// The interior BGP peer IP address for the transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayconnectpeer-transitgatewayattachmentbgpconfiguration.html#cfn-ec2-transitgatewayconnectpeer-transitgatewayattachmentbgpconfiguration-transitgatewayaddress
	//
	TransitGatewayAddress *string `field:"optional" json:"transitGatewayAddress" yaml:"transitGatewayAddress"`
	// The transit gateway Autonomous System Number (ASN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayconnectpeer-transitgatewayattachmentbgpconfiguration.html#cfn-ec2-transitgatewayconnectpeer-transitgatewayattachmentbgpconfiguration-transitgatewayasn
	//
	TransitGatewayAsn *float64 `field:"optional" json:"transitGatewayAsn" yaml:"transitGatewayAsn"`
}

