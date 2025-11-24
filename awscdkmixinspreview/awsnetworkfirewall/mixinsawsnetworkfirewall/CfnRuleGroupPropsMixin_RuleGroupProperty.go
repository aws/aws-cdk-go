package mixinsawsnetworkfirewall


// The object that defines the rules in a rule group.
//
// AWS Network Firewall uses a rule group to inspect and control network traffic. You define stateless rule groups to inspect individual packets and you define stateful rule groups to inspect packets in the context of their traffic flow.
//
// To use a rule group, you include it by reference in an Network Firewall firewall policy, then you use the policy in a firewall. You can reference a rule group from more than one firewall policy, and you can use a firewall policy in more than one firewall.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleGroupProperty := &RuleGroupProperty{
//   	ReferenceSets: &ReferenceSetsProperty{
//   		IpSetReferences: map[string]interface{}{
//   			"ipSetReferencesKey": map[string]*string{
//   				"referenceArn": jsii.String("referenceArn"),
//   			},
//   		},
//   	},
//   	RulesSource: &RulesSourceProperty{
//   		RulesSourceList: &RulesSourceListProperty{
//   			GeneratedRulesType: jsii.String("generatedRulesType"),
//   			Targets: []*string{
//   				jsii.String("targets"),
//   			},
//   			TargetTypes: []*string{
//   				jsii.String("targetTypes"),
//   			},
//   		},
//   		RulesString: jsii.String("rulesString"),
//   		StatefulRules: []interface{}{
//   			&StatefulRuleProperty{
//   				Action: jsii.String("action"),
//   				Header: &HeaderProperty{
//   					Destination: jsii.String("destination"),
//   					DestinationPort: jsii.String("destinationPort"),
//   					Direction: jsii.String("direction"),
//   					Protocol: jsii.String("protocol"),
//   					Source: jsii.String("source"),
//   					SourcePort: jsii.String("sourcePort"),
//   				},
//   				RuleOptions: []interface{}{
//   					&RuleOptionProperty{
//   						Keyword: jsii.String("keyword"),
//   						Settings: []*string{
//   							jsii.String("settings"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		StatelessRulesAndCustomActions: &StatelessRulesAndCustomActionsProperty{
//   			CustomActions: []interface{}{
//   				&CustomActionProperty{
//   					ActionDefinition: &ActionDefinitionProperty{
//   						PublishMetricAction: &PublishMetricActionProperty{
//   							Dimensions: []interface{}{
//   								&DimensionProperty{
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					ActionName: jsii.String("actionName"),
//   				},
//   			},
//   			StatelessRules: []interface{}{
//   				&StatelessRuleProperty{
//   					Priority: jsii.Number(123),
//   					RuleDefinition: &RuleDefinitionProperty{
//   						Actions: []*string{
//   							jsii.String("actions"),
//   						},
//   						MatchAttributes: &MatchAttributesProperty{
//   							DestinationPorts: []interface{}{
//   								&PortRangeProperty{
//   									FromPort: jsii.Number(123),
//   									ToPort: jsii.Number(123),
//   								},
//   							},
//   							Destinations: []interface{}{
//   								&AddressProperty{
//   									AddressDefinition: jsii.String("addressDefinition"),
//   								},
//   							},
//   							Protocols: []interface{}{
//   								jsii.Number(123),
//   							},
//   							SourcePorts: []interface{}{
//   								&PortRangeProperty{
//   									FromPort: jsii.Number(123),
//   									ToPort: jsii.Number(123),
//   								},
//   							},
//   							Sources: []interface{}{
//   								&AddressProperty{
//   									AddressDefinition: jsii.String("addressDefinition"),
//   								},
//   							},
//   							TcpFlags: []interface{}{
//   								&TCPFlagFieldProperty{
//   									Flags: []*string{
//   										jsii.String("flags"),
//   									},
//   									Masks: []*string{
//   										jsii.String("masks"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	RuleVariables: &RuleVariablesProperty{
//   		IpSets: map[string]interface{}{
//   			"ipSetsKey": map[string][]*string{
//   				"definition": []*string{
//   					jsii.String("definition"),
//   				},
//   			},
//   		},
//   		PortSets: map[string]interface{}{
//   			"portSetsKey": &PortSetProperty{
//   				"definition": []*string{
//   					jsii.String("definition"),
//   				},
//   			},
//   		},
//   	},
//   	StatefulRuleOptions: &StatefulRuleOptionsProperty{
//   		RuleOrder: jsii.String("ruleOrder"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulegroup.html
//
type CfnRuleGroupPropsMixin_RuleGroupProperty struct {
	// The reference sets for the stateful rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulegroup.html#cfn-networkfirewall-rulegroup-rulegroup-referencesets
	//
	ReferenceSets interface{} `field:"optional" json:"referenceSets" yaml:"referenceSets"`
	// The stateful rules or stateless rules for the rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulegroup.html#cfn-networkfirewall-rulegroup-rulegroup-rulessource
	//
	RulesSource interface{} `field:"optional" json:"rulesSource" yaml:"rulesSource"`
	// Settings that are available for use in the rules in the rule group.
	//
	// You can only use these for stateful rule groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulegroup.html#cfn-networkfirewall-rulegroup-rulegroup-rulevariables
	//
	RuleVariables interface{} `field:"optional" json:"ruleVariables" yaml:"ruleVariables"`
	// Additional options governing how Network Firewall handles stateful rules.
	//
	// The policies where you use your stateful rule group must have stateful rule options settings that are compatible with these settings. Some limitations apply; for more information, see [Strict evaluation order](https://docs.aws.amazon.com/network-firewall/latest/developerguide/suricata-limitations-caveats.html) in the *AWS Network Firewall Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulegroup.html#cfn-networkfirewall-rulegroup-rulegroup-statefulruleoptions
	//
	StatefulRuleOptions interface{} `field:"optional" json:"statefulRuleOptions" yaml:"statefulRuleOptions"`
}

