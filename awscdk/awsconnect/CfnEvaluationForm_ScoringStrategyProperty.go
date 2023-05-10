package awsconnect


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
	// `CfnEvaluationForm.ScoringStrategyProperty.Mode`.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// `CfnEvaluationForm.ScoringStrategyProperty.Status`.
	Status *string `field:"required" json:"status" yaml:"status"`
}

