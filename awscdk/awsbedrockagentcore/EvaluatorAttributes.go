package awsbedrockagentcore


// Attributes for importing an existing Evaluator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluatorAttributes := &EvaluatorAttributes{
//   	EvaluatorArn: jsii.String("evaluatorArn"),
//   	EvaluatorId: jsii.String("evaluatorId"),
//
//   	// the properties below are optional
//   	EvaluatorName: jsii.String("evaluatorName"),
//   }
//
type EvaluatorAttributes struct {
	// The ARN of the evaluator.
	EvaluatorArn *string `field:"required" json:"evaluatorArn" yaml:"evaluatorArn"`
	// The ID of the evaluator.
	EvaluatorId *string `field:"required" json:"evaluatorId" yaml:"evaluatorId"`
	// The name of the evaluator.
	// Default: - No name available.
	//
	EvaluatorName *string `field:"optional" json:"evaluatorName" yaml:"evaluatorName"`
}

