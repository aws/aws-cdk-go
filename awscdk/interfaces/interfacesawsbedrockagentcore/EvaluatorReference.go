package interfacesawsbedrockagentcore


// A reference to a Evaluator resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluatorReference := &EvaluatorReference{
//   	EvaluatorArn: jsii.String("evaluatorArn"),
//   }
//
type EvaluatorReference struct {
	// The EvaluatorArn of the Evaluator resource.
	EvaluatorArn *string `field:"required" json:"evaluatorArn" yaml:"evaluatorArn"`
}

