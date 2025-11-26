package previewawsec2mixins


// Properties for CfnVPNConnectionRoutePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPNConnectionRouteMixinProps := &CfnVPNConnectionRouteMixinProps{
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	VpnConnectionId: jsii.String("vpnConnectionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnectionroute.html
//
type CfnVPNConnectionRouteMixinProps struct {
	// The CIDR block associated with the local subnet of the customer network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnectionroute.html#cfn-ec2-vpnconnectionroute-destinationcidrblock
	//
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The ID of the VPN connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnectionroute.html#cfn-ec2-vpnconnectionroute-vpnconnectionid
	//
	VpnConnectionId *string `field:"optional" json:"vpnConnectionId" yaml:"vpnConnectionId"`
}

