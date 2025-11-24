package mixinsawsnetworkfirewall


// An optional, non-standard action to use for stateless packet handling.
//
// You can define this in addition to the standard action that you must specify.
//
// You define and name the custom actions that you want to be able to use, and then you reference them by name in your actions settings.
//
// You can use custom actions in the following places:
//
// - In a `StatelessRulesAndCustomActions` . The custom actions are available for use by name inside the `StatelessRulesAndCustomActions` where you define them. You can use them for your stateless rule actions to specify what to do with a packet that matches the rule's match attributes.
// - In an firewall policy specification, in `StatelessCustomActions` . The custom actions are available for use inside the policy where you define them. You can use them for the policy's default stateless actions settings to specify what to do with packets that don't match any of the policy's stateless rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customActionProperty := &CustomActionProperty{
//   	ActionDefinition: &ActionDefinitionProperty{
//   		PublishMetricAction: &PublishMetricActionProperty{
//   			Dimensions: []interface{}{
//   				&DimensionProperty{
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	ActionName: jsii.String("actionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-customaction.html
//
type CfnRuleGroupPropsMixin_CustomActionProperty struct {
	// The custom action associated with the action name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-customaction.html#cfn-networkfirewall-rulegroup-customaction-actiondefinition
	//
	ActionDefinition interface{} `field:"optional" json:"actionDefinition" yaml:"actionDefinition"`
	// The descriptive name of the custom action.
	//
	// You can't change the name of a custom action after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-customaction.html#cfn-networkfirewall-rulegroup-customaction-actionname
	//
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
}

