package awsconnect


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
	// `CfnEvaluationForm.EvaluationFormNumericQuestionOptionProperty.MaxValue`.
	MaxValue *float64 `field:"required" json:"maxValue" yaml:"maxValue"`
	// `CfnEvaluationForm.EvaluationFormNumericQuestionOptionProperty.MinValue`.
	MinValue *float64 `field:"required" json:"minValue" yaml:"minValue"`
	// `CfnEvaluationForm.EvaluationFormNumericQuestionOptionProperty.AutomaticFail`.
	AutomaticFail interface{} `field:"optional" json:"automaticFail" yaml:"automaticFail"`
	// `CfnEvaluationForm.EvaluationFormNumericQuestionOptionProperty.Score`.
	Score *float64 `field:"optional" json:"score" yaml:"score"`
}

