package mixinsawsec2


// Properties for CfnClientVpnRoutePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClientVpnRouteMixinProps := &CfnClientVpnRouteMixinProps{
//   	ClientVpnEndpointId: jsii.String("clientVpnEndpointId"),
//   	Description: jsii.String("description"),
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	TargetVpcSubnetId: jsii.String("targetVpcSubnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnroute.html
//
type CfnClientVpnRouteMixinProps struct {
	// The ID of the Client VPN endpoint to which to add the route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnroute.html#cfn-ec2-clientvpnroute-clientvpnendpointid
	//
	ClientVpnEndpointId *string `field:"optional" json:"clientVpnEndpointId" yaml:"clientVpnEndpointId"`
	// A brief description of the route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnroute.html#cfn-ec2-clientvpnroute-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IPv4 address range, in CIDR notation, of the route destination. For example:.
	//
	// - To add a route for Internet access, enter `0.0.0.0/0`
	// - To add a route for a peered VPC, enter the peered VPC's IPv4 CIDR range
	// - To add a route for an on-premises network, enter the AWS Site-to-Site VPN connection's IPv4 CIDR range
	// - To add a route for the local network, enter the client CIDR range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnroute.html#cfn-ec2-clientvpnroute-destinationcidrblock
	//
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The ID of the subnet through which you want to route traffic.
	//
	// The specified subnet must be an existing target network of the Client VPN endpoint.
	//
	// Alternatively, if you're adding a route for the local network, specify `local` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnroute.html#cfn-ec2-clientvpnroute-targetvpcsubnetid
	//
	TargetVpcSubnetId *string `field:"optional" json:"targetVpcSubnetId" yaml:"targetVpcSubnetId"`
}

