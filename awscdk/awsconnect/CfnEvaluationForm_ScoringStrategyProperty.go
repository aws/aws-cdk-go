package awsconnect


// A scoring strategy of the evaluation form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scoringStrategyProperty := &ScoringStrategyProperty{
//   	Mode: jsii.String("mode"),
//   	Status: jsii.String("status"),
//   }
//
type CfnEvaluationForm_ScoringStrategyProperty struct {
	// The scoring mode of the evaluation form.
	//
	// *Allowed values* : `QUESTION_ONLY` | `SECTION_ONLY`.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The scoring status of the evaluation form.
	//
	// *Allowed values* : `ENABLED` | `DISABLED`.
	Status *string `field:"required" json:"status" yaml:"status"`
}

