package awsnetworkfirewall


// Stateless inspection criteria.
//
// Each stateless rule group uses exactly one of these data types to define its stateless rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statelessRulesAndCustomActionsProperty := &statelessRulesAndCustomActionsProperty{
//   	statelessRules: []interface{}{
//   		&statelessRuleProperty{
//   			priority: jsii.Number(123),
//   			ruleDefinition: &ruleDefinitionProperty{
//   				actions: []*string{
//   					jsii.String("actions"),
//   				},
//   				matchAttributes: &matchAttributesProperty{
//   					destinationPorts: []interface{}{
//   						&portRangeProperty{
//   							fromPort: jsii.Number(123),
//   							toPort: jsii.Number(123),
//   						},
//   					},
//   					destinations: []interface{}{
//   						&addressProperty{
//   							addressDefinition: jsii.String("addressDefinition"),
//   						},
//   					},
//   					protocols: []interface{}{
//   						jsii.Number(123),
//   					},
//   					sourcePorts: []interface{}{
//   						&portRangeProperty{
//   							fromPort: jsii.Number(123),
//   							toPort: jsii.Number(123),
//   						},
//   					},
//   					sources: []interface{}{
//   						&addressProperty{
//   							addressDefinition: jsii.String("addressDefinition"),
//   						},
//   					},
//   					tcpFlags: []interface{}{
//   						&tCPFlagFieldProperty{
//   							flags: []*string{
//   								jsii.String("flags"),
//   							},
//
//   							// the properties below are optional
//   							masks: []*string{
//   								jsii.String("masks"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	customActions: []interface{}{
//   		&customActionProperty{
//   			actionDefinition: &actionDefinitionProperty{
//   				publishMetricAction: &publishMetricActionProperty{
//   					dimensions: []interface{}{
//   						&dimensionProperty{
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			actionName: jsii.String("actionName"),
//   		},
//   	},
//   }
//
type CfnRuleGroup_StatelessRulesAndCustomActionsProperty struct {
	// Defines the set of stateless rules for use in a stateless rule group.
	StatelessRules interface{} `field:"required" json:"statelessRules" yaml:"statelessRules"`
	// Defines an array of individual custom action definitions that are available for use by the stateless rules in this `StatelessRulesAndCustomActions` specification.
	//
	// You name each custom action that you define, and then you can use it by name in your stateless rule `RuleGroup.RuleDefinition` `Actions` specification.
	CustomActions interface{} `field:"optional" json:"customActions" yaml:"customActions"`
}

