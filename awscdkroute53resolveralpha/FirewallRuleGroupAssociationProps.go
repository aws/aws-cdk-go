// The CDK Construct Library for AWS::Route53Resolver
package awscdkroute53resolveralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for a Firewall Rule Group Association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import route53resolver_alpha "github.com/aws/aws-cdk-go/awscdkroute53resolveralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var firewallRuleGroup firewallRuleGroup
//   var vpc vpc
//
//   firewallRuleGroupAssociationProps := &FirewallRuleGroupAssociationProps{
//   	FirewallRuleGroup: firewallRuleGroup,
//   	Priority: jsii.Number(123),
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	MutationProtection: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   }
//
// Experimental.
type FirewallRuleGroupAssociationProps struct {
	// The setting that determines the processing order of the rule group among the rule groups that are associated with a single VPC.
	//
	// DNS Firewall filters VPC
	// traffic starting from rule group with the lowest numeric priority setting.
	//
	// This value must be greater than 100 and less than 9,000.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The VPC that to associate with the rule group.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.
	// Experimental.
	MutationProtection *bool `field:"optional" json:"mutationProtection" yaml:"mutationProtection"`
	// The name of the association.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The firewall rule group which must be associated.
	// Experimental.
	FirewallRuleGroup IFirewallRuleGroup `field:"required" json:"firewallRuleGroup" yaml:"firewallRuleGroup"`
}

