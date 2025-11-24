package mixinsawsnetworkfirewall

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRuleGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuleGroupMixinProps := &CfnRuleGroupMixinProps{
//   	Capacity: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	RuleGroup: &RuleGroupProperty{
//   		ReferenceSets: &ReferenceSetsProperty{
//   			IpSetReferences: map[string]interface{}{
//   				"ipSetReferencesKey": map[string]*string{
//   					"referenceArn": jsii.String("referenceArn"),
//   				},
//   			},
//   		},
//   		RulesSource: &RulesSourceProperty{
//   			RulesSourceList: &RulesSourceListProperty{
//   				GeneratedRulesType: jsii.String("generatedRulesType"),
//   				Targets: []*string{
//   					jsii.String("targets"),
//   				},
//   				TargetTypes: []*string{
//   					jsii.String("targetTypes"),
//   				},
//   			},
//   			RulesString: jsii.String("rulesString"),
//   			StatefulRules: []interface{}{
//   				&StatefulRuleProperty{
//   					Action: jsii.String("action"),
//   					Header: &HeaderProperty{
//   						Destination: jsii.String("destination"),
//   						DestinationPort: jsii.String("destinationPort"),
//   						Direction: jsii.String("direction"),
//   						Protocol: jsii.String("protocol"),
//   						Source: jsii.String("source"),
//   						SourcePort: jsii.String("sourcePort"),
//   					},
//   					RuleOptions: []interface{}{
//   						&RuleOptionProperty{
//   							Keyword: jsii.String("keyword"),
//   							Settings: []*string{
//   								jsii.String("settings"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			StatelessRulesAndCustomActions: &StatelessRulesAndCustomActionsProperty{
//   				CustomActions: []interface{}{
//   					&CustomActionProperty{
//   						ActionDefinition: &ActionDefinitionProperty{
//   							PublishMetricAction: &PublishMetricActionProperty{
//   								Dimensions: []interface{}{
//   									&DimensionProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   						ActionName: jsii.String("actionName"),
//   					},
//   				},
//   				StatelessRules: []interface{}{
//   					&StatelessRuleProperty{
//   						Priority: jsii.Number(123),
//   						RuleDefinition: &RuleDefinitionProperty{
//   							Actions: []*string{
//   								jsii.String("actions"),
//   							},
//   							MatchAttributes: &MatchAttributesProperty{
//   								DestinationPorts: []interface{}{
//   									&PortRangeProperty{
//   										FromPort: jsii.Number(123),
//   										ToPort: jsii.Number(123),
//   									},
//   								},
//   								Destinations: []interface{}{
//   									&AddressProperty{
//   										AddressDefinition: jsii.String("addressDefinition"),
//   									},
//   								},
//   								Protocols: []interface{}{
//   									jsii.Number(123),
//   								},
//   								SourcePorts: []interface{}{
//   									&PortRangeProperty{
//   										FromPort: jsii.Number(123),
//   										ToPort: jsii.Number(123),
//   									},
//   								},
//   								Sources: []interface{}{
//   									&AddressProperty{
//   										AddressDefinition: jsii.String("addressDefinition"),
//   									},
//   								},
//   								TcpFlags: []interface{}{
//   									&TCPFlagFieldProperty{
//   										Flags: []*string{
//   											jsii.String("flags"),
//   										},
//   										Masks: []*string{
//   											jsii.String("masks"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		RuleVariables: &RuleVariablesProperty{
//   			IpSets: map[string]interface{}{
//   				"ipSetsKey": map[string][]*string{
//   					"definition": []*string{
//   						jsii.String("definition"),
//   					},
//   				},
//   			},
//   			PortSets: map[string]interface{}{
//   				"portSetsKey": &PortSetProperty{
//   					"definition": []*string{
//   						jsii.String("definition"),
//   					},
//   				},
//   			},
//   		},
//   		StatefulRuleOptions: &StatefulRuleOptionsProperty{
//   			RuleOrder: jsii.String("ruleOrder"),
//   		},
//   	},
//   	RuleGroupName: jsii.String("ruleGroupName"),
//   	SummaryConfiguration: &SummaryConfigurationProperty{
//   		RuleOptions: []*string{
//   			jsii.String("ruleOptions"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-rulegroup.html
//
type CfnRuleGroupMixinProps struct {
	// The maximum operating resources that this rule group can use.
	//
	// You can't change a rule group's capacity setting after you create the rule group. When you update a rule group, you are limited to this capacity. When you reference a rule group from a firewall policy, Network Firewall reserves this capacity for the rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-rulegroup.html#cfn-networkfirewall-rulegroup-capacity
	//
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// A description of the rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-rulegroup.html#cfn-networkfirewall-rulegroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An object that defines the rule group rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-rulegroup.html#cfn-networkfirewall-rulegroup-rulegroup
	//
	RuleGroup interface{} `field:"optional" json:"ruleGroup" yaml:"ruleGroup"`
	// The descriptive name of the rule group.
	//
	// You can't change the name of a rule group after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-rulegroup.html#cfn-networkfirewall-rulegroup-rulegroupname
	//
	RuleGroupName *string `field:"optional" json:"ruleGroupName" yaml:"ruleGroupName"`
	// A complex type containing the currently selected rule option fields that will be displayed for rule summarization returned by `DescribeRuleGroupSummary` .
	//
	// - The `RuleOptions` specified in `SummaryConfiguration`
	// - Rule metadata organization preferences.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-rulegroup.html#cfn-networkfirewall-rulegroup-summaryconfiguration
	//
	SummaryConfiguration interface{} `field:"optional" json:"summaryConfiguration" yaml:"summaryConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-rulegroup.html#cfn-networkfirewall-rulegroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Indicates whether the rule group is stateless or stateful.
	//
	// If the rule group is stateless, it contains
	// stateless rules. If it is stateful, it contains stateful rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-rulegroup.html#cfn-networkfirewall-rulegroup-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

