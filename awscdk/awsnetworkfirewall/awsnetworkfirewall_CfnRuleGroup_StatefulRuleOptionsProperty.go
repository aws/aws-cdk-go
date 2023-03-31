package awsnetworkfirewall


// Additional options governing how Network Firewall handles the rule group.
//
// You can only use these for stateful rule groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statefulRuleOptionsProperty := &statefulRuleOptionsProperty{
//   	ruleOrder: jsii.String("ruleOrder"),
//   }
//
type CfnRuleGroup_StatefulRuleOptionsProperty struct {
	// Indicates how to manage the order of the rule evaluation for the rule group.
	//
	// `DEFAULT_ACTION_ORDER` is the default behavior. Stateful rules are provided to the rule engine as Suricata compatible strings, and Suricata evaluates them based on certain settings. For more information, see [Evaluation order for stateful rules](https://docs.aws.amazon.com/network-firewall/latest/developerguide/suricata-rule-evaluation-order.html) in the *AWS Network Firewall Developer Guide* .
	RuleOrder *string `field:"optional" json:"ruleOrder" yaml:"ruleOrder"`
}

