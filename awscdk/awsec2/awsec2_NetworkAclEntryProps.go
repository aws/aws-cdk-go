package awsec2


// Properties to create NetworkAclEntry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var aclCidr aclCidr
//   var aclTraffic aclTraffic
//   var networkAcl networkAcl
//
//   networkAclEntryProps := &networkAclEntryProps{
//   	cidr: aclCidr,
//   	networkAcl: networkAcl,
//   	ruleNumber: jsii.Number(123),
//   	traffic: aclTraffic,
//
//   	// the properties below are optional
//   	direction: awscdk.Aws_ec2.trafficDirection_EGRESS,
//   	networkAclEntryName: jsii.String("networkAclEntryName"),
//   	ruleAction: awscdk.*Aws_ec2.action_ALLOW,
//   }
//
type NetworkAclEntryProps struct {
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
	Direction TrafficDirection `field:"optional" json:"direction" yaml:"direction"`
	// The name of the NetworkAclEntry.
	//
	// It is not recommended to use an explicit group name.
	NetworkAclEntryName *string `field:"optional" json:"networkAclEntryName" yaml:"networkAclEntryName"`
	// Whether to allow or deny traffic that matches the rule; valid values are "allow" or "deny".
	//
	// Any traffic that is not explicitly allowed is automatically denied in a custom
	// ACL, all traffic is automatically allowed in a default ACL.
	RuleAction Action `field:"optional" json:"ruleAction" yaml:"ruleAction"`
	// The network ACL this entry applies to.
	NetworkAcl INetworkAcl `field:"required" json:"networkAcl" yaml:"networkAcl"`
}

