package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIPAMPool`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPAMPoolProps := &CfnIPAMPoolProps{
//   	AddressFamily: jsii.String("addressFamily"),
//   	IpamScopeId: jsii.String("ipamScopeId"),
//
//   	// the properties below are optional
//   	AllocationDefaultNetmaskLength: jsii.Number(123),
//   	AllocationMaxNetmaskLength: jsii.Number(123),
//   	AllocationMinNetmaskLength: jsii.Number(123),
//   	AllocationResourceTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	AutoImport: jsii.Boolean(false),
//   	AwsService: jsii.String("awsService"),
//   	Description: jsii.String("description"),
//   	Locale: jsii.String("locale"),
//   	ProvisionedCidrs: []interface{}{
//   		&ProvisionedCidrProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   	PublicIpSource: jsii.String("publicIpSource"),
//   	PubliclyAdvertisable: jsii.Boolean(false),
//   	SourceIpamPoolId: jsii.String("sourceIpamPoolId"),
//   	SourceResource: &SourceResourceProperty{
//   		ResourceId: jsii.String("resourceId"),
//   		ResourceOwner: jsii.String("resourceOwner"),
//   		ResourceRegion: jsii.String("resourceRegion"),
//   		ResourceType: jsii.String("resourceType"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html
//
type CfnIPAMPoolProps struct {
	// The address family of the pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-addressfamily
	//
	AddressFamily *string `field:"required" json:"addressFamily" yaml:"addressFamily"`
	// The ID of the scope in which you would like to create the IPAM pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-ipamscopeid
	//
	IpamScopeId *string `field:"required" json:"ipamScopeId" yaml:"ipamScopeId"`
	// The default netmask length for allocations added to this pool.
	//
	// If, for example, the CIDR assigned to this pool is 10.0.0.0/8 and you enter 16 here, new allocations will default to 10.0.0.0/16.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-allocationdefaultnetmasklength
	//
	AllocationDefaultNetmaskLength *float64 `field:"optional" json:"allocationDefaultNetmaskLength" yaml:"allocationDefaultNetmaskLength"`
	// The maximum netmask length possible for CIDR allocations in this IPAM pool to be compliant.
	//
	// The maximum netmask length must be greater than the minimum netmask length. Possible netmask lengths for IPv4 addresses are 0 - 32. Possible netmask lengths for IPv6 addresses are 0 - 128.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-allocationmaxnetmasklength
	//
	AllocationMaxNetmaskLength *float64 `field:"optional" json:"allocationMaxNetmaskLength" yaml:"allocationMaxNetmaskLength"`
	// The minimum netmask length required for CIDR allocations in this IPAM pool to be compliant.
	//
	// The minimum netmask length must be less than the maximum netmask length. Possible netmask lengths for IPv4 addresses are 0 - 32. Possible netmask lengths for IPv6 addresses are 0 - 128.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-allocationminnetmasklength
	//
	AllocationMinNetmaskLength *float64 `field:"optional" json:"allocationMinNetmaskLength" yaml:"allocationMinNetmaskLength"`
	// Tags that are required for resources that use CIDRs from this IPAM pool.
	//
	// Resources that do not have these tags will not be allowed to allocate space from the pool. If the resources have their tags changed after they have allocated space or if the allocation tagging requirements are changed on the pool, the resource may be marked as noncompliant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-allocationresourcetags
	//
	AllocationResourceTags interface{} `field:"optional" json:"allocationResourceTags" yaml:"allocationResourceTags"`
	// If selected, IPAM will continuously look for resources within the CIDR range of this pool and automatically import them as allocations into your IPAM.
	//
	// The CIDRs that will be allocated for these resources must not already be allocated to other resources in order for the import to succeed. IPAM will import a CIDR regardless of its compliance with the pool's allocation rules, so a resource might be imported and subsequently marked as noncompliant. If IPAM discovers multiple CIDRs that overlap, IPAM will import the largest CIDR only. If IPAM discovers multiple CIDRs with matching CIDRs, IPAM will randomly import one of them only.
	//
	// A locale must be set on the pool for this feature to work.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-autoimport
	//
	AutoImport interface{} `field:"optional" json:"autoImport" yaml:"autoImport"`
	// Limits which service in AWS that the pool can be used in.
	//
	// "ec2", for example, allows users to use space for Elastic IP addresses and VPCs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-awsservice
	//
	AwsService *string `field:"optional" json:"awsService" yaml:"awsService"`
	// The description of the IPAM pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The locale of the IPAM pool.
	//
	// In IPAM, the locale is the AWS Region where you want to make an IPAM pool available for allocations. Only resources in the same Region as the locale of the pool can get IP address allocations from the pool. You can only allocate a CIDR for a VPC, for example, from an IPAM pool that shares a locale with the VPC’s Region. Note that once you choose a Locale for a pool, you cannot modify it. If you choose an AWS Region for locale that has not been configured as an operating Region for the IPAM, you'll get an error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Information about the CIDRs provisioned to an IPAM pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-provisionedcidrs
	//
	ProvisionedCidrs interface{} `field:"optional" json:"provisionedCidrs" yaml:"provisionedCidrs"`
	// The IP address source for pools in the public scope.
	//
	// Only used for provisioning IP address CIDRs to pools in the public scope. Default is `BYOIP` . For more information, see [Create IPv6 pools](https://docs.aws.amazon.com//vpc/latest/ipam/intro-create-ipv6-pools.html) in the *Amazon VPC IPAM User Guide* . By default, you can add only one Amazon-provided IPv6 CIDR block to a top-level IPv6 pool. For information on increasing the default limit, see [Quotas for your IPAM](https://docs.aws.amazon.com//vpc/latest/ipam/quotas-ipam.html) in the *Amazon VPC IPAM User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-publicipsource
	//
	PublicIpSource *string `field:"optional" json:"publicIpSource" yaml:"publicIpSource"`
	// Determines if a pool is publicly advertisable.
	//
	// This option is not available for pools with AddressFamily set to `ipv4` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-publiclyadvertisable
	//
	PubliclyAdvertisable interface{} `field:"optional" json:"publiclyAdvertisable" yaml:"publiclyAdvertisable"`
	// The ID of the source IPAM pool.
	//
	// You can use this option to create an IPAM pool within an existing source pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-sourceipampoolid
	//
	SourceIpamPoolId *string `field:"optional" json:"sourceIpamPoolId" yaml:"sourceIpamPoolId"`
	// The resource used to provision CIDRs to a resource planning pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-sourceresource
	//
	SourceResource interface{} `field:"optional" json:"sourceResource" yaml:"sourceResource"`
	// The key/value combination of a tag assigned to the resource.
	//
	// Use the tag key in the filter name and the tag value as the filter value. For example, to find all resources that have a tag with the key `Owner` and the value `TeamA` , specify `tag:Owner` for the filter name and `TeamA` for the filter value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

