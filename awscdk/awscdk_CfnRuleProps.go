// Version 2 of the AWS Cloud Development Kit library
package awscdk


// A rule can include a RuleCondition property and must include an Assertions property.
//
// For each rule, you can define only one rule condition; you can define one or more asserts within the Assertions property.
// You define a rule condition and assertions by using rule-specific intrinsic functions.
//
// You can use the following rule-specific intrinsic functions to define rule conditions and assertions:
//
//   Fn::And
//   Fn::Contains
//   Fn::EachMemberEquals
//   Fn::EachMemberIn
//   Fn::Equals
//   Fn::If
//   Fn::Not
//   Fn::Or
//   Fn::RefAll
//   Fn::ValueOf
//   Fn::ValueOfAll
//
// https://docs.aws.amazon.com/servicecatalog/latest/adminguide/reference-template_constraint_rules.html
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnConditionExpression iCfnConditionExpression
//
//   cfnRuleProps := &cfnRuleProps{
//   	assertions: []cfnRuleAssertion{
//   		&cfnRuleAssertion{
//   			assert: cfnConditionExpression,
//   			assertDescription: jsii.String("assertDescription"),
//   		},
//   	},
//   	ruleCondition: cfnConditionExpression,
//   }
//
type CfnRuleProps struct {
	// Assertions which define the rule.
	Assertions *[]*CfnRuleAssertion `field:"optional" json:"assertions" yaml:"assertions"`
	// If the rule condition evaluates to false, the rule doesn't take effect.
	//
	// If the function in the rule condition evaluates to true, expressions in each assert are evaluated and applied.
	RuleCondition ICfnConditionExpression `field:"optional" json:"ruleCondition" yaml:"ruleCondition"`
}

