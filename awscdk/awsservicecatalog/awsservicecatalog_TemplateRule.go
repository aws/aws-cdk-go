package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   portfolio.constrainCloudFormationParameters(product, &cloudFormationRuleConstraintOptions{
//   	rule: &templateRule{
//   		ruleName: jsii.String("testInstanceType"),
//   		condition: cdk.fn.conditionEquals(cdk.*fn.ref(jsii.String("Environment")), jsii.String("test")),
//   		assertions: []templateRuleAssertion{
//   			&templateRuleAssertion{
//   				assert: cdk.*fn.conditionContains([]*string{
//   					jsii.String("t2.micro"),
//   					jsii.String("t2.small"),
//   				}, cdk.*fn.ref(jsii.String("InstanceType"))),
//   				description: jsii.String("For test environment, the instance type should be small"),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type TemplateRule struct {
	// A list of assertions that make up the rule.
	// Experimental.
	Assertions *[]*TemplateRuleAssertion `field:"required" json:"assertions" yaml:"assertions"`
	// Name of the rule.
	// Experimental.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Specify when to apply rule with a rule-specific intrinsic function.
	// Experimental.
	Condition awscdk.ICfnRuleConditionExpression `field:"optional" json:"condition" yaml:"condition"`
}

