package awsec2


// Basic NetworkACL entry props.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var aclCidr aclCidr
//   var aclTraffic aclTraffic
//
//   commonNetworkAclEntryOptions := &CommonNetworkAclEntryOptions{
//   	Cidr: aclCidr,
//   	RuleNumber: jsii.Number(123),
//   	Traffic: aclTraffic,
//
//   	// the properties below are optional
//   	Direction: awscdk.Aws_ec2.TrafficDirection_EGRESS,
//   	NetworkAclEntryName: jsii.String("networkAclEntryName"),
//   	RuleAction: awscdk.*Aws_ec2.Action_ALLOW,
//   }
//
type CommonNetworkAclEntryOptions struct {
	// The CIDR range to allow or deny.
	Cidr AclCidr `field:"required" json:"cidr" yaml:"cidr"`
	// Rule number to assign to the entry, such as 100.
	//
	// ACL entries are processed in ascending order by rule number.
	// Entries can't use the same rule number unless one is an egress rule and the other is an ingress rule.
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
	// What kind of traffic this ACL rule applies to.
	Traffic AclTraffic `field:"required" json:"traffic" yaml:"traffic"`
	// Traffic direction, with respect to the subnet, this rule applies to.
	// Default: TrafficDirection.INGRESS
	//
	Direction TrafficDirection `field:"optional" json:"direction" yaml:"direction"`
	// The name of the NetworkAclEntry.
	//
	// It is not recommended to use an explicit group name.
	// Default: If you don't specify a NetworkAclName, AWS CloudFormation generates a
	// unique physical ID and uses that ID for the group name.
	//
	NetworkAclEntryName *string `field:"optional" json:"networkAclEntryName" yaml:"networkAclEntryName"`
	// Whether to allow or deny traffic that matches the rule; valid values are "allow" or "deny".
	//
	// Any traffic that is not explicitly allowed is automatically denied in a custom
	// ACL, all traffic is automatically allowed in a default ACL.
	// Default: ALLOW.
	//
	RuleAction Action `field:"optional" json:"ruleAction" yaml:"ruleAction"`
}

