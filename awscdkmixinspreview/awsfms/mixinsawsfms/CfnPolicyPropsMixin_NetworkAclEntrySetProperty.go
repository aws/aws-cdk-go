package mixinsawsfms


// The configuration of the first and last rules for the network ACL policy, and the remediation settings for each.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkAclEntrySetProperty := &NetworkAclEntrySetProperty{
//   	FirstEntries: []interface{}{
//   		&NetworkAclEntryProperty{
//   			CidrBlock: jsii.String("cidrBlock"),
//   			Egress: jsii.Boolean(false),
//   			IcmpTypeCode: &IcmpTypeCodeProperty{
//   				Code: jsii.Number(123),
//   				Type: jsii.Number(123),
//   			},
//   			Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   			PortRange: &PortRangeProperty{
//   				From: jsii.Number(123),
//   				To: jsii.Number(123),
//   			},
//   			Protocol: jsii.String("protocol"),
//   			RuleAction: jsii.String("ruleAction"),
//   		},
//   	},
//   	ForceRemediateForFirstEntries: jsii.Boolean(false),
//   	ForceRemediateForLastEntries: jsii.Boolean(false),
//   	LastEntries: []interface{}{
//   		&NetworkAclEntryProperty{
//   			CidrBlock: jsii.String("cidrBlock"),
//   			Egress: jsii.Boolean(false),
//   			IcmpTypeCode: &IcmpTypeCodeProperty{
//   				Code: jsii.Number(123),
//   				Type: jsii.Number(123),
//   			},
//   			Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   			PortRange: &PortRangeProperty{
//   				From: jsii.Number(123),
//   				To: jsii.Number(123),
//   			},
//   			Protocol: jsii.String("protocol"),
//   			RuleAction: jsii.String("ruleAction"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclentryset.html
//
type CfnPolicyPropsMixin_NetworkAclEntrySetProperty struct {
	// The rules that you want to run first in the Firewall Manager managed network ACLs.
	//
	// > Provide these in the order in which you want them to run. Firewall Manager will assign the specific rule numbers for you, in the network ACLs that it creates.
	//
	// You must specify at least one first entry or one last entry in any network ACL policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclentryset.html#cfn-fms-policy-networkaclentryset-firstentries
	//
	FirstEntries interface{} `field:"optional" json:"firstEntries" yaml:"firstEntries"`
	// Applies only when remediation is enabled for the policy as a whole.
	//
	// Firewall Manager uses this setting when it finds policy violations that involve conflicts between the custom entries and the policy entries.
	//
	// If forced remediation is disabled, Firewall Manager marks the network ACL as noncompliant and does not try to remediate. For more information about the remediation behavior, see [Remediation for managed network ACLs](https://docs.aws.amazon.com/waf/latest/developerguide/network-acl-policies.html#network-acls-remediation) in the *AWS Firewall Manager Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclentryset.html#cfn-fms-policy-networkaclentryset-forceremediateforfirstentries
	//
	ForceRemediateForFirstEntries interface{} `field:"optional" json:"forceRemediateForFirstEntries" yaml:"forceRemediateForFirstEntries"`
	// Applies only when remediation is enabled for the policy as a whole.
	//
	// Firewall Manager uses this setting when it finds policy violations that involve conflicts between the custom entries and the policy entries.
	//
	// If forced remediation is disabled, Firewall Manager marks the network ACL as noncompliant and does not try to remediate. For more information about the remediation behavior, see [Remediation for managed network ACLs](https://docs.aws.amazon.com/waf/latest/developerguide/network-acl-policies.html#network-acls-remediation) in the *AWS Firewall Manager Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclentryset.html#cfn-fms-policy-networkaclentryset-forceremediateforlastentries
	//
	ForceRemediateForLastEntries interface{} `field:"optional" json:"forceRemediateForLastEntries" yaml:"forceRemediateForLastEntries"`
	// The rules that you want to run last in the Firewall Manager managed network ACLs.
	//
	// > Provide these in the order in which you want them to run. Firewall Manager will assign the specific rule numbers for you, in the network ACLs that it creates.
	//
	// You must specify at least one first entry or one last entry in any network ACL policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclentryset.html#cfn-fms-policy-networkaclentryset-lastentries
	//
	LastEntries interface{} `field:"optional" json:"lastEntries" yaml:"lastEntries"`
}

