package awsnetworkfirewall


// A single Suricata rules specification, for use in a stateful rule group.
//
// Use this option to specify a simple Suricata rule with protocol, source and destination, ports, direction, and rule options. For information about the Suricata `Rules` format, see [Rules Format](https://docs.aws.amazon.com/https://suricata.readthedocs.io/en/suricata-5.0.0/rules/intro.html#) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statefulRuleProperty := &statefulRuleProperty{
//   	action: jsii.String("action"),
//   	header: &headerProperty{
//   		destination: jsii.String("destination"),
//   		destinationPort: jsii.String("destinationPort"),
//   		direction: jsii.String("direction"),
//   		protocol: jsii.String("protocol"),
//   		source: jsii.String("source"),
//   		sourcePort: jsii.String("sourcePort"),
//   	},
//   	ruleOptions: []interface{}{
//   		&ruleOptionProperty{
//   			keyword: jsii.String("keyword"),
//
//   			// the properties below are optional
//   			settings: []*string{
//   				jsii.String("settings"),
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_StatefulRuleProperty struct {
	// Defines what Network Firewall should do with the packets in a traffic flow when the flow matches the stateful rule criteria.
	//
	// For all actions, Network Firewall performs the specified action and discontinues stateful inspection of the traffic flow.
	//
	// The actions for a stateful rule are defined as follows:
	//
	// - *PASS* - Permits the packets to go to the intended destination.
	// - *DROP* - Blocks the packets from going to the intended destination and sends an alert log message, if alert logging is configured in the firewall's `LoggingConfiguration` .
	// - *ALERT* - Permits the packets to go to the intended destination and sends an alert log message, if alert logging is configured in the firewall's `LoggingConfiguration` .
	//
	// You can use this action to test a rule that you intend to use to drop traffic. You can enable the rule with `ALERT` action, verify in the logs that the rule is filtering as you want, then change the action to `DROP` .
	Action *string `field:"required" json:"action" yaml:"action"`
	// The stateful inspection criteria for this rule, used to inspect traffic flows.
	Header interface{} `field:"required" json:"header" yaml:"header"`
	// Additional settings for a stateful rule, provided as keywords and settings.
	RuleOptions interface{} `field:"required" json:"ruleOptions" yaml:"ruleOptions"`
}

