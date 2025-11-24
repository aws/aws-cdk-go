package mixinsawsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRecipePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRecipeMixinProps := &CfnRecipeMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Steps: []interface{}{
//   		&RecipeStepProperty{
//   			Action: &ActionProperty{
//   				Operation: jsii.String("operation"),
//   				Parameters: map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   			},
//   			ConditionExpressions: []interface{}{
//   				&ConditionExpressionProperty{
//   					Condition: jsii.String("condition"),
//   					TargetColumn: jsii.String("targetColumn"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-recipe.html
//
type CfnRecipeMixinProps struct {
	// The description of the recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-recipe.html#cfn-databrew-recipe-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique name for the recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-recipe.html#cfn-databrew-recipe-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of steps that are defined by the recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-recipe.html#cfn-databrew-recipe-steps
	//
	Steps interface{} `field:"optional" json:"steps" yaml:"steps"`
	// Metadata tags that have been applied to the recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-recipe.html#cfn-databrew-recipe-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

