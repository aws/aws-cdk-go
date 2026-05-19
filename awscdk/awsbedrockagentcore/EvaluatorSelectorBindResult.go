package awsbedrockagentcore


// The result of binding an EvaluatorSelector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluatorSelectorBindResult := &EvaluatorSelectorBindResult{
//   	EvaluatorId: jsii.String("evaluatorId"),
//   }
//
type EvaluatorSelectorBindResult struct {
	// The evaluator identifier.
	EvaluatorId *string `field:"required" json:"evaluatorId" yaml:"evaluatorId"`
}

