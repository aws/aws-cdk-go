package awsbedrockagentcore


// A categorical rating scale option for custom evaluators.
//
// Categorical scales define discrete labels for scoring agent performance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   categoricalRatingOption := &CategoricalRatingOption{
//   	Definition: jsii.String("definition"),
//   	Label: jsii.String("label"),
//   }
//
type CategoricalRatingOption struct {
	// The description that explains what this rating represents.
	//
	// Example:
	//   "The response fully addresses the user query with accurate information."
	//
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// The label for this rating option.
	//
	// Example:
	//   "Good"
	//
	Label *string `field:"required" json:"label" yaml:"label"`
}

