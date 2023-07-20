package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIPAMResourceDiscovery`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPAMResourceDiscoveryProps := &CfnIPAMResourceDiscoveryProps{
//   	Description: jsii.String("description"),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamresourcediscovery.html
//
type CfnIPAMResourceDiscoveryProps struct {
	// The resource discovery description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamresourcediscovery.html#cfn-ec2-ipamresourcediscovery-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The operating Regions for the resource discovery.
	//
	// Operating Regions are AWS Regions where the IPAM is allowed to manage IP address CIDRs. IPAM only discovers and monitors resources in the AWS Regions you select as operating Regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamresourcediscovery.html#cfn-ec2-ipamresourcediscovery-operatingregions
	//
	OperatingRegions interface{} `field:"optional" json:"operatingRegions" yaml:"operatingRegions"`
	// A tag is a label that you assign to an AWS resource.
	//
	// Each tag consists of a key and an optional value. You can use tags to search and filter your resources or track your AWS costs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamresourcediscovery.html#cfn-ec2-ipamresourcediscovery-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

