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
//   rulesSourceProperty := &rulesSourceProperty{
//   	rulesSourceList: &rulesSourceListProperty{
//   		generatedRulesType: jsii.String("generatedRulesType"),
//   		targets: []*string{
//   			jsii.String("targets"),
//   		},
//   		targetTypes: []*string{
//   			jsii.String("targetTypes"),
//   		},
//   	},
//   	rulesString: jsii.String("rulesString"),
//   	statefulRules: []interface{}{
//   		&statefulRuleProperty{
//   			action: jsii.String("action"),
//   			header: &headerProperty{
//   				destination: jsii.String("destination"),
//   				destinationPort: jsii.String("destinationPort"),
//   				direction: jsii.String("direction"),
//   				protocol: jsii.String("protocol"),
//   				source: jsii.String("source"),
//   				sourcePort: jsii.String("sourcePort"),
//   			},
//   			ruleOptions: []interface{}{
//   				&ruleOptionProperty{
//   					keyword: jsii.String("keyword"),
//
//   					// the properties below are optional
//   					settings: []*string{
//   						jsii.String("settings"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	statelessRulesAndCustomActions: &statelessRulesAndCustomActionsProperty{
//   		statelessRules: []interface{}{
//   			&statelessRuleProperty{
//   				priority: jsii.Number(123),
//   				ruleDefinition: &ruleDefinitionProperty{
//   					actions: []*string{
//   						jsii.String("actions"),
//   					},
//   					matchAttributes: &matchAttributesProperty{
//   						destinationPorts: []interface{}{
//   							&portRangeProperty{
//   								fromPort: jsii.Number(123),
//   								toPort: jsii.Number(123),
//   							},
//   						},
//   						destinations: []interface{}{
//   							&addressProperty{
//   								addressDefinition: jsii.String("addressDefinition"),
//   							},
//   						},
//   						protocols: []interface{}{
//   							jsii.Number(123),
//   						},
//   						sourcePorts: []interface{}{
//   							&portRangeProperty{
//   								fromPort: jsii.Number(123),
//   								toPort: jsii.Number(123),
//   							},
//   						},
//   						sources: []interface{}{
//   							&addressProperty{
//   								addressDefinition: jsii.String("addressDefinition"),
//   							},
//   						},
//   						tcpFlags: []interface{}{
//   							&tCPFlagFieldProperty{
//   								flags: []*string{
//   									jsii.String("flags"),
//   								},
//
//   								// the properties below are optional
//   								masks: []*string{
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
//   		customActions: []interface{}{
//   			&customActionProperty{
//   				actionDefinition: &actionDefinitionProperty{
//   					publishMetricAction: &publishMetricActionProperty{
//   						dimensions: []interface{}{
//   							&dimensionProperty{
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				actionName: jsii.String("actionName"),
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_RulesSourceProperty struct {
	// Stateful inspection criteria for a domain list rule group.
	RulesSourceList interface{} `field:"optional" json:"rulesSourceList" yaml:"rulesSourceList"`
	// Stateful inspection criteria, provided in Suricata compatible intrusion prevention system (IPS) rules.
	//
	// Suricata is an open-source network IPS that includes a standard rule-based language for network traffic inspection.
	//
	// These rules contain the inspection criteria and the action to take for traffic that matches the criteria, so this type of rule group doesn't have a separate action setting.
	RulesString *string `field:"optional" json:"rulesString" yaml:"rulesString"`
	// An array of individual stateful rules inspection criteria to be used together in a stateful rule group.
	//
	// Use this option to specify simple Suricata rules with protocol, source and destination, ports, direction, and rule options. For information about the Suricata `Rules` format, see [Rules Format](https://docs.aws.amazon.com/https://suricata.readthedocs.io/en/suricata-5.0.0/rules/intro.html#) .
	StatefulRules interface{} `field:"optional" json:"statefulRules" yaml:"statefulRules"`
	// Stateless inspection criteria to be used in a stateless rule group.
	StatelessRulesAndCustomActions interface{} `field:"optional" json:"statelessRulesAndCustomActions" yaml:"statelessRulesAndCustomActions"`
}

