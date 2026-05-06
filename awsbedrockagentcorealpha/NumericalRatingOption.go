package awsbedrockagentcorealpha


// A numerical rating scale option for custom evaluators.
//
// Numerical scales define labeled numeric values for scoring agent performance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   numericalRatingOption := &NumericalRatingOption{
//   	Definition: jsii.String("definition"),
//   	Label: jsii.String("label"),
//   	Value: jsii.Number(123),
//   }
//
// Experimental.
type NumericalRatingOption struct {
	// The description that explains what this numerical rating represents.
	//
	// Example:
	//   "The response is comprehensive, accurate, and well-structured."
	//
	// Experimental.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// The label for this rating option.
	//
	// Example:
	//   "Excellent"
	//
	// Experimental.
	Label *string `field:"required" json:"label" yaml:"label"`
	// The numerical value for this rating scale option.
	//
	// Example:
	//   5
	//
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

