package awsnetworkfirewall

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVpcEndpointAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcEndpointAssociationProps := &CfnVpcEndpointAssociationProps{
//   	FirewallArn: jsii.String("firewallArn"),
//   	SubnetMapping: &SubnetMappingProperty{
//   		SubnetId: jsii.String("subnetId"),
//
//   		// the properties below are optional
//   		IpAddressType: jsii.String("ipAddressType"),
//   	},
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-vpcendpointassociation.html
//
type CfnVpcEndpointAssociationProps struct {
	// The Amazon Resource Name (ARN) of the firewall.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-vpcendpointassociation.html#cfn-networkfirewall-vpcendpointassociation-firewallarn
	//
	FirewallArn *string `field:"required" json:"firewallArn" yaml:"firewallArn"`
	// The ID for a subnet that's used in an association with a firewall.
	//
	// This is used in `CreateFirewall` , `AssociateSubnets` , and `CreateVpcEndpointAssociation` . AWS Network Firewall creates an instance of the associated firewall in each subnet that you specify, to filter traffic in the subnet's Availability Zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-vpcendpointassociation.html#cfn-networkfirewall-vpcendpointassociation-subnetmapping
	//
	SubnetMapping interface{} `field:"required" json:"subnetMapping" yaml:"subnetMapping"`
	// The unique identifier of the VPC for the endpoint association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-vpcendpointassociation.html#cfn-networkfirewall-vpcendpointassociation-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// A description of the VPC endpoint association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-vpcendpointassociation.html#cfn-networkfirewall-vpcendpointassociation-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The key:value pairs to associate with the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-vpcendpointassociation.html#cfn-networkfirewall-vpcendpointassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

