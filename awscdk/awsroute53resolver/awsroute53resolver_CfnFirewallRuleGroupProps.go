package awsroute53resolver

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFirewallRuleGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFirewallRuleGroupProps := &cfnFirewallRuleGroupProps{
//   	firewallRules: []interface{}{
//   		&firewallRuleProperty{
//   			action: jsii.String("action"),
//   			firewallDomainListId: jsii.String("firewallDomainListId"),
//   			priority: jsii.Number(123),
//
//   			// the properties below are optional
//   			blockOverrideDnsType: jsii.String("blockOverrideDnsType"),
//   			blockOverrideDomain: jsii.String("blockOverrideDomain"),
//   			blockOverrideTtl: jsii.Number(123),
//   			blockResponse: jsii.String("blockResponse"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFirewallRuleGroupProps struct {
	// A list of the rules that you have defined.
	FirewallRules interface{} `field:"optional" json:"firewallRules" yaml:"firewallRules"`
	// The name of the rule group.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of the tag keys and values that you want to associate with the rule group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

