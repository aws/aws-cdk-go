package awsdatabrew


// Represents a single step from a DataBrew recipe to be performed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recipeStepProperty := &recipeStepProperty{
//   	action: &actionProperty{
//   		operation: jsii.String("operation"),
//
//   		// the properties below are optional
//   		parameters: map[string]*string{
//   			"parametersKey": jsii.String("parameters"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	conditionExpressions: []interface{}{
//   		&conditionExpressionProperty{
//   			condition: jsii.String("condition"),
//   			targetColumn: jsii.String("targetColumn"),
//
//   			// the properties below are optional
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRecipe_RecipeStepProperty struct {
	// The particular action to be performed in the recipe step.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// One or more conditions that must be met for the recipe step to succeed.
	//
	// > All of the conditions in the array must be met. In other words, all of the conditions must be combined using a logical AND operation.
	ConditionExpressions interface{} `field:"optional" json:"conditionExpressions" yaml:"conditionExpressions"`
}

