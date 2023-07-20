package awsnetworkfirewall


// Contains variables that you can use to override default Suricata settings in your firewall policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyVariablesProperty := &PolicyVariablesProperty{
//   	RuleVariables: map[string]interface{}{
//   		"ruleVariablesKey": map[string][]*string{
//   			"definition": []*string{
//   				jsii.String("definition"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-policyvariables.html
//
type CfnFirewallPolicy_PolicyVariablesProperty struct {
	// The IPv4 or IPv6 addresses in CIDR notation to use for the Suricata `HOME_NET` variable.
	//
	// If your firewall uses an inspection VPC, you might want to override the `HOME_NET` variable with the CIDRs of your home networks. If you don't override `HOME_NET` with your own CIDRs, Network Firewall by default uses the CIDR of your inspection VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-policyvariables.html#cfn-networkfirewall-firewallpolicy-policyvariables-rulevariables
	//
	RuleVariables interface{} `field:"optional" json:"ruleVariables" yaml:"ruleVariables"`
}

