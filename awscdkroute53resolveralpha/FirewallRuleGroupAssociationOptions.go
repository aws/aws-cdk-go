package awscdkroute53resolveralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Options for a Firewall Rule Group Association.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var ruleGroup FirewallRuleGroup
//   var myVpc Vpc
//
//
//   ruleGroup.Associate(jsii.String("Association"), &FirewallRuleGroupAssociationOptions{
//   	Priority: jsii.Number(101),
//   	Vpc: myVpc,
//   })
//
// Experimental.
type FirewallRuleGroupAssociationOptions struct {
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
	//
	// Note that mutation protection also blocks CloudFormation from updating or
	// deleting the association, so leave it disabled for associations whose
	// lifecycle is managed by this stack.
	// Default: - mutation protection is disabled; the association can be modified or removed.
	//
	// Experimental.
	MutationProtection *bool `field:"optional" json:"mutationProtection" yaml:"mutationProtection"`
	// The name of the association.
	// Default: - a CloudFormation generated name.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

