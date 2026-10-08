package awsnetworksecuritymanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRuleProps := &CfnRuleProps{
//   	RuleName: jsii.String("ruleName"),
//
//   	// the properties below are optional
//   	Configuration: jsii.String("configuration"),
//   	FirewallType: jsii.String("firewallType"),
//   	RuleDescription: jsii.String("ruleDescription"),
//   	RuleType: jsii.String("ruleType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-rule.html
//
type CfnRuleProps struct {
	// The name of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-rule.html#cfn-networksecuritymanager-rule-rulename
	//
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// The rule configuration as a JSON string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-rule.html#cfn-networksecuritymanager-rule-configuration
	//
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// The type of firewall for this rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-rule.html#cfn-networksecuritymanager-rule-firewalltype
	//
	FirewallType *string `field:"optional" json:"firewallType" yaml:"firewallType"`
	// A description of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-rule.html#cfn-networksecuritymanager-rule-ruledescription
	//
	RuleDescription *string `field:"optional" json:"ruleDescription" yaml:"ruleDescription"`
	// The type of rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-rule.html#cfn-networksecuritymanager-rule-ruletype
	//
	RuleType *string `field:"optional" json:"ruleType" yaml:"ruleType"`
	// The tags associated with the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-rule.html#cfn-networksecuritymanager-rule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

