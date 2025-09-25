package awsnetworkfirewall


// The ID for a subnet that's used in an association with a firewall.
//
// This is used in `CreateFirewall` , `AssociateSubnets` , and `CreateVpcEndpointAssociation` . AWS Network Firewall creates an instance of the associated firewall in each subnet that you specify, to filter traffic in the subnet's Availability Zone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetMappingProperty := &SubnetMappingProperty{
//   	SubnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	IpAddressType: jsii.String("ipAddressType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-vpcendpointassociation-subnetmapping.html
//
type CfnVpcEndpointAssociation_SubnetMappingProperty struct {
	// The unique identifier for the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-vpcendpointassociation-subnetmapping.html#cfn-networkfirewall-vpcendpointassociation-subnetmapping-subnetid
	//
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The subnet's IP address type.
	//
	// You can't change the IP address type after you create the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-vpcendpointassociation-subnetmapping.html#cfn-networkfirewall-vpcendpointassociation-subnetmapping-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
}

