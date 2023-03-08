package awsnetworkfirewall

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRuleGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRuleGroupProps := &cfnRuleGroupProps{
//   	capacity: jsii.Number(123),
//   	ruleGroupName: jsii.String("ruleGroupName"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	ruleGroup: &ruleGroupProperty{
//   		rulesSource: &rulesSourceProperty{
//   			rulesSourceList: &rulesSourceListProperty{
//   				generatedRulesType: jsii.String("generatedRulesType"),
//   				targets: []*string{
//   					jsii.String("targets"),
//   				},
//   				targetTypes: []*string{
//   					jsii.String("targetTypes"),
//   				},
//   			},
//   			rulesString: jsii.String("rulesString"),
//   			statefulRules: []interface{}{
//   				&statefulRuleProperty{
//   					action: jsii.String("action"),
//   					header: &headerProperty{
//   						destination: jsii.String("destination"),
//   						destinationPort: jsii.String("destinationPort"),
//   						direction: jsii.String("direction"),
//   						protocol: jsii.String("protocol"),
//   						source: jsii.String("source"),
//   						sourcePort: jsii.String("sourcePort"),
//   					},
//   					ruleOptions: []interface{}{
//   						&ruleOptionProperty{
//   							keyword: jsii.String("keyword"),
//
//   							// the properties below are optional
//   							settings: []*string{
//   								jsii.String("settings"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			statelessRulesAndCustomActions: &statelessRulesAndCustomActionsProperty{
//   				statelessRules: []interface{}{
//   					&statelessRuleProperty{
//   						priority: jsii.Number(123),
//   						ruleDefinition: &ruleDefinitionProperty{
//   							actions: []*string{
//   								jsii.String("actions"),
//   							},
//   							matchAttributes: &matchAttributesProperty{
//   								destinationPorts: []interface{}{
//   									&portRangeProperty{
//   										fromPort: jsii.Number(123),
//   										toPort: jsii.Number(123),
//   									},
//   								},
//   								destinations: []interface{}{
//   									&addressProperty{
//   										addressDefinition: jsii.String("addressDefinition"),
//   									},
//   								},
//   								protocols: []interface{}{
//   									jsii.Number(123),
//   								},
//   								sourcePorts: []interface{}{
//   									&portRangeProperty{
//   										fromPort: jsii.Number(123),
//   										toPort: jsii.Number(123),
//   									},
//   								},
//   								sources: []interface{}{
//   									&addressProperty{
//   										addressDefinition: jsii.String("addressDefinition"),
//   									},
//   								},
//   								tcpFlags: []interface{}{
//   									&tCPFlagFieldProperty{
//   										flags: []*string{
//   											jsii.String("flags"),
//   										},
//
//   										// the properties below are optional
//   										masks: []*string{
//   											jsii.String("masks"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				customActions: []interface{}{
//   					&customActionProperty{
//   						actionDefinition: &actionDefinitionProperty{
//   							publishMetricAction: &publishMetricActionProperty{
//   								dimensions: []interface{}{
//   									&dimensionProperty{
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   						actionName: jsii.String("actionName"),
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		referenceSets: &referenceSetsProperty{
//   			ipSetReferences: map[string]interface{}{
//   				"ipSetReferencesKey": map[string]*string{
//   					"referenceArn": jsii.String("referenceArn"),
//   				},
//   			},
//   		},
//   		ruleVariables: &ruleVariablesProperty{
//   			ipSets: map[string]interface{}{
//   				"ipSetsKey": map[string][]*string{
//   					"definition": []*string{
//   						jsii.String("definition"),
//   					},
//   				},
//   			},
//   			portSets: map[string]interface{}{
//   				"portSetsKey": &PortSetProperty{
//   					"definition": []*string{
//   						jsii.String("definition"),
//   					},
//   				},
//   			},
//   		},
//   		statefulRuleOptions: &statefulRuleOptionsProperty{
//   			ruleOrder: jsii.String("ruleOrder"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRuleGroupProps struct {
	// The maximum operating resources that this rule group can use.
	//
	// You can't change a rule group's capacity setting after you create the rule group. When you update a rule group, you are limited to this capacity. When you reference a rule group from a firewall policy, Network Firewall reserves this capacity for the rule group.
	Capacity *float64 `field:"required" json:"capacity" yaml:"capacity"`
	// The descriptive name of the rule group.
	//
	// You can't change the name of a rule group after you create it.
	RuleGroupName *string `field:"required" json:"ruleGroupName" yaml:"ruleGroupName"`
	// Indicates whether the rule group is stateless or stateful.
	//
	// If the rule group is stateless, it contains
	// stateless rules. If it is stateful, it contains stateful rules.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description of the rule group.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An object that defines the rule group rules.
	RuleGroup interface{} `field:"optional" json:"ruleGroup" yaml:"ruleGroup"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

