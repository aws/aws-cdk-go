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
//   cfnSecurityGroupProps := &cfnSecurityGroupProps{
//   	groupDescription: jsii.String("groupDescription"),
//
//   	// the properties below are optional
//   	groupName: jsii.String("groupName"),
//   	securityGroupEgress: []interface{}{
//   		&egressProperty{
//   			ipProtocol: jsii.String("ipProtocol"),
//
//   			// the properties below are optional
//   			cidrIp: jsii.String("cidrIp"),
//   			cidrIpv6: jsii.String("cidrIpv6"),
//   			description: jsii.String("description"),
//   			destinationPrefixListId: jsii.String("destinationPrefixListId"),
//   			destinationSecurityGroupId: jsii.String("destinationSecurityGroupId"),
//   			fromPort: jsii.Number(123),
//   			toPort: jsii.Number(123),
//   		},
//   	},
//   	securityGroupIngress: []interface{}{
//   		&ingressProperty{
//   			ipProtocol: jsii.String("ipProtocol"),
//
//   			// the properties below are optional
//   			cidrIp: jsii.String("cidrIp"),
//   			cidrIpv6: jsii.String("cidrIpv6"),
//   			description: jsii.String("description"),
//   			fromPort: jsii.Number(123),
//   			sourcePrefixListId: jsii.String("sourcePrefixListId"),
//   			sourceSecurityGroupId: jsii.String("sourceSecurityGroupId"),
//   			sourceSecurityGroupName: jsii.String("sourceSecurityGroupName"),
//   			sourceSecurityGroupOwnerId: jsii.String("sourceSecurityGroupOwnerId"),
//   			toPort: jsii.Number(123),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnSecurityGroupProps struct {
	// A description for the security group. This is informational only.
	//
	// Constraints: Up to 255 characters in length
	//
	// Constraints for EC2-Classic: ASCII characters
	//
	// Constraints for EC2-VPC: a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&;{}!$*
	GroupDescription *string `field:"required" json:"groupDescription" yaml:"groupDescription"`
	// The name of the security group.
	//
	// Constraints: Up to 255 characters in length. Cannot start with `sg-` .
	//
	// Constraints for EC2-Classic: ASCII characters
	//
	// Constraints for EC2-VPC: a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&;{}!$*
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

