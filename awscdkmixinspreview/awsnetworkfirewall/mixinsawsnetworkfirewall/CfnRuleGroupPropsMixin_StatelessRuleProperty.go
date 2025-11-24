package mixinsawsnetworkfirewall


// A single stateless rule.
//
// This is used in `StatelessRulesAndCustomActions` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   statelessRuleProperty := &StatelessRuleProperty{
//   	Priority: jsii.Number(123),
//   	RuleDefinition: &RuleDefinitionProperty{
//   		Actions: []*string{
//   			jsii.String("actions"),
//   		},
//   		MatchAttributes: &MatchAttributesProperty{
//   			DestinationPorts: []interface{}{
//   				&PortRangeProperty{
//   					FromPort: jsii.Number(123),
//   					ToPort: jsii.Number(123),
//   				},
//   			},
//   			Destinations: []interface{}{
//   				&AddressProperty{
//   					AddressDefinition: jsii.String("addressDefinition"),
//   				},
//   			},
//   			Protocols: []interface{}{
//   				jsii.Number(123),
//   			},
//   			SourcePorts: []interface{}{
//   				&PortRangeProperty{
//   					FromPort: jsii.Number(123),
//   					ToPort: jsii.Number(123),
//   				},
//   			},
//   			Sources: []interface{}{
//   				&AddressProperty{
//   					AddressDefinition: jsii.String("addressDefinition"),
//   				},
//   			},
//   			TcpFlags: []interface{}{
//   				&TCPFlagFieldProperty{
//   					Flags: []*string{
//   						jsii.String("flags"),
//   					},
//   					Masks: []*string{
//   						jsii.String("masks"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-statelessrule.html
//
type CfnRuleGroupPropsMixin_StatelessRuleProperty struct {
	// Indicates the order in which to run this rule relative to all of the rules that are defined for a stateless rule group.
	//
	// Network Firewall evaluates the rules in a rule group starting with the lowest priority setting. You must ensure that the priority settings are unique for the rule group.
	//
	// Each stateless rule group uses exactly one `StatelessRulesAndCustomActions` object, and each `StatelessRulesAndCustomActions` contains exactly one `StatelessRules` object. To ensure unique priority settings for your rule groups, set unique priorities for the stateless rules that you define inside any single `StatelessRules` object.
	//
	// You can change the priority settings of your rules at any time. To make it easier to insert rules later, number them so there's a wide range in between, for example use 100, 200, and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-statelessrule.html#cfn-networkfirewall-rulegroup-statelessrule-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Defines the stateless 5-tuple packet inspection criteria and the action to take on a packet that matches the criteria.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-statelessrule.html#cfn-networkfirewall-rulegroup-statelessrule-ruledefinition
	//
	RuleDefinition interface{} `field:"optional" json:"ruleDefinition" yaml:"ruleDefinition"`
}

