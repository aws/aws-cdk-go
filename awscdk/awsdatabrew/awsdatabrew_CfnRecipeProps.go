package awsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRecipe`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecipeProps := &cfnRecipeProps{
//   	name: jsii.String("name"),
//   	steps: []interface{}{
//   		&recipeStepProperty{
//   			action: &actionProperty{
//   				operation: jsii.String("operation"),
//
//   				// the properties below are optional
//   				parameters: map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			conditionExpressions: []interface{}{
//   				&conditionExpressionProperty{
//   					condition: jsii.String("condition"),
//   					targetColumn: jsii.String("targetColumn"),
//
//   					// the properties below are optional
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

