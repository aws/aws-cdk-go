package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnIPAMResourceDiscovery`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPAMResourceDiscoveryProps := &cfnIPAMResourceDiscoveryProps{
//   	description: jsii.String("description"),
//   	operatingRegions: []interface{}{
//   		&ipamOperatingRegionProperty{
//   			regionName: jsii.String("regionName"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnIPAMResourceDiscoveryProps struct {
	// The resource discovery description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The operating Regions for the resource discovery.
	//
	// Operating Regions are AWS Regions where the IPAM is allowed to manage IP address CIDRs. IPAM only discovers and monitors resources in the AWS Regions you select as operating Regions.
	OperatingRegions interface{} `field:"optional" json:"operatingRegions" yaml:"operatingRegions"`
	// A tag is a label that you assign to an AWS resource.
	//
	// Each tag consists of a key and an optional value. You can use tags to search and filter your resources or track your AWS costs.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

