package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormSingleSelectQuestionOptionProperty := &EvaluationFormSingleSelectQuestionOptionProperty{
//   	RefId: jsii.String("refId"),
//   	Text: jsii.String("text"),
//
//   	// the properties below are optional
//   	AutomaticFail: jsii.Boolean(false),
//   	Score: jsii.Number(123),
//   }
//
type CfnEvaluationForm_EvaluationFormSingleSelectQuestionOptionProperty struct {
	// `CfnEvaluationForm.EvaluationFormSingleSelectQuestionOptionProperty.RefId`.
	RefId *string `field:"required" json:"refId" yaml:"refId"`
	// `CfnEvaluationForm.EvaluationFormSingleSelectQuestionOptionProperty.Text`.
	Text *string `field:"required" json:"text" yaml:"text"`
	// `CfnEvaluationForm.EvaluationFormSingleSelectQuestionOptionProperty.AutomaticFail`.
	AutomaticFail interface{} `field:"optional" json:"automaticFail" yaml:"automaticFail"`
	// `CfnEvaluationForm.EvaluationFormSingleSelectQuestionOptionProperty.Score`.
	Score *float64 `field:"optional" json:"score" yaml:"score"`
}

