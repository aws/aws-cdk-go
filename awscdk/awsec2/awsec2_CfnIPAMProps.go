package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnIPAM`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPAMProps := &CfnIPAMProps{
//   	DefaultResourceDiscoveryAssociationId: jsii.String("defaultResourceDiscoveryAssociationId"),
//   	DefaultResourceDiscoveryId: jsii.String("defaultResourceDiscoveryId"),
//   	Description: jsii.String("description"),
//   	OperatingRegions: []interface{}{
//   		&IpamOperatingRegionProperty{
//   			RegionName: jsii.String("regionName"),
//   		},
//   	},
//   	ResourceDiscoveryAssociationCount: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnIPAMProps struct {
	// The IPAM's default resource discovery association ID.
	DefaultResourceDiscoveryAssociationId *string `field:"optional" json:"defaultResourceDiscoveryAssociationId" yaml:"defaultResourceDiscoveryAssociationId"`
	// The IPAM's default resource discovery ID.
	DefaultResourceDiscoveryId *string `field:"optional" json:"defaultResourceDiscoveryId" yaml:"defaultResourceDiscoveryId"`
	// The description for the IPAM.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The operating Regions for an IPAM.
	//
	// Operating Regions are AWS Regions where the IPAM is allowed to manage IP address CIDRs. IPAM only discovers and monitors resources in the AWS Regions you select as operating Regions.
	//
	// For more information about operating Regions, see [Create an IPAM](https://docs.aws.amazon.com//vpc/latest/ipam/create-ipam.html) in the *Amazon VPC IPAM User Guide* .
	OperatingRegions interface{} `field:"optional" json:"operatingRegions" yaml:"operatingRegions"`
	// The IPAM's resource discovery association count.
	ResourceDiscoveryAssociationCount *float64 `field:"optional" json:"resourceDiscoveryAssociationCount" yaml:"resourceDiscoveryAssociationCount"`
	// The key/value combination of a tag assigned to the resource.
	//
	// Use the tag key in the filter name and the tag value as the filter value. For example, to find all resources that have a tag with the key `Owner` and the value `TeamA` , specify `tag:Owner` for the filter name and `TeamA` for the filter value.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

