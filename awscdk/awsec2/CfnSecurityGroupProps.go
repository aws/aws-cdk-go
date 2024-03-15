package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSecurityGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityGroupProps := &CfnSecurityGroupProps{
//   	GroupDescription: jsii.String("groupDescription"),
//
//   	// the properties below are optional
//   	GroupName: jsii.String("groupName"),
//   	SecurityGroupEgress: []interface{}{
//   		&EgressProperty{
//   			IpProtocol: jsii.String("ipProtocol"),
//
//   			// the properties below are optional
//   			CidrIp: jsii.String("cidrIp"),
//   			CidrIpv6: jsii.String("cidrIpv6"),
//   			Description: jsii.String("description"),
//   			DestinationPrefixListId: jsii.String("destinationPrefixListId"),
//   			DestinationSecurityGroupId: jsii.String("destinationSecurityGroupId"),
//   			FromPort: jsii.Number(123),
//   			SourceSecurityGroupId: jsii.String("sourceSecurityGroupId"),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	SecurityGroupIngress: []interface{}{
//   		&IngressProperty{
//   			IpProtocol: jsii.String("ipProtocol"),
//
//   			// the properties below are optional
//   			CidrIp: jsii.String("cidrIp"),
//   			CidrIpv6: jsii.String("cidrIpv6"),
//   			Description: jsii.String("description"),
//   			FromPort: jsii.Number(123),
//   			SourcePrefixListId: jsii.String("sourcePrefixListId"),
//   			SourceSecurityGroupId: jsii.String("sourceSecurityGroupId"),
//   			SourceSecurityGroupName: jsii.String("sourceSecurityGroupName"),
//   			SourceSecurityGroupOwnerId: jsii.String("sourceSecurityGroupOwnerId"),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroup.html
//
type CfnSecurityGroupProps struct {
	// A description for the security group.
	//
	// Constraints: Up to 255 characters in length
	//
	// Valid characters: a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&;{}!$*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroup.html#cfn-ec2-securitygroup-groupdescription
	//
	GroupDescription *string `field:"required" json:"groupDescription" yaml:"groupDescription"`
	// The name of the security group.
	//
	// Constraints: Up to 255 characters in length. Cannot start with `sg-` .
	//
	// Valid characters: a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&;{}!$*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroup.html#cfn-ec2-securitygroup-groupname
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The outbound rules associated with the security group.
	//
	// There is a short interruption during which you cannot connect to the security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroup.html#cfn-ec2-securitygroup-securitygroupegress
	//
	SecurityGroupEgress interface{} `field:"optional" json:"securityGroupEgress" yaml:"securityGroupEgress"`
	// The inbound rules associated with the security group.
	//
	// There is a short interruption during which you cannot connect to the security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroup.html#cfn-ec2-securitygroup-securitygroupingress
	//
	SecurityGroupIngress interface{} `field:"optional" json:"securityGroupIngress" yaml:"securityGroupIngress"`
	// Any tags assigned to the security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroup.html#cfn-ec2-securitygroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the VPC for the security group.
	//
	// If you do not specify a VPC, the default is to use the default VPC for the Region. If there's no specified VPC and no default VPC, security group creation fails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroup.html#cfn-ec2-securitygroup-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

