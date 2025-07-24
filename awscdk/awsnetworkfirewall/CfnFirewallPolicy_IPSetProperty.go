package awsnetworkfirewall


// A list of IP addresses and address ranges, in CIDR notation.
//
// This is part of a `RuleVariables` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPSetProperty := map[string][]*string{
//   	"definition": []*string{
//   		jsii.String("definition"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-ipset.html
//
type CfnFirewallPolicy_IPSetProperty struct {
	// The list of IP addresses and address ranges, in CIDR notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-ipset.html#cfn-networkfirewall-firewallpolicy-ipset-definition
	//
	Definition *[]*string `field:"optional" json:"definition" yaml:"definition"`
}

