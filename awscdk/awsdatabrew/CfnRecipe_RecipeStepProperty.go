package awsdatabrew


// Represents a single step from a DataBrew recipe to be performed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recipeStepProperty := &RecipeStepProperty{
//   	Action: &ActionProperty{
//   		Operation: jsii.String("operation"),
//
//   		// the properties below are optional
//   		Parameters: map[string]*string{
//   			"parametersKey": jsii.String("parameters"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ConditionExpressions: []interface{}{
//   		&ConditionExpressionProperty{
//   			Condition: jsii.String("condition"),
//   			TargetColumn: jsii.String("targetColumn"),
//
//   			// the properties below are optional
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipestep.html
//
type CfnRecipe_RecipeStepProperty struct {
	// The particular action to be performed in the recipe step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipestep.html#cfn-databrew-recipe-recipestep-action
	//
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// One or more conditions that must be met for the recipe step to succeed.
	//
	// > All of the conditions in the array must be met. In other words, all of the conditions must be combined using a logical AND operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipestep.html#cfn-databrew-recipe-recipestep-conditionexpressions
	//
	ConditionExpressions interface{} `field:"optional" json:"conditionExpressions" yaml:"conditionExpressions"`
}

