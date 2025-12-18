package awsconnect


// Automation configuration for multi-select questions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormMultiSelectQuestionAutomationProperty := &EvaluationFormMultiSelectQuestionAutomationProperty{
//   	Options: []interface{}{
//   		&EvaluationFormMultiSelectQuestionAutomationOptionProperty{
//   			RuleCategory: &MultiSelectQuestionRuleCategoryAutomationProperty{
//   				Category: jsii.String("category"),
//   				Condition: jsii.String("condition"),
//   				OptionRefIds: []*string{
//   					jsii.String("optionRefIds"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   		SourceType: jsii.String("sourceType"),
//   	},
//   	DefaultOptionRefIds: []*string{
//   		jsii.String("defaultOptionRefIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionautomation.html
//
type CfnEvaluationForm_EvaluationFormMultiSelectQuestionAutomationProperty struct {
	// Automation options for the multi-select question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionautomation.html#cfn-connect-evaluationform-evaluationformmultiselectquestionautomation-options
	//
	Options interface{} `field:"required" json:"options" yaml:"options"`
	// A question automation answer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionautomation.html#cfn-connect-evaluationform-evaluationformmultiselectquestionautomation-answersource
	//
	AnswerSource interface{} `field:"optional" json:"answerSource" yaml:"answerSource"`
	// Reference IDs of default options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionautomation.html#cfn-connect-evaluationform-evaluationformmultiselectquestionautomation-defaultoptionrefids
	//
	DefaultOptionRefIds *[]*string `field:"optional" json:"defaultOptionRefIds" yaml:"defaultOptionRefIds"`
}

