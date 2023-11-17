package awsimagebuilder


// The resource selection for the lifecycle policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceSelectionProperty := &ResourceSelectionProperty{
//   	Recipes: []interface{}{
//   		&RecipeSelectionProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			SemanticVersion: jsii.String("semanticVersion"),
//   		},
//   	},
//   	TagMap: map[string]*string{
//   		"tagMapKey": jsii.String("tagMap"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-resourceselection.html
//
type CfnLifecyclePolicy_ResourceSelectionProperty struct {
	// The recipes to select.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-resourceselection.html#cfn-imagebuilder-lifecyclepolicy-resourceselection-recipes
	//
	Recipes interface{} `field:"optional" json:"recipes" yaml:"recipes"`
	// The Image Builder resources to select by tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-resourceselection.html#cfn-imagebuilder-lifecyclepolicy-resourceselection-tagmap
	//
	TagMap interface{} `field:"optional" json:"tagMap" yaml:"tagMap"`
}

