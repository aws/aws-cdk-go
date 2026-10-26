package awsnetworkfirewall


// A mapping to a NAT gateway that a proxy-mode firewall proxies egress traffic through.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   natGatewayMappingProperty := &NatGatewayMappingProperty{
//   	NatGatewayId: jsii.String("natGatewayId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewall-natgatewaymapping.html
//
type CfnFirewall_NatGatewayMappingProperty struct {
	// The unique identifier of a NAT gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewall-natgatewaymapping.html#cfn-networkfirewall-firewall-natgatewaymapping-natgatewayid
	//
	NatGatewayId *string `field:"required" json:"natGatewayId" yaml:"natGatewayId"`
}

