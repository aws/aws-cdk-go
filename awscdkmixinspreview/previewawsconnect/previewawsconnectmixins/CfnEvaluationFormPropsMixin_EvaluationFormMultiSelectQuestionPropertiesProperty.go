package previewawsconnectmixins


// Properties for a multi-select question in an evaluation form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evaluationFormMultiSelectQuestionPropertiesProperty := &EvaluationFormMultiSelectQuestionPropertiesProperty{
//   	Automation: &EvaluationFormMultiSelectQuestionAutomationProperty{
//   		AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   			SourceType: jsii.String("sourceType"),
//   		},
//   		DefaultOptionRefIds: []*string{
//   			jsii.String("defaultOptionRefIds"),
//   		},
//   		Options: []interface{}{
//   			&EvaluationFormMultiSelectQuestionAutomationOptionProperty{
//   				RuleCategory: &MultiSelectQuestionRuleCategoryAutomationProperty{
//   					Category: jsii.String("category"),
//   					Condition: jsii.String("condition"),
//   					OptionRefIds: []*string{
//   						jsii.String("optionRefIds"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	DisplayAs: jsii.String("displayAs"),
//   	Options: []interface{}{
//   		&EvaluationFormMultiSelectQuestionOptionProperty{
//   			RefId: jsii.String("refId"),
//   			Text: jsii.String("text"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionproperties.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormMultiSelectQuestionPropertiesProperty struct {
	// Automation configuration for this multi-select question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionproperties.html#cfn-connect-evaluationform-evaluationformmultiselectquestionproperties-automation
	//
	Automation interface{} `field:"optional" json:"automation" yaml:"automation"`
	// Display format for the multi-select question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionproperties.html#cfn-connect-evaluationform-evaluationformmultiselectquestionproperties-displayas
	//
	DisplayAs *string `field:"optional" json:"displayAs" yaml:"displayAs"`
	// Options available for this multi-select question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionproperties.html#cfn-connect-evaluationform-evaluationformmultiselectquestionproperties-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
}

