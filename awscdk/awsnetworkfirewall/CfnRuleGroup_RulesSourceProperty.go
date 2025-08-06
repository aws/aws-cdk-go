package awsnetworkfirewall


// The stateless or stateful rules definitions for use in a single rule group.
//
// Each rule group requires a single `RulesSource` . You can use an instance of this for either stateless rules or stateful rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rulesSourceProperty := &RulesSourceProperty{
//   	RulesSourceList: &RulesSourceListProperty{
//   		GeneratedRulesType: jsii.String("generatedRulesType"),
//   		Targets: []*string{
//   			jsii.String("targets"),
//   		},
//   		TargetTypes: []*string{
//   			jsii.String("targetTypes"),
//   		},
//   	},
//   	RulesString: jsii.String("rulesString"),
//   	StatefulRules: []interface{}{
//   		&StatefulRuleProperty{
//   			Action: jsii.String("action"),
//   			Header: &HeaderProperty{
//   				Destination: jsii.String("destination"),
//   				DestinationPort: jsii.String("destinationPort"),
//   				Direction: jsii.String("direction"),
//   				Protocol: jsii.String("protocol"),
//   				Source: jsii.String("source"),
//   				SourcePort: jsii.String("sourcePort"),
//   			},
//   			RuleOptions: []interface{}{
//   				&RuleOptionProperty{
//   					Keyword: jsii.String("keyword"),
//
//   					// the properties below are optional
//   					Settings: []*string{
//   						jsii.String("settings"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	StatelessRulesAndCustomActions: &StatelessRulesAndCustomActionsProperty{
//   		StatelessRules: []interface{}{
//   			&StatelessRuleProperty{
//   				Priority: jsii.Number(123),
//   				RuleDefinition: &RuleDefinitionProperty{
//   					Actions: []*string{
//   						jsii.String("actions"),
//   					},
//   					MatchAttributes: &MatchAttributesProperty{
//   						DestinationPorts: []interface{}{
//   							&PortRangeProperty{
//   								FromPort: jsii.Number(123),
//   								ToPort: jsii.Number(123),
//   							},
//   						},
//   						Destinations: []interface{}{
//   							&AddressProperty{
//   								AddressDefinition: jsii.String("addressDefinition"),
//   							},
//   						},
//   						Protocols: []interface{}{
//   							jsii.Number(123),
//   						},
//   						SourcePorts: []interface{}{
//   							&PortRangeProperty{
//   								FromPort: jsii.Number(123),
//   								ToPort: jsii.Number(123),
//   							},
//   						},
//   						Sources: []interface{}{
//   							&AddressProperty{
//   								AddressDefinition: jsii.String("addressDefinition"),
//   							},
//   						},
//   						TcpFlags: []interface{}{
//   							&TCPFlagFieldProperty{
//   								Flags: []*string{
//   									jsii.String("flags"),
//   								},
//
//   								// the properties below are optional
//   								Masks: []*string{
//   									jsii.String("masks"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		CustomActions: []interface{}{
//   			&CustomActionProperty{
//   				ActionDefinition: &ActionDefinitionProperty{
//   					PublishMetricAction: &PublishMetricActionProperty{
//   						Dimensions: []interface{}{
//   							&DimensionProperty{
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				ActionName: jsii.String("actionName"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulessource.html
//
type CfnRuleGroup_RulesSourceProperty struct {
	// Stateful inspection criteria for a domain list rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulessource.html#cfn-networkfirewall-rulegroup-rulessource-rulessourcelist
	//
	RulesSourceList interface{} `field:"optional" json:"rulesSourceList" yaml:"rulesSourceList"`
	// Stateful inspection criteria, provided in Suricata compatible rules.
	//
	// Suricata is an open-source threat detection framework that includes a standard rule-based language for network traffic inspection.
	//
	// These rules contain the inspection criteria and the action to take for traffic that matches the criteria, so this type of rule group doesn't have a separate action setting.
	//
	// > You can't use the `priority` keyword if the `RuleOrder` option in `StatefulRuleOptions` is set to `STRICT_ORDER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulessource.html#cfn-networkfirewall-rulegroup-rulessource-rulesstring
	//
	RulesString *string `field:"optional" json:"rulesString" yaml:"rulesString"`
	// An array of individual stateful rules inspection criteria to be used together in a stateful rule group.
	//
	// Use this option to specify simple Suricata rules with protocol, source and destination, ports, direction, and rule options. For information about the Suricata `Rules` format, see [Rules Format](https://docs.aws.amazon.com/https://suricata.readthedocs.io/en/suricata-7.0.3/rules/intro.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulessource.html#cfn-networkfirewall-rulegroup-rulessource-statefulrules
	//
	StatefulRules interface{} `field:"optional" json:"statefulRules" yaml:"statefulRules"`
	// Stateless inspection criteria to be used in a stateless rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulessource.html#cfn-networkfirewall-rulegroup-rulessource-statelessrulesandcustomactions
	//
	StatelessRulesAndCustomActions interface{} `field:"optional" json:"statelessRulesAndCustomActions" yaml:"statelessRulesAndCustomActions"`
}

