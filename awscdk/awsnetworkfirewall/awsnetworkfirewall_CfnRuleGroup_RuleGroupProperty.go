package awsnetworkfirewall


// The object that defines the rules in a rule group.
//
// AWS Network Firewall uses a rule group to inspect and control network traffic. You define stateless rule groups to inspect individual packets and you define stateful rule groups to inspect packets in the context of their traffic flow.
//
// To use a rule group, you include it by reference in an Network Firewall firewall policy, then you use the policy in a firewall. You can reference a rule group from more than one firewall policy, and you can use a firewall policy in more than one firewall.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleGroupProperty := &ruleGroupProperty{
//   	rulesSource: &rulesSourceProperty{
//   		rulesSourceList: &rulesSourceListProperty{
//   			generatedRulesType: jsii.String("generatedRulesType"),
//   			targets: []*string{
//   				jsii.String("targets"),
//   			},
//   			targetTypes: []*string{
//   				jsii.String("targetTypes"),
//   			},
//   		},
//   		rulesString: jsii.String("rulesString"),
//   		statefulRules: []interface{}{
//   			&statefulRuleProperty{
//   				action: jsii.String("action"),
//   				header: &headerProperty{
//   					destination: jsii.String("destination"),
//   					destinationPort: jsii.String("destinationPort"),
//   					direction: jsii.String("direction"),
//   					protocol: jsii.String("protocol"),
//   					source: jsii.String("source"),
//   					sourcePort: jsii.String("sourcePort"),
//   				},
//   				ruleOptions: []interface{}{
//   					&ruleOptionProperty{
//   						keyword: jsii.String("keyword"),
//
//   						// the properties below are optional
//   						settings: []*string{
//   							jsii.String("settings"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		statelessRulesAndCustomActions: &statelessRulesAndCustomActionsProperty{
//   			statelessRules: []interface{}{
//   				&statelessRuleProperty{
//   					priority: jsii.Number(123),
//   					ruleDefinition: &ruleDefinitionProperty{
//   						actions: []*string{
//   							jsii.String("actions"),
//   						},
//   						matchAttributes: &matchAttributesProperty{
//   							destinationPorts: []interface{}{
//   								&portRangeProperty{
//   									fromPort: jsii.Number(123),
//   									toPort: jsii.Number(123),
//   								},
//   							},
//   							destinations: []interface{}{
//   								&addressProperty{
//   									addressDefinition: jsii.String("addressDefinition"),
//   								},
//   							},
//   							protocols: []interface{}{
//   								jsii.Number(123),
//   							},
//   							sourcePorts: []interface{}{
//   								&portRangeProperty{
//   									fromPort: jsii.Number(123),
//   									toPort: jsii.Number(123),
//   								},
//   							},
//   							sources: []interface{}{
//   								&addressProperty{
//   									addressDefinition: jsii.String("addressDefinition"),
//   								},
//   							},
//   							tcpFlags: []interface{}{
//   								&tCPFlagFieldProperty{
//   									flags: []*string{
//   										jsii.String("flags"),
//   									},
//
//   									// the properties below are optional
//   									masks: []*string{
//   										jsii.String("masks"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			customActions: []interface{}{
//   				&customActionProperty{
//   					actionDefinition: &actionDefinitionProperty{
//   						publishMetricAction: &publishMetricActionProperty{
//   							dimensions: []interface{}{
//   								&dimensionProperty{
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					actionName: jsii.String("actionName"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	referenceSets: &referenceSetsProperty{
//   		ipSetReferences: map[string]interface{}{
//   			"ipSetReferencesKey": map[string]*string{
//   				"referenceArn": jsii.String("referenceArn"),
//   			},
//   		},
//   	},
//   	ruleVariables: &ruleVariablesProperty{
//   		ipSets: map[string]interface{}{
//   			"ipSetsKey": map[string][]*string{
//   				"definition": []*string{
//   					jsii.String("definition"),
//   				},
//   			},
//   		},
//   		portSets: map[string]interface{}{
//   			"portSetsKey": &PortSetProperty{
//   				"definition": []*string{
//   					jsii.String("definition"),
//   				},
//   			},
//   		},
//   	},
//   	statefulRuleOptions: &statefulRuleOptionsProperty{
//   		ruleOrder: jsii.String("ruleOrder"),
//   	},
//   }
//
type CfnRuleGroup_RuleGroupProperty struct {
	// The stateful rules or stateless rules for the rule group.
	RulesSource interface{} `field:"required" json:"rulesSource" yaml:"rulesSource"`
	// `CfnRuleGroup.RuleGroupProperty.ReferenceSets`.
	ReferenceSets interface{} `field:"optional" json:"referenceSets" yaml:"referenceSets"`
	// Settings that are available for use in the rules in the rule group.
	//
	// You can only use these for stateful rule groups.
	RuleVariables interface{} `field:"optional" json:"ruleVariables" yaml:"ruleVariables"`
	// Additional options governing how Network Firewall handles stateful rules.
	//
	// The policies where you use your stateful rule group must have stateful rule options settings that are compatible with these settings.
	StatefulRuleOptions interface{} `field:"optional" json:"statefulRuleOptions" yaml:"statefulRuleOptions"`
}

