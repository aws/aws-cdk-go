package awsnetworkfirewall


// A single Suricata rules specification, for use in a stateful rule group.
//
// Use this option to specify a simple Suricata rule with protocol, source and destination, ports, direction, and rule options. For information about the Suricata `Rules` format, see [Rules Format](https://docs.aws.amazon.com/https://suricata.readthedocs.io/en/suricata-6.0.9/rules/intro.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statefulRuleProperty := &StatefulRuleProperty{
//   	Action: jsii.String("action"),
//   	Header: &HeaderProperty{
//   		Destination: jsii.String("destination"),
//   		DestinationPort: jsii.String("destinationPort"),
//   		Direction: jsii.String("direction"),
//   		Protocol: jsii.String("protocol"),
//   		Source: jsii.String("source"),
//   		SourcePort: jsii.String("sourcePort"),
//   	},
//   	RuleOptions: []interface{}{
//   		&RuleOptionProperty{
//   			Keyword: jsii.String("keyword"),
//
//   			// the properties below are optional
//   			Settings: []*string{
//   				jsii.String("settings"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-statefulrule.html
//
type CfnRuleGroup_StatefulRuleProperty struct {
	// Defines what Network Firewall should do with the packets in a traffic flow when the flow matches the stateful rule criteria.
	//
	// For all actions, Network Firewall performs the specified action and discontinues stateful inspection of the traffic flow.
	//
	// The actions for a stateful rule are defined as follows:
	//
	// - *PASS* - Permits the packets to go to the intended destination.
	// - *DROP* - Blocks the packets from going to the intended destination and sends an alert log message, if alert logging is configured in the `Firewall` `LoggingConfiguration` .
	// - *REJECT* - Drops traffic that matches the conditions of the stateful rule and sends a TCP reset packet back to sender of the packet. A TCP reset packet is a packet with no payload and a `RST` bit contained in the TCP header flags. `REJECT` is available only for TCP traffic.
	// - *ALERT* - Permits the packets to go to the intended destination and sends an alert log message, if alert logging is configured in the `Firewall` `LoggingConfiguration` .
	//
	// You can use this action to test a rule that you intend to use to drop traffic. You can enable the rule with `ALERT` action, verify in the logs that the rule is filtering as you want, then change the action to `DROP` .
	// - *REJECT* - Drops TCP traffic that matches the conditions of the stateful rule, and sends a TCP reset packet back to sender of the packet. A TCP reset packet is a packet with no payload and a `RST` bit contained in the TCP header flags. Also sends an alert log mesage if alert logging is configured in the `Firewall` `LoggingConfiguration` .
	//
	// `REJECT` isn't currently available for use with IMAP and FTP protocols.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-statefulrule.html#cfn-networkfirewall-rulegroup-statefulrule-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The stateful inspection criteria for this rule, used to inspect traffic flows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-statefulrule.html#cfn-networkfirewall-rulegroup-statefulrule-header
	//
	Header interface{} `field:"required" json:"header" yaml:"header"`
	// Additional settings for a stateful rule, provided as keywords and settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-statefulrule.html#cfn-networkfirewall-rulegroup-statefulrule-ruleoptions
	//
	RuleOptions interface{} `field:"required" json:"ruleOptions" yaml:"ruleOptions"`
}

