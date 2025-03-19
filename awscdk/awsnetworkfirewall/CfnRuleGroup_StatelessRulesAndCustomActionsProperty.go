package awsnetworkfirewall


// Stateless inspection criteria.
//
// Each stateless rule group uses exactly one of these data types to define its stateless rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statelessRulesAndCustomActionsProperty := &StatelessRulesAndCustomActionsProperty{
//   	StatelessRules: []interface{}{
//   		&StatelessRuleProperty{
//   			Priority: jsii.Number(123),
//   			RuleDefinition: &RuleDefinitionProperty{
//   				Actions: []*string{
//   					jsii.String("actions"),
//   				},
//   				MatchAttributes: &MatchAttributesProperty{
//   					DestinationPorts: []interface{}{
//   						&PortRangeProperty{
//   							FromPort: jsii.Number(123),
//   							ToPort: jsii.Number(123),
//   						},
//   					},
//   					Destinations: []interface{}{
//   						&AddressProperty{
//   							AddressDefinition: jsii.String("addressDefinition"),
//   						},
//   					},
//   					Protocols: []interface{}{
//   						jsii.Number(123),
//   					},
//   					SourcePorts: []interface{}{
//   						&PortRangeProperty{
//   							FromPort: jsii.Number(123),
//   							ToPort: jsii.Number(123),
//   						},
//   					},
//   					Sources: []interface{}{
//   						&AddressProperty{
//   							AddressDefinition: jsii.String("addressDefinition"),
//   						},
//   					},
//   					TcpFlags: []interface{}{
//   						&TCPFlagFieldProperty{
//   							Flags: []*string{
//   								jsii.String("flags"),
//   							},
//
//   							// the properties below are optional
//   							Masks: []*string{
//   								jsii.String("masks"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	CustomActions: []interface{}{
//   		&CustomActionProperty{
//   			ActionDefinition: &ActionDefinitionProperty{
//   				PublishMetricAction: &PublishMetricActionProperty{
//   					Dimensions: []interface{}{
//   						&DimensionProperty{
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			ActionName: jsii.String("actionName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-statelessrulesandcustomactions.html
//
type CfnRuleGroup_StatelessRulesAndCustomActionsProperty struct {
	// Defines the set of stateless rules for use in a stateless rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-statelessrulesandcustomactions.html#cfn-networkfirewall-rulegroup-statelessrulesandcustomactions-statelessrules
	//
	StatelessRules interface{} `field:"required" json:"statelessRules" yaml:"statelessRules"`
	// Defines an array of individual custom action definitions that are available for use by the stateless rules in this `StatelessRulesAndCustomActions` specification.
	//
	// You name each custom action that you define, and then you can use it by name in your stateless rule `RuleGroup.RuleDefinition` `Actions` specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-statelessrulesandcustomactions.html#cfn-networkfirewall-rulegroup-statelessrulesandcustomactions-customactions
	//
	CustomActions interface{} `field:"optional" json:"customActions" yaml:"customActions"`
}

