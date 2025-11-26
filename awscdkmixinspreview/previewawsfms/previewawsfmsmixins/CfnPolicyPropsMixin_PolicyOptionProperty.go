package previewawsfmsmixins


// Contains the settings to configure a network ACL policy, a AWS Network Firewall firewall policy deployment model, or a third-party firewall policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   policyOptionProperty := &PolicyOptionProperty{
//   	NetworkAclCommonPolicy: &NetworkAclCommonPolicyProperty{
//   		NetworkAclEntrySet: &NetworkAclEntrySetProperty{
//   			FirstEntries: []interface{}{
//   				&NetworkAclEntryProperty{
//   					CidrBlock: jsii.String("cidrBlock"),
//   					Egress: jsii.Boolean(false),
//   					IcmpTypeCode: &IcmpTypeCodeProperty{
//   						Code: jsii.Number(123),
//   						Type: jsii.Number(123),
//   					},
//   					Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   					PortRange: &PortRangeProperty{
//   						From: jsii.Number(123),
//   						To: jsii.Number(123),
//   					},
//   					Protocol: jsii.String("protocol"),
//   					RuleAction: jsii.String("ruleAction"),
//   				},
//   			},
//   			ForceRemediateForFirstEntries: jsii.Boolean(false),
//   			ForceRemediateForLastEntries: jsii.Boolean(false),
//   			LastEntries: []interface{}{
//   				&NetworkAclEntryProperty{
//   					CidrBlock: jsii.String("cidrBlock"),
//   					Egress: jsii.Boolean(false),
//   					IcmpTypeCode: &IcmpTypeCodeProperty{
//   						Code: jsii.Number(123),
//   						Type: jsii.Number(123),
//   					},
//   					Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   					PortRange: &PortRangeProperty{
//   						From: jsii.Number(123),
//   						To: jsii.Number(123),
//   					},
//   					Protocol: jsii.String("protocol"),
//   					RuleAction: jsii.String("ruleAction"),
//   				},
//   			},
//   		},
//   	},
//   	NetworkFirewallPolicy: &NetworkFirewallPolicyProperty{
//   		FirewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   	},
//   	ThirdPartyFirewallPolicy: &ThirdPartyFirewallPolicyProperty{
//   		FirewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policyoption.html
//
type CfnPolicyPropsMixin_PolicyOptionProperty struct {
	// Defines a Firewall Manager network ACL policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policyoption.html#cfn-fms-policy-policyoption-networkaclcommonpolicy
	//
	NetworkAclCommonPolicy interface{} `field:"optional" json:"networkAclCommonPolicy" yaml:"networkAclCommonPolicy"`
	// Defines the deployment model to use for the firewall policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policyoption.html#cfn-fms-policy-policyoption-networkfirewallpolicy
	//
	NetworkFirewallPolicy interface{} `field:"optional" json:"networkFirewallPolicy" yaml:"networkFirewallPolicy"`
	// Defines the policy options for a third-party firewall policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policyoption.html#cfn-fms-policy-policyoption-thirdpartyfirewallpolicy
	//
	ThirdPartyFirewallPolicy interface{} `field:"optional" json:"thirdPartyFirewallPolicy" yaml:"thirdPartyFirewallPolicy"`
}

