package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSubnetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubnetGroupProps := &CfnSubnetGroupProps{
//   	Description: jsii.String("description"),
//   	SubnetIds: []interface{}{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	CacheSubnetGroupName: jsii.String("cacheSubnetGroupName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-subnetgroup.html
//
type CfnSubnetGroupProps struct {
	// The description for the cache subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-subnetgroup.html#cfn-elasticache-subnetgroup-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The EC2 subnet IDs for the cache subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-subnetgroup.html#cfn-elasticache-subnetgroup-subnetids
	//
	SubnetIds *[]interface{} `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The name for the cache subnet group. This value is stored as a lowercase string.
	//
	// Constraints: Must contain no more than 255 alphanumeric characters or hyphens.
	//
	// Example: `mysubnetgroup`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-subnetgroup.html#cfn-elasticache-subnetgroup-cachesubnetgroupname
	//
	CacheSubnetGroupName *string `field:"optional" json:"cacheSubnetGroupName" yaml:"cacheSubnetGroupName"`
	// A tag that can be added to an ElastiCache subnet group.
	//
	// Tags are composed of a Key/Value pair. You can use tags to categorize and track all your subnet groups. A tag with a null Value is permitted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-subnetgroup.html#cfn-elasticache-subnetgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

