package awsservicecatalog


// Properties for provisoning rule constraint.
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
type CloudFormationRuleConstraintOptions struct {
	// The description of the constraint.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The language code.
	//
	// Configures the language for error messages from service catalog.
	MessageLanguage MessageLanguage `field:"optional" json:"messageLanguage" yaml:"messageLanguage"`
	// The rule with condition and assertions to apply to template.
	Rule *TemplateRule `field:"required" json:"rule" yaml:"rule"`
}

