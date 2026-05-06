package awsbedrockagentcorealpha


// Attributes for importing an existing Evaluator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   evaluatorAttributes := &EvaluatorAttributes{
//   	EvaluatorArn: jsii.String("evaluatorArn"),
//   	EvaluatorId: jsii.String("evaluatorId"),
//
//   	// the properties below are optional
//   	EvaluatorName: jsii.String("evaluatorName"),
//   }
//
// Experimental.
type EvaluatorAttributes struct {
	// The ARN of the evaluator.
	// Experimental.
	EvaluatorArn *string `field:"required" json:"evaluatorArn" yaml:"evaluatorArn"`
	// The ID of the evaluator.
	// Experimental.
	EvaluatorId *string `field:"required" json:"evaluatorId" yaml:"evaluatorId"`
	// The name of the evaluator.
	// Default: - No name available.
	//
	// Experimental.
	EvaluatorName *string `field:"optional" json:"evaluatorName" yaml:"evaluatorName"`
}

