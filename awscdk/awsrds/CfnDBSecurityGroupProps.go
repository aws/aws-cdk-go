package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDBSecurityGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBSecurityGroupProps := &CfnDBSecurityGroupProps{
//   	DbSecurityGroupIngress: []interface{}{
//   		&IngressProperty{
//   			Cidrip: jsii.String("cidrip"),
//   			Ec2SecurityGroupId: jsii.String("ec2SecurityGroupId"),
//   			Ec2SecurityGroupName: jsii.String("ec2SecurityGroupName"),
//   			Ec2SecurityGroupOwnerId: jsii.String("ec2SecurityGroupOwnerId"),
//   		},
//   	},
//   	GroupDescription: jsii.String("groupDescription"),
//
//   	// the properties below are optional
//   	Ec2VpcId: jsii.String("ec2VpcId"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroup.html
//
type CfnDBSecurityGroupProps struct {
	// Ingress rules to be applied to the DB security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroup.html#cfn-rds-dbsecuritygroup-dbsecuritygroupingress
	//
	DbSecurityGroupIngress interface{} `field:"required" json:"dbSecurityGroupIngress" yaml:"dbSecurityGroupIngress"`
	// Provides the description of the DB security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroup.html#cfn-rds-dbsecuritygroup-groupdescription
	//
	GroupDescription *string `field:"required" json:"groupDescription" yaml:"groupDescription"`
	// The identifier of an Amazon VPC. This property indicates the VPC that this DB security group belongs to.
	//
	// > The `EC2VpcId` property is for backward compatibility with older regions, and is no longer recommended for providing security information to an RDS DB instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroup.html#cfn-rds-dbsecuritygroup-ec2vpcid
	//
	Ec2VpcId *string `field:"optional" json:"ec2VpcId" yaml:"ec2VpcId"`
	// An optional array of key-value pairs to apply to this DB security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroup.html#cfn-rds-dbsecuritygroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

