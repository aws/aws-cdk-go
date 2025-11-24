package mixinsawsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDBSecurityGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDBSecurityGroupMixinProps := &CfnDBSecurityGroupMixinProps{
//   	DbSecurityGroupIngress: []interface{}{
//   		&IngressProperty{
//   			Cidrip: jsii.String("cidrip"),
//   			Ec2SecurityGroupId: jsii.String("ec2SecurityGroupId"),
//   			Ec2SecurityGroupName: jsii.String("ec2SecurityGroupName"),
//   			Ec2SecurityGroupOwnerId: jsii.String("ec2SecurityGroupOwnerId"),
//   		},
//   	},
//   	Ec2VpcId: jsii.String("ec2VpcId"),
//   	GroupDescription: jsii.String("groupDescription"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroup.html
//
type CfnDBSecurityGroupMixinProps struct {
	// Ingress rules to be applied to the DB security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroup.html#cfn-rds-dbsecuritygroup-dbsecuritygroupingress
	//
	DbSecurityGroupIngress interface{} `field:"optional" json:"dbSecurityGroupIngress" yaml:"dbSecurityGroupIngress"`
	// The identifier of an Amazon virtual private cloud (VPC).
	//
	// This property indicates the VPC that this DB security group belongs to.
	//
	// > This property is included for backwards compatibility and is no longer recommended for providing security information to an RDS DB instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroup.html#cfn-rds-dbsecuritygroup-ec2vpcid
	//
	Ec2VpcId *string `field:"optional" json:"ec2VpcId" yaml:"ec2VpcId"`
	// Provides the description of the DB security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroup.html#cfn-rds-dbsecuritygroup-groupdescription
	//
	GroupDescription *string `field:"optional" json:"groupDescription" yaml:"groupDescription"`
	// Metadata assigned to an Amazon RDS resource consisting of a key-value pair.
	//
	// For more information, see [Tagging Amazon RDS resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide* or [Tagging Amazon Aurora and Amazon RDS resources](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Tagging.html) in the *Amazon Aurora User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroup.html#cfn-rds-dbsecuritygroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

