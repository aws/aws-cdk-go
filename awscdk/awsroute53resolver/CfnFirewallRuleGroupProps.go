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
//   cfnFirewallRuleGroupProps := &CfnFirewallRuleGroupProps{
//   	FirewallRules: []interface{}{
//   		&FirewallRuleProperty{
//   			Action: jsii.String("action"),
//   			FirewallDomainListId: jsii.String("firewallDomainListId"),
//   			Priority: jsii.Number(123),
//
//   			// the properties below are optional
//   			BlockOverrideDnsType: jsii.String("blockOverrideDnsType"),
//   			BlockOverrideDomain: jsii.String("blockOverrideDomain"),
//   			BlockOverrideTtl: jsii.Number(123),
//   			BlockResponse: jsii.String("blockResponse"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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

