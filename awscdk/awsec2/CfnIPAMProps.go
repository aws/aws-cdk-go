package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIPAM`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPAMProps := &CfnIPAMProps{
//   	DefaultResourceDiscoveryOrganizationalUnitExclusions: []interface{}{
//   		&IpamOrganizationalUnitExclusionProperty{
//   			OrganizationsEntityPath: jsii.String("organizationsEntityPath"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	EnablePrivateGua: jsii.Boolean(false),
//   	MeteredAccount: jsii.String("meteredAccount"),
//   	OperatingRegions: []interface{}{
//   		&IpamOperatingRegionProperty{
//   			RegionName: jsii.String("regionName"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tier: jsii.String("tier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipam.html
//
type CfnIPAMProps struct {
	// If your IPAM is integrated with AWS Organizations, you can exclude an [organizational unit (OU)](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#organizationalunit) from being managed by IPAM. When you exclude an OU, IPAM will not manage the IP addresses in accounts in that OU. For more information, see [Exclude organizational units from IPAM](https://docs.aws.amazon.com/vpc/latest/ipam/exclude-ous.html) in the *Amazon Virtual Private Cloud IP Address Manager User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipam.html#cfn-ec2-ipam-defaultresourcediscoveryorganizationalunitexclusions
	//
	DefaultResourceDiscoveryOrganizationalUnitExclusions interface{} `field:"optional" json:"defaultResourceDiscoveryOrganizationalUnitExclusions" yaml:"defaultResourceDiscoveryOrganizationalUnitExclusions"`
	// The description for the IPAM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipam.html#cfn-ec2-ipam-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable this option to use your own GUA ranges as private IPv6 addresses.
	//
	// This option is disabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipam.html#cfn-ec2-ipam-enableprivategua
	//
	EnablePrivateGua interface{} `field:"optional" json:"enablePrivateGua" yaml:"enablePrivateGua"`
	// A metered account is an account that is charged for active IP addresses managed in IPAM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipam.html#cfn-ec2-ipam-meteredaccount
	//
	MeteredAccount *string `field:"optional" json:"meteredAccount" yaml:"meteredAccount"`
	// The operating Regions for an IPAM.
	//
	// Operating Regions are AWS Regions where the IPAM is allowed to manage IP address CIDRs. IPAM only discovers and monitors resources in the AWS Regions you select as operating Regions.
	//
	// For more information about operating Regions, see [Create an IPAM](https://docs.aws.amazon.com//vpc/latest/ipam/create-ipam.html) in the *Amazon VPC IPAM User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipam.html#cfn-ec2-ipam-operatingregions
	//
	OperatingRegions interface{} `field:"optional" json:"operatingRegions" yaml:"operatingRegions"`
	// The key/value combination of a tag assigned to the resource.
	//
	// Use the tag key in the filter name and the tag value as the filter value. For example, to find all resources that have a tag with the key `Owner` and the value `TeamA` , specify `tag:Owner` for the filter name and `TeamA` for the filter value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipam.html#cfn-ec2-ipam-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// IPAM is offered in a Free Tier and an Advanced Tier.
	//
	// For more information about the features available in each tier and the costs associated with the tiers, see the [VPC IPAM product pricing page](https://docs.aws.amazon.com//vpc/pricing/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipam.html#cfn-ec2-ipam-tier
	//
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

