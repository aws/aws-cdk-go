package awsnetworkfirewall


// A single stateless rule.
//
// This is used in `RuleGroup.StatelessRulesAndCustomActions` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statelessRuleProperty := &statelessRuleProperty{
//   	priority: jsii.Number(123),
//   	ruleDefinition: &ruleDefinitionProperty{
//   		actions: []*string{
//   			jsii.String("actions"),
//   		},
//   		matchAttributes: &matchAttributesProperty{
//   			destinationPorts: []interface{}{
//   				&portRangeProperty{
//   					fromPort: jsii.Number(123),
//   					toPort: jsii.Number(123),
//   				},
//   			},
//   			destinations: []interface{}{
//   				&addressProperty{
//   					addressDefinition: jsii.String("addressDefinition"),
//   				},
//   			},
//   			protocols: []interface{}{
//   				jsii.Number(123),
//   			},
//   			sourcePorts: []interface{}{
//   				&portRangeProperty{
//   					fromPort: jsii.Number(123),
//   					toPort: jsii.Number(123),
//   				},
//   			},
//   			sources: []interface{}{
//   				&addressProperty{
//   					addressDefinition: jsii.String("addressDefinition"),
//   				},
//   			},
//   			tcpFlags: []interface{}{
//   				&tCPFlagFieldProperty{
//   					flags: []*string{
//   						jsii.String("flags"),
//   					},
//
//   					// the properties below are optional
//   					masks: []*string{
//   						jsii.String("masks"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_StatelessRuleProperty struct {
	// Indicates the order in which to run this rule relative to all of the rules that are defined for a stateless rule group.
	//
	// Network Firewall evaluates the rules in a rule group starting with the lowest priority setting. You must ensure that the priority settings are unique for the rule group.
	//
	// Each stateless rule group uses exactly one `StatelessRulesAndCustomActions` object, and each `StatelessRulesAndCustomActions` contains exactly one `StatelessRules` object. To ensure unique priority settings for your rule groups, set unique priorities for the stateless rules that you define inside any single `StatelessRules` object.
	//
	// You can change the priority settings of your rules at any time. To make it easier to insert rules later, number them so there's a wide range in between, for example use 100, 200, and so on.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Defines the stateless 5-tuple packet inspection criteria and the action to take on a packet that matches the criteria.
	RuleDefinition interface{} `field:"required" json:"ruleDefinition" yaml:"ruleDefinition"`
}

