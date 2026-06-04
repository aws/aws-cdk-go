package awsbedrockagentcorealpha


// The result of binding an EvaluatorReference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   evaluatorReferenceBindResult := &EvaluatorReferenceBindResult{
//   	EvaluatorId: jsii.String("evaluatorId"),
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type EvaluatorReferenceBindResult struct {
	// The evaluator identifier.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	EvaluatorId *string `field:"required" json:"evaluatorId" yaml:"evaluatorId"`
}

