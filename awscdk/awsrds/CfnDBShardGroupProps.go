package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDBShardGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBShardGroupProps := &CfnDBShardGroupProps{
//   	DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	MaxAcu: jsii.Number(123),
//
//   	// the properties below are optional
//   	ComputeRedundancy: jsii.Number(123),
//   	DbShardGroupIdentifier: jsii.String("dbShardGroupIdentifier"),
//   	MinAcu: jsii.Number(123),
//   	PubliclyAccessible: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbshardgroup.html
//
type CfnDBShardGroupProps struct {
	// The name of the primary DB cluster for the DB shard group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbshardgroup.html#cfn-rds-dbshardgroup-dbclusteridentifier
	//
	DbClusterIdentifier *string `field:"required" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The maximum capacity of the DB shard group in Aurora capacity units (ACUs).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbshardgroup.html#cfn-rds-dbshardgroup-maxacu
	//
	MaxAcu *float64 `field:"required" json:"maxAcu" yaml:"maxAcu"`
	// Specifies whether to create standby DB shard groups for the DB shard group. Valid values are the following:.
	//
	// - 0 - Creates a DB shard group without a standby DB shard group. This is the default value.
	// - 1 - Creates a DB shard group with a standby DB shard group in a different Availability Zone (AZ).
	// - 2 - Creates a DB shard group with two standby DB shard groups in two different AZs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbshardgroup.html#cfn-rds-dbshardgroup-computeredundancy
	//
	ComputeRedundancy *float64 `field:"optional" json:"computeRedundancy" yaml:"computeRedundancy"`
	// The name of the DB shard group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbshardgroup.html#cfn-rds-dbshardgroup-dbshardgroupidentifier
	//
	DbShardGroupIdentifier *string `field:"optional" json:"dbShardGroupIdentifier" yaml:"dbShardGroupIdentifier"`
	// The minimum capacity of the DB shard group in Aurora capacity units (ACUs).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbshardgroup.html#cfn-rds-dbshardgroup-minacu
	//
	MinAcu *float64 `field:"optional" json:"minAcu" yaml:"minAcu"`
	// Specifies whether the DB shard group is publicly accessible.
	//
	// When the DB shard group is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from within the DB shard group's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB shard group's VPC. Access to the DB shard group is ultimately controlled by the security group it uses. That public access is not permitted if the security group assigned to the DB shard group doesn't permit it.
	//
	// When the DB shard group isn't publicly accessible, it is an internal DB shard group with a DNS name that resolves to a private IP address.
	//
	// Default: The default behavior varies depending on whether `DBSubnetGroupName` is specified.
	//
	// If `DBSubnetGroupName` isn't specified, and `PubliclyAccessible` isn't specified, the following applies:
	//
	// - If the default VPC in the target Region doesn’t have an internet gateway attached to it, the DB shard group is private.
	// - If the default VPC in the target Region has an internet gateway attached to it, the DB shard group is public.
	//
	// If `DBSubnetGroupName` is specified, and `PubliclyAccessible` isn't specified, the following applies:
	//
	// - If the subnets are part of a VPC that doesn’t have an internet gateway attached to it, the DB shard group is private.
	// - If the subnets are part of a VPC that has an internet gateway attached to it, the DB shard group is public.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbshardgroup.html#cfn-rds-dbshardgroup-publiclyaccessible
	//
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// An optional set of key-value pairs to associate arbitrary data of your choosing with the DB shard group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbshardgroup.html#cfn-rds-dbshardgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

