package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnReplicationSubnetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationSubnetGroupProps := &CfnReplicationSubnetGroupProps{
//   	ReplicationSubnetGroupDescription: jsii.String("replicationSubnetGroupDescription"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	ReplicationSubnetGroupIdentifier: jsii.String("replicationSubnetGroupIdentifier"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html
//
type CfnReplicationSubnetGroupProps struct {
	// The description for the subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-replicationsubnetgroupdescription
	//
	ReplicationSubnetGroupDescription *string `field:"required" json:"replicationSubnetGroupDescription" yaml:"replicationSubnetGroupDescription"`
	// One or more subnet IDs to be assigned to the subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The identifier for the replication subnet group.
	//
	// If you don't specify a name, CloudFormation generates a unique ID and uses that ID for the identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-replicationsubnetgroupidentifier
	//
	ReplicationSubnetGroupIdentifier *string `field:"optional" json:"replicationSubnetGroupIdentifier" yaml:"replicationSubnetGroupIdentifier"`
	// One or more tags to be assigned to the subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

