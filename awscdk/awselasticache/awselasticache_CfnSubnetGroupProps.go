package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnSubnetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubnetGroupProps := &cfnSubnetGroupProps{
//   	description: jsii.String("description"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	cacheSubnetGroupName: jsii.String("cacheSubnetGroupName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSubnetGroupProps struct {
	// The description for the cache subnet group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The EC2 subnet IDs for the cache subnet group.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The name for the cache subnet group. This value is stored as a lowercase string.
	//
	// Constraints: Must contain no more than 255 alphanumeric characters or hyphens.
	//
	// Example: `mysubnetgroup`.
	CacheSubnetGroupName *string `field:"optional" json:"cacheSubnetGroupName" yaml:"cacheSubnetGroupName"`
	// A tag that can be added to an ElastiCache subnet group.
	//
	// Tags are composed of a Key/Value pair. You can use tags to categorize and track all your subnet groups. A tag with a null Value is permitted.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

