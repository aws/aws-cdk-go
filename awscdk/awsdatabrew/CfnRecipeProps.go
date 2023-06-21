package awsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRecipe`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecipeProps := &CfnRecipeProps{
//   	Name: jsii.String("name"),
//   	Steps: []interface{}{
//   		&RecipeStepProperty{
//   			Action: &ActionProperty{
//   				Operation: jsii.String("operation"),
//
//   				// the properties below are optional
//   				Parameters: map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			ConditionExpressions: []interface{}{
//   				&ConditionExpressionProperty{
//   					Condition: jsii.String("condition"),
//   					TargetColumn: jsii.String("targetColumn"),
//
//   					// the properties below are optional
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRecipeProps struct {
	// The unique name for the recipe.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of steps that are defined by the recipe.
	Steps interface{} `field:"required" json:"steps" yaml:"steps"`
	// The description of the recipe.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata tags that have been applied to the recipe.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

