package awsconnect


// Information about the automation configuration in single select questions.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsingleselectquestionoption.html
//
type CfnEvaluationForm_EvaluationFormSingleSelectQuestionOptionProperty struct {
	// The identifier of the answer option. An identifier must be unique within the question.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 40.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsingleselectquestionoption.html#cfn-connect-evaluationform-evaluationformsingleselectquestionoption-refid
	//
	RefId *string `field:"required" json:"refId" yaml:"refId"`
	// The title of the answer option.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 128.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsingleselectquestionoption.html#cfn-connect-evaluationform-evaluationformsingleselectquestionoption-text
	//
	Text *string `field:"required" json:"text" yaml:"text"`
	// The flag to mark the option as automatic fail.
	//
	// If an automatic fail answer is provided, the overall evaluation gets a score of 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsingleselectquestionoption.html#cfn-connect-evaluationform-evaluationformsingleselectquestionoption-automaticfail
	//
	AutomaticFail interface{} `field:"optional" json:"automaticFail" yaml:"automaticFail"`
	// The score assigned to the answer option.
	//
	// *Minimum* : 0
	//
	// *Maximum* : 10.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsingleselectquestionoption.html#cfn-connect-evaluationform-evaluationformsingleselectquestionoption-score
	//
	Score *float64 `field:"optional" json:"score" yaml:"score"`
}

