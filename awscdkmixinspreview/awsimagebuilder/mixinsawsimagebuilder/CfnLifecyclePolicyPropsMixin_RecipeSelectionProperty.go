package mixinsawsimagebuilder


// Specifies an Image Builder recipe that the lifecycle policy uses for resource selection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recipeSelectionProperty := &RecipeSelectionProperty{
//   	Name: jsii.String("name"),
//   	SemanticVersion: jsii.String("semanticVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-recipeselection.html
//
type CfnLifecyclePolicyPropsMixin_RecipeSelectionProperty struct {
	// The name of an Image Builder recipe that the lifecycle policy uses for resource selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-recipeselection.html#cfn-imagebuilder-lifecyclepolicy-recipeselection-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The version of the Image Builder recipe specified by the `name` field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-recipeselection.html#cfn-imagebuilder-lifecyclepolicy-recipeselection-semanticversion
	//
	SemanticVersion *string `field:"optional" json:"semanticVersion" yaml:"semanticVersion"`
}

