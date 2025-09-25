package awsconnect


// A reference to a EvaluationForm resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormReference := &EvaluationFormReference{
//   	EvaluationFormArn: jsii.String("evaluationFormArn"),
//   }
//
type EvaluationFormReference struct {
	// The EvaluationFormArn of the EvaluationForm resource.
	EvaluationFormArn *string `field:"required" json:"evaluationFormArn" yaml:"evaluationFormArn"`
}

