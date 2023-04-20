package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
type CfnSecurityGroupProps struct {
	// A description for the security group.
	//
	// Constraints: Up to 255 characters in length
	//
	// Valid characters: a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&;{}!$*
	GroupDescription *string `field:"required" json:"groupDescription" yaml:"groupDescription"`
	// The name of the security group.
	//
	// Constraints: Up to 255 characters in length. Cannot start with `sg-` .
	//
	// Valid characters: a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&;{}!$*
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// [VPC only] The outbound rules associated with the security group.
	//
	// There is a short interruption during which you cannot connect to the security group.
	SecurityGroupEgress interface{} `field:"optional" json:"securityGroupEgress" yaml:"securityGroupEgress"`
	// The inbound rules associated with the security group.
	//
	// There is a short interruption during which you cannot connect to the security group.
	SecurityGroupIngress interface{} `field:"optional" json:"securityGroupIngress" yaml:"securityGroupIngress"`
	// Any tags assigned to the security group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// [VPC only] The ID of the VPC for the security group.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

