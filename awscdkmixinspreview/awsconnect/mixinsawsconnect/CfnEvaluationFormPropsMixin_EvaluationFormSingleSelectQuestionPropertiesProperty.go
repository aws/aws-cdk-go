package mixinsawsconnect


// Information about the options in single select questions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evaluationFormSingleSelectQuestionPropertiesProperty := &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   	Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   		AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   			SourceType: jsii.String("sourceType"),
//   		},
//   		DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   		Options: []interface{}{
//   			&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   				RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   					Category: jsii.String("category"),
//   					Condition: jsii.String("condition"),
//   					OptionRefId: jsii.String("optionRefId"),
//   				},
//   			},
//   		},
//   	},
//   	DisplayAs: jsii.String("displayAs"),
//   	Options: []interface{}{
//   		&EvaluationFormSingleSelectQuestionOptionProperty{
//   			AutomaticFail: jsii.Boolean(false),
//   			AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   				TargetSection: jsii.String("targetSection"),
//   			},
//   			RefId: jsii.String("refId"),
//   			Score: jsii.Number(123),
//   			Text: jsii.String("text"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsingleselectquestionproperties.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormSingleSelectQuestionPropertiesProperty struct {
	// The display mode of the single select question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsingleselectquestionproperties.html#cfn-connect-evaluationform-evaluationformsingleselectquestionproperties-automation
	//
	Automation interface{} `field:"optional" json:"automation" yaml:"automation"`
	// The display mode of the single select question.
	//
	// *Allowed values* : `DROPDOWN` | `RADIO`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsingleselectquestionproperties.html#cfn-connect-evaluationform-evaluationformsingleselectquestionproperties-displayas
	//
	DisplayAs *string `field:"optional" json:"displayAs" yaml:"displayAs"`
	// The answer options of the single select question.
	//
	// *Minimum* : 2
	//
	// *Maximum* : 256.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsingleselectquestionproperties.html#cfn-connect-evaluationform-evaluationformsingleselectquestionproperties-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
}

