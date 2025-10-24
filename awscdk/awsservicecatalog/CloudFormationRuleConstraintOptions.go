package awsservicecatalog


// Properties for provisoning rule constraint.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var portfolio Portfolio
//   var product CloudFormationProduct
//
//
//   portfolio.constrainCloudFormationParameters(product, &CloudFormationRuleConstraintOptions{
//   	Rule: &TemplateRule{
//   		RuleName: jsii.String("testInstanceType"),
//   		Condition: awscdk.Fn_ConditionEquals(awscdk.Fn_Ref(jsii.String("Environment")), jsii.String("test")),
//   		Assertions: []TemplateRuleAssertion{
//   			&TemplateRuleAssertion{
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
type CloudFormationRuleConstraintOptions struct {
	// The description of the constraint.
	// Default: - No description provided.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The language code.
	//
	// Configures the language for error messages from service catalog.
	// Default: - English.
	//
	MessageLanguage MessageLanguage `field:"optional" json:"messageLanguage" yaml:"messageLanguage"`
	// The rule with condition and assertions to apply to template.
	Rule *TemplateRule `field:"required" json:"rule" yaml:"rule"`
}

