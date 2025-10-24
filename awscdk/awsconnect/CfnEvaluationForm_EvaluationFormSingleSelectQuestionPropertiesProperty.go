package awsconnect


// Information about the options in single select questions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormSingleSelectQuestionPropertiesProperty := &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   	Options: []interface{}{
//   		&EvaluationFormSingleSelectQuestionOptionProperty{
//   			RefId: jsii.String("refId"),
//   			Text: jsii.String("text"),
//
//   			// the properties below are optional
//   			AutomaticFail: jsii.Boolean(false),
//   			AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   				TargetSection: jsii.String("targetSection"),
//   			},
//   			Score: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   		Options: []interface{}{
//   			&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   				RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   					Category: jsii.String("category"),
//   					Condition: jsii.String("condition"),
//   					OptionRefId: jsii.String("optionRefId"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   			SourceType: jsii.String("sourceType"),
//   		},
//   		DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   	},
//   	DisplayAs: jsii.String("displayAs"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsingleselectquestionproperties.html
//
type CfnEvaluationForm_EvaluationFormSingleSelectQuestionPropertiesProperty struct {
	// The answer options of the single select question.
	//
	// *Minimum* : 2
	//
	// *Maximum* : 256.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsingleselectquestionproperties.html#cfn-connect-evaluationform-evaluationformsingleselectquestionproperties-options
	//
	Options interface{} `field:"required" json:"options" yaml:"options"`
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
}

