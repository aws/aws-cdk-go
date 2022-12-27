package awsnetworkfirewall


// Configuration settings for the handling of the stateful rule groups in a firewall policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statefulEngineOptionsProperty := &statefulEngineOptionsProperty{
//   	ruleOrder: jsii.String("ruleOrder"),
//   	streamExceptionPolicy: jsii.String("streamExceptionPolicy"),
//   }
//
type CfnFirewallPolicy_StatefulEngineOptionsProperty struct {
	// Indicates how to manage the order of stateful rule evaluation for the policy.
	//
	// `DEFAULT_ACTION_ORDER` is the default behavior. Stateful rules are provided to the rule engine as Suricata compatible strings, and Suricata evaluates them based on certain settings. For more information, see [Evaluation order for stateful rules](https://docs.aws.amazon.com/network-firewall/latest/developerguide/suricata-rule-evaluation-order.html) in the *AWS Network Firewall Developer Guide* .
	RuleOrder *string `field:"optional" json:"ruleOrder" yaml:"ruleOrder"`
	// `CfnFirewallPolicy.StatefulEngineOptionsProperty.StreamExceptionPolicy`.
	StreamExceptionPolicy *string `field:"optional" json:"streamExceptionPolicy" yaml:"streamExceptionPolicy"`
}

