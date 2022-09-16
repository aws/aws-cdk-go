package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDBSecurityGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBSecurityGroupProps := &cfnDBSecurityGroupProps{
//   	dbSecurityGroupIngress: []interface{}{
//   		&ingressProperty{
//   			cidrip: jsii.String("cidrip"),
//   			ec2SecurityGroupId: jsii.String("ec2SecurityGroupId"),
//   			ec2SecurityGroupName: jsii.String("ec2SecurityGroupName"),
//   			ec2SecurityGroupOwnerId: jsii.String("ec2SecurityGroupOwnerId"),
//   		},
//   	},
//   	groupDescription: jsii.String("groupDescription"),
//
//   	// the properties below are optional
//   	ec2VpcId: jsii.String("ec2VpcId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDBSecurityGroupProps struct {
	// Ingress rules to be applied to the DB security group.
	DbSecurityGroupIngress interface{} `field:"required" json:"dbSecurityGroupIngress" yaml:"dbSecurityGroupIngress"`
	// Provides the description of the DB security group.
	GroupDescription *string `field:"required" json:"groupDescription" yaml:"groupDescription"`
	// The identifier of an Amazon VPC. This property indicates the VPC that this DB security group belongs to.
	//
	// > The `EC2VpcId` property is for backward compatibility with older regions, and is no longer recommended for providing security information to an RDS DB instance.
	Ec2VpcId *string `field:"optional" json:"ec2VpcId" yaml:"ec2VpcId"`
	// Tags to assign to the DB security group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

