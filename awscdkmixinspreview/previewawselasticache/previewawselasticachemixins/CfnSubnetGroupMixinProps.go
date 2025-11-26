package previewawselasticachemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSubnetGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSubnetGroupMixinProps := &CfnSubnetGroupMixinProps{
//   	CacheSubnetGroupName: jsii.String("cacheSubnetGroupName"),
//   	Description: jsii.String("description"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
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
type CfnSubnetGroupMixinProps struct {
	// The name for the cache subnet group. This value is stored as a lowercase string.
	//
	// Constraints: Must contain no more than 255 alphanumeric characters or hyphens.
	//
	// Example: `mysubnetgroup`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-subnetgroup.html#cfn-elasticache-subnetgroup-cachesubnetgroupname
	//
	CacheSubnetGroupName *string `field:"optional" json:"cacheSubnetGroupName" yaml:"cacheSubnetGroupName"`
	// The description for the cache subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-subnetgroup.html#cfn-elasticache-subnetgroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The EC2 subnet IDs for the cache subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-subnetgroup.html#cfn-elasticache-subnetgroup-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// A tag that can be added to an ElastiCache subnet group.
	//
	// Tags are composed of a Key/Value pair. You can use tags to categorize and track all your subnet groups. A tag with a null Value is permitted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-subnetgroup.html#cfn-elasticache-subnetgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

