package mixinsawsnetworkfirewall


// Stateful inspection criteria for a domain list rule group.
//
// For HTTPS traffic, domain filtering is SNI-based. It uses the server name indicator extension of the TLS handshake.
//
// By default, Network Firewall domain list inspection only includes traffic coming from the VPC where you deploy the firewall. To inspect traffic from IP addresses outside of the deployment VPC, you set the `HOME_NET` rule variable to include the CIDR range of the deployment VPC plus the other CIDR ranges. For more information, see `RuleVariables` in this guide and [Stateful domain list rule groups in AWS Network Firewall](https://docs.aws.amazon.com/network-firewall/latest/developerguide/stateful-rule-groups-domain-names.html) in the *Network Firewall Developer Guide*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rulesSourceListProperty := &RulesSourceListProperty{
//   	GeneratedRulesType: jsii.String("generatedRulesType"),
//   	Targets: []*string{
//   		jsii.String("targets"),
//   	},
//   	TargetTypes: []*string{
//   		jsii.String("targetTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulessourcelist.html
//
type CfnRuleGroupPropsMixin_RulesSourceListProperty struct {
	// Whether you want to apply allow, reject, alert, or drop behavior to the domains in your target list.
	//
	// > When logging is enabled and you choose Alert, traffic that matches the domain specifications generates an alert in the firewall's logs. Then, traffic either passes, is rejected, or drops based on other rules in the firewall policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulessourcelist.html#cfn-networkfirewall-rulegroup-rulessourcelist-generatedrulestype
	//
	GeneratedRulesType *string `field:"optional" json:"generatedRulesType" yaml:"generatedRulesType"`
	// The domains that you want to inspect for in your traffic flows. Valid domain specifications are the following:.
	//
	// - Explicit names. For example, `abc.example.com` matches only the domain `abc.example.com` .
	// - Names that use a domain wildcard, which you indicate with an initial ' `.` '. For example, `.example.com` matches `example.com` and matches all subdomains of `example.com` , such as `abc.example.com` and `www.example.com` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulessourcelist.html#cfn-networkfirewall-rulegroup-rulessourcelist-targets
	//
	Targets *[]*string `field:"optional" json:"targets" yaml:"targets"`
	// The types of targets to inspect for.
	//
	// Valid values are `TLS_SNI` and `HTTP_HOST` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulessourcelist.html#cfn-networkfirewall-rulegroup-rulessourcelist-targettypes
	//
	TargetTypes *[]*string `field:"optional" json:"targetTypes" yaml:"targetTypes"`
}

