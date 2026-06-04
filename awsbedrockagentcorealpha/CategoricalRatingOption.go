package awsbedrockagentcorealpha


// A categorical rating scale option for custom evaluators.
//
// Categorical scales define discrete labels for scoring agent performance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   categoricalRatingOption := &CategoricalRatingOption{
//   	Definition: jsii.String("definition"),
//   	Label: jsii.String("label"),
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type CategoricalRatingOption struct {
	// The description that explains what this rating represents.
	//
	// Example:
	//   "The response fully addresses the user query with accurate information."
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// The label for this rating option.
	//
	// Example:
	//   "Good"
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Label *string `field:"required" json:"label" yaml:"label"`
}

