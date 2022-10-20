package awsroute53resolver

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFirewallRuleGroupAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFirewallRuleGroupAssociationProps := &cfnFirewallRuleGroupAssociationProps{
//   	firewallRuleGroupId: jsii.String("firewallRuleGroupId"),
//   	priority: jsii.Number(123),
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	mutationProtection: jsii.String("mutationProtection"),
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFirewallRuleGroupAssociationProps struct {
	// The unique identifier of the firewall rule group.
	FirewallRuleGroupId *string `field:"required" json:"firewallRuleGroupId" yaml:"firewallRuleGroupId"`
	// The setting that determines the processing order of the rule group among the rule groups that are associated with a single VPC.
	//
	// DNS Firewall filters VPC traffic starting from rule group with the lowest numeric priority setting.
	//
	// You must specify a unique priority for each rule group that you associate with a single VPC. To make it easier to insert rule groups later, leave space between the numbers, for example, use 101, 200, and so on. You can change the priority setting for a rule group association after you create it.
	//
	// The allowed values for `Priority` are between 100 and 9900 (excluding 100 and 9900).
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The unique identifier of the VPC that is associated with the rule group.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.
	MutationProtection *string `field:"optional" json:"mutationProtection" yaml:"mutationProtection"`
	// The name of the association.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of the tag keys and values that you want to associate with the rule group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

