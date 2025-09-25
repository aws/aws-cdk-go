package awsnetworkfirewall


// The inspection criteria and action for a single stateless rule.
//
// AWS Network Firewall inspects each packet for the specified matching criteria. When a packet matches the criteria, Network Firewall performs the rule's actions on the packet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleDefinitionProperty := &RuleDefinitionProperty{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	MatchAttributes: &MatchAttributesProperty{
//   		DestinationPorts: []interface{}{
//   			&PortRangeProperty{
//   				FromPort: jsii.Number(123),
//   				ToPort: jsii.Number(123),
//   			},
//   		},
//   		Destinations: []interface{}{
//   			&AddressProperty{
//   				AddressDefinition: jsii.String("addressDefinition"),
//   			},
//   		},
//   		Protocols: []interface{}{
//   			jsii.Number(123),
//   		},
//   		SourcePorts: []interface{}{
//   			&PortRangeProperty{
//   				FromPort: jsii.Number(123),
//   				ToPort: jsii.Number(123),
//   			},
//   		},
//   		Sources: []interface{}{
//   			&AddressProperty{
//   				AddressDefinition: jsii.String("addressDefinition"),
//   			},
//   		},
//   		TcpFlags: []interface{}{
//   			&TCPFlagFieldProperty{
//   				Flags: []*string{
//   					jsii.String("flags"),
//   				},
//
//   				// the properties below are optional
//   				Masks: []*string{
//   					jsii.String("masks"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-ruledefinition.html
//
type CfnRuleGroup_RuleDefinitionProperty struct {
	// The actions to take on a packet that matches one of the stateless rule definition's match attributes.
	//
	// You must specify a standard action and you can add custom actions.
	//
	// > Network Firewall only forwards a packet for stateful rule inspection if you specify `aws:forward_to_sfe` for a rule that the packet matches, or if the packet doesn't match any stateless rule and you specify `aws:forward_to_sfe` for the `StatelessDefaultActions` setting for the firewall policy.
	//
	// For every rule, you must specify exactly one of the following standard actions.
	//
	// - *aws:pass* - Discontinues all inspection of the packet and permits it to go to its intended destination.
	// - *aws:drop* - Discontinues all inspection of the packet and blocks it from going to its intended destination.
	// - *aws:forward_to_sfe* - Discontinues stateless inspection of the packet and forwards it to the stateful rule engine for inspection.
	//
	// Additionally, you can specify a custom action. To do this, you define a custom action by name and type, then provide the name you've assigned to the action in this `Actions` setting.
	//
	// To provide more than one action in this setting, separate the settings with a comma. For example, if you have a publish metrics custom action that you've named `MyMetricsAction` , then you could specify the standard action `aws:pass` combined with the custom action using `[“aws:pass”, “MyMetricsAction”]` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-ruledefinition.html#cfn-networkfirewall-rulegroup-ruledefinition-actions
	//
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Criteria for Network Firewall to use to inspect an individual packet in stateless rule inspection.
	//
	// Each match attributes set can include one or more items such as IP address, CIDR range, port number, protocol, and TCP flags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-ruledefinition.html#cfn-networkfirewall-rulegroup-ruledefinition-matchattributes
	//
	MatchAttributes interface{} `field:"required" json:"matchAttributes" yaml:"matchAttributes"`
}

