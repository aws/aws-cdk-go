package previewawsnetworkfirewallmixins


// The setting that allows the policy owner to change the behavior of the rule group within a policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   statefulRuleGroupOverrideProperty := &StatefulRuleGroupOverrideProperty{
//   	Action: jsii.String("action"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statefulrulegroupoverride.html
//
type CfnFirewallPolicyPropsMixin_StatefulRuleGroupOverrideProperty struct {
	// The action that changes the rule group from `DROP` to `ALERT` .
	//
	// This only applies to managed rule groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statefulrulegroupoverride.html#cfn-networkfirewall-firewallpolicy-statefulrulegroupoverride-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
}

