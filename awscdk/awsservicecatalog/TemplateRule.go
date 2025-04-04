package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Defines the provisioning template constraints.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var portfolio portfolio
//   var product cloudFormationProduct
//
//
//   portfolio.constrainCloudFormationParameters(product, &CloudFormationRuleConstraintOptions{
//   	Rule: &TemplateRule{
//   		RuleName: jsii.String("testInstanceType"),
//   		Condition: awscdk.Fn_ConditionEquals(awscdk.Fn_Ref(jsii.String("Environment")), jsii.String("test")),
//   		Assertions: []templateRuleAssertion{
//   			&templateRuleAssertion{
//   				Assert: awscdk.Fn_ConditionContains([]*string{
//   					jsii.String("t2.micro"),
//   					jsii.String("t2.small"),
//   				}, awscdk.Fn_*Ref(jsii.String("InstanceType"))),
//   				Description: jsii.String("For test environment, the instance type should be small"),
//   			},
//   		},
//   	},
//   })
//
type TemplateRule struct {
	// A list of assertions that make up the rule.
	Assertions *[]*TemplateRuleAssertion `field:"required" json:"assertions" yaml:"assertions"`
	// Name of the rule.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Specify when to apply rule with a rule-specific intrinsic function.
	// Default: - no rule condition provided.
	//
	Condition awscdk.ICfnRuleConditionExpression `field:"optional" json:"condition" yaml:"condition"`
}

