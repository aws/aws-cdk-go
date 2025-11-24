package mixinsawsroute53resolver

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFirewallRuleGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallRuleGroupMixinProps := &CfnFirewallRuleGroupMixinProps{
//   	FirewallRules: []interface{}{
//   		&FirewallRuleProperty{
//   			Action: jsii.String("action"),
//   			BlockOverrideDnsType: jsii.String("blockOverrideDnsType"),
//   			BlockOverrideDomain: jsii.String("blockOverrideDomain"),
//   			BlockOverrideTtl: jsii.Number(123),
//   			BlockResponse: jsii.String("blockResponse"),
//   			ConfidenceThreshold: jsii.String("confidenceThreshold"),
//   			DnsThreatProtection: jsii.String("dnsThreatProtection"),
//   			FirewallDomainListId: jsii.String("firewallDomainListId"),
//   			FirewallDomainRedirectionAction: jsii.String("firewallDomainRedirectionAction"),
//   			FirewallThreatProtectionId: jsii.String("firewallThreatProtectionId"),
//   			Priority: jsii.Number(123),
//   			Qtype: jsii.String("qtype"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-firewallrulegroup.html
//
type CfnFirewallRuleGroupMixinProps struct {
	// A list of the rules that you have defined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-firewallrulegroup.html#cfn-route53resolver-firewallrulegroup-firewallrules
	//
	FirewallRules interface{} `field:"optional" json:"firewallRules" yaml:"firewallRules"`
	// The name of the rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-firewallrulegroup.html#cfn-route53resolver-firewallrulegroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of the tag keys and values that you want to associate with the rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-firewallrulegroup.html#cfn-route53resolver-firewallrulegroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

