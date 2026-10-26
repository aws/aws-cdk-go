package awsnetworkfirewall


// The VPC and subnets for a proxy-mode firewall's endpoint.
//
// Provide this instead of the top-level VpcId and SubnetMappings when NoSourcePreservation is true.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   vpcEndpointProperty := &VpcEndpointProperty{
//   	SubnetMappings: []interface{}{
//   		&SubnetMappingProperty{
//   			IpAddressType: jsii.String("ipAddressType"),
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewall-vpcendpoint.html
//
type CfnFirewallPropsMixin_VpcEndpointProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewall-vpcendpoint.html#cfn-networkfirewall-firewall-vpcendpoint-subnetmappings
	//
	SubnetMappings interface{} `field:"optional" json:"subnetMappings" yaml:"subnetMappings"`
	// The unique identifier of the VPC for the firewall endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewall-vpcendpoint.html#cfn-networkfirewall-firewall-vpcendpoint-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

