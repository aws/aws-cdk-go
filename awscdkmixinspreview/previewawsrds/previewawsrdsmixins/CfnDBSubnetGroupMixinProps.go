package previewawsrdsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDBSubnetGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDBSubnetGroupMixinProps := &CfnDBSubnetGroupMixinProps{
//   	DbSubnetGroupDescription: jsii.String("dbSubnetGroupDescription"),
//   	DbSubnetGroupName: jsii.String("dbSubnetGroupName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnetgroup.html
//
type CfnDBSubnetGroupMixinProps struct {
	// The description for the DB subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnetgroup.html#cfn-rds-dbsubnetgroup-dbsubnetgroupdescription
	//
	DbSubnetGroupDescription *string `field:"optional" json:"dbSubnetGroupDescription" yaml:"dbSubnetGroupDescription"`
	// The name for the DB subnet group. This value is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain no more than 255 letters, numbers, periods, underscores, spaces, or hyphens.
	// - Must not be default.
	// - First character must be a letter.
	//
	// Example: `mydbsubnetgroup`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnetgroup.html#cfn-rds-dbsubnetgroup-dbsubnetgroupname
	//
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// The EC2 Subnet IDs for the DB subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnetgroup.html#cfn-rds-dbsubnetgroup-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Tags to assign to the DB subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnetgroup.html#cfn-rds-dbsubnetgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

