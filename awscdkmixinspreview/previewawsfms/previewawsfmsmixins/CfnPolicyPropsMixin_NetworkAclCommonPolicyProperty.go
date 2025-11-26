package previewawsfmsmixins


// Defines a Firewall Manager network ACL policy.
//
// This is used in the `PolicyOption` of a `SecurityServicePolicyData` for a `Policy` , when the `SecurityServicePolicyData` type is set to `NETWORK_ACL_COMMON` .
//
// For information about network ACLs, see [Control traffic to subnets using network ACLs](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-network-acls.html) in the *Amazon Virtual Private Cloud User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkAclCommonPolicyProperty := &NetworkAclCommonPolicyProperty{
//   	NetworkAclEntrySet: &NetworkAclEntrySetProperty{
//   		FirstEntries: []interface{}{
//   			&NetworkAclEntryProperty{
//   				CidrBlock: jsii.String("cidrBlock"),
//   				Egress: jsii.Boolean(false),
//   				IcmpTypeCode: &IcmpTypeCodeProperty{
//   					Code: jsii.Number(123),
//   					Type: jsii.Number(123),
//   				},
//   				Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   				PortRange: &PortRangeProperty{
//   					From: jsii.Number(123),
//   					To: jsii.Number(123),
//   				},
//   				Protocol: jsii.String("protocol"),
//   				RuleAction: jsii.String("ruleAction"),
//   			},
//   		},
//   		ForceRemediateForFirstEntries: jsii.Boolean(false),
//   		ForceRemediateForLastEntries: jsii.Boolean(false),
//   		LastEntries: []interface{}{
//   			&NetworkAclEntryProperty{
//   				CidrBlock: jsii.String("cidrBlock"),
//   				Egress: jsii.Boolean(false),
//   				IcmpTypeCode: &IcmpTypeCodeProperty{
//   					Code: jsii.Number(123),
//   					Type: jsii.Number(123),
//   				},
//   				Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   				PortRange: &PortRangeProperty{
//   					From: jsii.Number(123),
//   					To: jsii.Number(123),
//   				},
//   				Protocol: jsii.String("protocol"),
//   				RuleAction: jsii.String("ruleAction"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclcommonpolicy.html
//
type CfnPolicyPropsMixin_NetworkAclCommonPolicyProperty struct {
	// The definition of the first and last rules for the network ACL policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclcommonpolicy.html#cfn-fms-policy-networkaclcommonpolicy-networkaclentryset
	//
	NetworkAclEntrySet interface{} `field:"optional" json:"networkAclEntrySet" yaml:"networkAclEntrySet"`
}

