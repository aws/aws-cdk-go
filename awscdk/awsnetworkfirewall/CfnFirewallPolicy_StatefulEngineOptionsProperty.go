package awsnetworkfirewall


// Configuration settings for the handling of the stateful rule groups in a firewall policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statefulEngineOptionsProperty := &StatefulEngineOptionsProperty{
//   	RuleOrder: jsii.String("ruleOrder"),
//   	StreamExceptionPolicy: jsii.String("streamExceptionPolicy"),
//   }
//
type CfnFirewallPolicy_StatefulEngineOptionsProperty struct {
	// Indicates how to manage the order of stateful rule evaluation for the policy.
	//
	// `DEFAULT_ACTION_ORDER` is the default behavior. Stateful rules are provided to the rule engine as Suricata compatible strings, and Suricata evaluates them based on certain settings. For more information, see [Evaluation order for stateful rules](https://docs.aws.amazon.com/network-firewall/latest/developerguide/suricata-rule-evaluation-order.html) in the *AWS Network Firewall Developer Guide* .
	RuleOrder *string `field:"optional" json:"ruleOrder" yaml:"ruleOrder"`
	// Configures how Network Firewall processes traffic when a network connection breaks midstream.
	//
	// Network connections can break due to disruptions in external networks or within the firewall itself.
	//
	// - `DROP` - Network Firewall fails closed and drops all subsequent traffic going to the firewall. This is the default behavior.
	// - `CONTINUE` - Network Firewall continues to apply rules to the subsequent traffic without context from traffic before the break. This impacts the behavior of rules that depend on this context. For example, if you have a stateful rule to `drop http` traffic, Network Firewall won't match the traffic for this rule because the service won't have the context from session initialization defining the application layer protocol as HTTP. However, this behavior is rule dependent—a TCP-layer rule using a `flow:stateless` rule would still match, as would the `aws:drop_strict` default action.
	StreamExceptionPolicy *string `field:"optional" json:"streamExceptionPolicy" yaml:"streamExceptionPolicy"`
}

