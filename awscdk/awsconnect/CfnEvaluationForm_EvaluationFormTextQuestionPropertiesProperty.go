package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormTextQuestionPropertiesProperty := &EvaluationFormTextQuestionPropertiesProperty{
//   	Automation: &EvaluationFormTextQuestionAutomationProperty{
//   		AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   			SourceType: jsii.String("sourceType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformtextquestionproperties.html
//
type CfnEvaluationForm_EvaluationFormTextQuestionPropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformtextquestionproperties.html#cfn-connect-evaluationform-evaluationformtextquestionproperties-automation
	//
	Automation interface{} `field:"optional" json:"automation" yaml:"automation"`
}

