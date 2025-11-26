package previewawsconnectmixins


// Information about the automation configuration in text questions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evaluationFormTextQuestionAutomationProperty := &EvaluationFormTextQuestionAutomationProperty{
//   	AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   		SourceType: jsii.String("sourceType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformtextquestionautomation.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormTextQuestionAutomationProperty struct {
	// Automation answer source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformtextquestionautomation.html#cfn-connect-evaluationform-evaluationformtextquestionautomation-answersource
	//
	AnswerSource interface{} `field:"optional" json:"answerSource" yaml:"answerSource"`
}

