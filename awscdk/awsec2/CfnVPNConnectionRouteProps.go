package awsec2


// Properties for defining a `CfnVPNConnectionRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPNConnectionRouteProps := &CfnVPNConnectionRouteProps{
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	VpnConnectionId: jsii.String("vpnConnectionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnectionroute.html
//
type CfnVPNConnectionRouteProps struct {
	// The CIDR block associated with the local subnet of the customer network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnectionroute.html#cfn-ec2-vpnconnectionroute-destinationcidrblock
	//
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The ID of the VPN connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnectionroute.html#cfn-ec2-vpnconnectionroute-vpnconnectionid
	//
	VpnConnectionId *string `field:"required" json:"vpnConnectionId" yaml:"vpnConnectionId"`
}

