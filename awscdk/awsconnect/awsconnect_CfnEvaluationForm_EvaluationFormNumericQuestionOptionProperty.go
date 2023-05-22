package awsconnect


// Information about the option range used for scoring in numeric questions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormNumericQuestionOptionProperty := &EvaluationFormNumericQuestionOptionProperty{
//   	MaxValue: jsii.Number(123),
//   	MinValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	AutomaticFail: jsii.Boolean(false),
//   	Score: jsii.Number(123),
//   }
//
type CfnEvaluationForm_EvaluationFormNumericQuestionOptionProperty struct {
	// The maximum answer value of the range option.
	MaxValue *float64 `field:"required" json:"maxValue" yaml:"maxValue"`
	// The minimum answer value of the range option.
	MinValue *float64 `field:"required" json:"minValue" yaml:"minValue"`
	// The flag to mark the option as automatic fail.
	//
	// If an automatic fail answer is provided, the overall evaluation gets a score of 0.
	AutomaticFail interface{} `field:"optional" json:"automaticFail" yaml:"automaticFail"`
	// The score assigned to answer values within the range option.
	//
	// *Minimum* : 0
	//
	// *Maximum* : 10.
	Score *float64 `field:"optional" json:"score" yaml:"score"`
}

