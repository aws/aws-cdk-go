package awsneptune

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDBSubnetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBSubnetGroupProps := &CfnDBSubnetGroupProps{
//   	DbSubnetGroupDescription: jsii.String("dbSubnetGroupDescription"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	DbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbsubnetgroup.html
//
type CfnDBSubnetGroupProps struct {
	// Provides the description of the DB subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbsubnetgroup.html#cfn-neptune-dbsubnetgroup-dbsubnetgroupdescription
	//
	DbSubnetGroupDescription *string `field:"required" json:"dbSubnetGroupDescription" yaml:"dbSubnetGroupDescription"`
	// The Amazon EC2 subnet IDs for the DB subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbsubnetgroup.html#cfn-neptune-dbsubnetgroup-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The name of the DB subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbsubnetgroup.html#cfn-neptune-dbsubnetgroup-dbsubnetgroupname
	//
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// The tags that you want to attach to the DB subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbsubnetgroup.html#cfn-neptune-dbsubnetgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

