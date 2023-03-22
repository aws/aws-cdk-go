package awsdatabrew


// Represents an individual condition that evaluates to true or false.
//
// Conditions are used with recipe actions. The action is only performed for column values where the condition evaluates to true.
//
// If a recipe requires more than one condition, then the recipe must specify multiple `ConditionExpression` elements. Each condition is applied to the rows in a dataset first, before the recipe action is performed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionExpressionProperty := &conditionExpressionProperty{
//   	condition: jsii.String("condition"),
//   	targetColumn: jsii.String("targetColumn"),
//
//   	// the properties below are optional
//   	value: jsii.String("value"),
//   }
//
type CfnRecipe_ConditionExpressionProperty struct {
	// A specific condition to apply to a recipe action.
	//
	// For more information, see [Recipe structure](https://docs.aws.amazon.com/databrew/latest/dg/recipe-structure.html) in the *AWS Glue DataBrew Developer Guide* .
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// A column to apply this condition to.
	TargetColumn *string `field:"required" json:"targetColumn" yaml:"targetColumn"`
	// A value that the condition must evaluate to for the condition to succeed.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

