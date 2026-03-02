package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIPAMPrefixListResolver`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPAMPrefixListResolverProps := &CfnIPAMPrefixListResolverProps{
//   	AddressFamily: jsii.String("addressFamily"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	IpamId: jsii.String("ipamId"),
//   	Rules: []interface{}{
//   		&IpamPrefixListResolverRuleProperty{
//   			RuleType: jsii.String("ruleType"),
//
//   			// the properties below are optional
//   			Conditions: []interface{}{
//   				&IpamPrefixListResolverRuleConditionProperty{
//   					Operation: jsii.String("operation"),
//
//   					// the properties below are optional
//   					Cidr: jsii.String("cidr"),
//   					IpamPoolId: jsii.String("ipamPoolId"),
//   					ResourceId: jsii.String("resourceId"),
//   					ResourceOwner: jsii.String("resourceOwner"),
//   					ResourceRegion: jsii.String("resourceRegion"),
//   					ResourceTag: &CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			IpamScopeId: jsii.String("ipamScopeId"),
//   			ResourceType: jsii.String("resourceType"),
//   			StaticCidr: jsii.String("staticCidr"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamprefixlistresolver.html
//
type CfnIPAMPrefixListResolverProps struct {
	// The address family of the address space in this Prefix List Resolver.
	//
	// Either IPv4 or IPv6.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamprefixlistresolver.html#cfn-ec2-ipamprefixlistresolver-addressfamily
	//
	AddressFamily *string `field:"required" json:"addressFamily" yaml:"addressFamily"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamprefixlistresolver.html#cfn-ec2-ipamprefixlistresolver-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Id of the IPAM this Prefix List Resolver is a part of.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamprefixlistresolver.html#cfn-ec2-ipamprefixlistresolver-ipamid
	//
	IpamId *string `field:"optional" json:"ipamId" yaml:"ipamId"`
	// Rules define the business logic for selecting CIDRs from IPAM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamprefixlistresolver.html#cfn-ec2-ipamprefixlistresolver-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamprefixlistresolver.html#cfn-ec2-ipamprefixlistresolver-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

