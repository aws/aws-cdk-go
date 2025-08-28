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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-scoringstrategy.html
//
type CfnEvaluationForm_ScoringStrategyProperty struct {
	// The scoring mode of the evaluation form.
	//
	// *Allowed values* : `QUESTION_ONLY` | `SECTION_ONLY`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-scoringstrategy.html#cfn-connect-evaluationform-scoringstrategy-mode
	//
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The scoring status of the evaluation form.
	//
	// *Allowed values* : `ENABLED` | `DISABLED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-scoringstrategy.html#cfn-connect-evaluationform-scoringstrategy-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
}

