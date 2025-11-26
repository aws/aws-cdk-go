package previewawsec2mixins


// Describes the Connect peer details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transitGatewayConnectPeerConfigurationProperty := &TransitGatewayConnectPeerConfigurationProperty{
//   	BgpConfigurations: []interface{}{
//   		&TransitGatewayAttachmentBgpConfigurationProperty{
//   			BgpStatus: jsii.String("bgpStatus"),
//   			PeerAddress: jsii.String("peerAddress"),
//   			PeerAsn: jsii.Number(123),
//   			TransitGatewayAddress: jsii.String("transitGatewayAddress"),
//   			TransitGatewayAsn: jsii.Number(123),
//   		},
//   	},
//   	InsideCidrBlocks: []*string{
//   		jsii.String("insideCidrBlocks"),
//   	},
//   	PeerAddress: jsii.String("peerAddress"),
//   	Protocol: jsii.String("protocol"),
//   	TransitGatewayAddress: jsii.String("transitGatewayAddress"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayconnectpeer-transitgatewayconnectpeerconfiguration.html
//
type CfnTransitGatewayConnectPeerPropsMixin_TransitGatewayConnectPeerConfigurationProperty struct {
	// The BGP configuration details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayconnectpeer-transitgatewayconnectpeerconfiguration.html#cfn-ec2-transitgatewayconnectpeer-transitgatewayconnectpeerconfiguration-bgpconfigurations
	//
	BgpConfigurations interface{} `field:"optional" json:"bgpConfigurations" yaml:"bgpConfigurations"`
	// The range of interior BGP peer IP addresses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayconnectpeer-transitgatewayconnectpeerconfiguration.html#cfn-ec2-transitgatewayconnectpeer-transitgatewayconnectpeerconfiguration-insidecidrblocks
	//
	InsideCidrBlocks *[]*string `field:"optional" json:"insideCidrBlocks" yaml:"insideCidrBlocks"`
	// The Connect peer IP address on the appliance side of the tunnel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayconnectpeer-transitgatewayconnectpeerconfiguration.html#cfn-ec2-transitgatewayconnectpeer-transitgatewayconnectpeerconfiguration-peeraddress
	//
	PeerAddress *string `field:"optional" json:"peerAddress" yaml:"peerAddress"`
	// The tunnel protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayconnectpeer-transitgatewayconnectpeerconfiguration.html#cfn-ec2-transitgatewayconnectpeer-transitgatewayconnectpeerconfiguration-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The Connect peer IP address on the transit gateway side of the tunnel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayconnectpeer-transitgatewayconnectpeerconfiguration.html#cfn-ec2-transitgatewayconnectpeer-transitgatewayconnectpeerconfiguration-transitgatewayaddress
	//
	TransitGatewayAddress *string `field:"optional" json:"transitGatewayAddress" yaml:"transitGatewayAddress"`
}

