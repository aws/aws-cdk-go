package awsnetworkfirewall


// The traffic filtering behavior of a firewall policy, defined in a collection of stateless and stateful rule groups and other settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firewallPolicyProperty := &firewallPolicyProperty{
//   	statelessDefaultActions: []*string{
//   		jsii.String("statelessDefaultActions"),
//   	},
//   	statelessFragmentDefaultActions: []*string{
//   		jsii.String("statelessFragmentDefaultActions"),
//   	},
//
//   	// the properties below are optional
//   	statefulDefaultActions: []*string{
//   		jsii.String("statefulDefaultActions"),
//   	},
//   	statefulEngineOptions: &statefulEngineOptionsProperty{
//   		ruleOrder: jsii.String("ruleOrder"),
//   		streamExceptionPolicy: jsii.String("streamExceptionPolicy"),
//   	},
//   	statefulRuleGroupReferences: []interface{}{
//   		&statefulRuleGroupReferenceProperty{
//   			resourceArn: jsii.String("resourceArn"),
//
//   			// the properties below are optional
//   			override: &statefulRuleGroupOverrideProperty{
//   				action: jsii.String("action"),
//   			},
//   			priority: jsii.Number(123),
//   		},
//   	},
//   	statelessCustomActions: []interface{}{
//   		&customActionProperty{
//   			actionDefinition: &actionDefinitionProperty{
//   				publishMetricAction: &publishMetricActionProperty{
//   					dimensions: []interface{}{
//   						&dimensionProperty{
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			actionName: jsii.String("actionName"),
//   		},
//   	},
//   	statelessRuleGroupReferences: []interface{}{
//   		&statelessRuleGroupReferenceProperty{
//   			priority: jsii.Number(123),
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   	},
//   }
//
type CfnFirewallPolicy_FirewallPolicyProperty struct {
	// The actions to take on a packet if it doesn't match any of the stateless rules in the policy.
	//
	// If you want non-matching packets to be forwarded for stateful inspection, specify `aws:forward_to_sfe` .
	//
	// You must specify one of the standard actions: `aws:pass` , `aws:drop` , or `aws:forward_to_sfe` . In addition, you can specify custom actions that are compatible with your standard section choice.
	//
	// For example, you could specify `["aws:pass"]` or you could specify `["aws:pass", “customActionName”]` . For information about compatibility, see the custom action descriptions.
	StatelessDefaultActions *[]*string `field:"required" json:"statelessDefaultActions" yaml:"statelessDefaultActions"`
	// The actions to take on a fragmented packet if it doesn't match any of the stateless rules in the policy.
	//
	// If you want non-matching fragmented packets to be forwarded for stateful inspection, specify `aws:forward_to_sfe` .
	//
	// You must specify one of the standard actions: `aws:pass` , `aws:drop` , or `aws:forward_to_sfe` . In addition, you can specify custom actions that are compatible with your standard section choice.
	//
	// For example, you could specify `["aws:pass"]` or you could specify `["aws:pass", “customActionName”]` . For information about compatibility, see the custom action descriptions.
	StatelessFragmentDefaultActions *[]*string `field:"required" json:"statelessFragmentDefaultActions" yaml:"statelessFragmentDefaultActions"`
	// The default actions to take on a packet that doesn't match any stateful rules.
	//
	// The stateful default action is optional, and is only valid when using the strict rule order.
	//
	// Valid values of the stateful default action:
	//
	// - aws:drop_strict
	// - aws:drop_established
	// - aws:alert_strict
	// - aws:alert_established
	//
	// For more information, see [Strict evaluation order](https://docs.aws.amazon.com/network-firewall/latest/developerguide/suricata-rule-evaluation-order.html#suricata-strict-rule-evaluation-order.html) in the *AWS Network Firewall Developer Guide* .
	StatefulDefaultActions *[]*string `field:"optional" json:"statefulDefaultActions" yaml:"statefulDefaultActions"`
	// Additional options governing how Network Firewall handles stateful rules.
	//
	// The stateful rule groups that you use in your policy must have stateful rule options settings that are compatible with these settings.
	StatefulEngineOptions interface{} `field:"optional" json:"statefulEngineOptions" yaml:"statefulEngineOptions"`
	// References to the stateful rule groups that are used in the policy.
	//
	// These define the inspection criteria in stateful rules.
	StatefulRuleGroupReferences interface{} `field:"optional" json:"statefulRuleGroupReferences" yaml:"statefulRuleGroupReferences"`
	// The custom action definitions that are available for use in the firewall policy's `StatelessDefaultActions` setting.
	//
	// You name each custom action that you define, and then you can use it by name in your default actions specifications.
	StatelessCustomActions interface{} `field:"optional" json:"statelessCustomActions" yaml:"statelessCustomActions"`
	// References to the stateless rule groups that are used in the policy.
	//
	// These define the matching criteria in stateless rules.
	StatelessRuleGroupReferences interface{} `field:"optional" json:"statelessRuleGroupReferences" yaml:"statelessRuleGroupReferences"`
}

