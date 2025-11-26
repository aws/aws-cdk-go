package previewawsconnectmixins


// Information about a question from an evaluation form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evaluationFormQuestionProperty := &EvaluationFormQuestionProperty{
//   	Enablement: &EvaluationFormItemEnablementConfigurationProperty{
//   		Action: jsii.String("action"),
//   		Condition: &EvaluationFormItemEnablementConditionProperty{
//   			Operands: []interface{}{
//   				&EvaluationFormItemEnablementConditionOperandProperty{
//   					Expression: &EvaluationFormItemEnablementExpressionProperty{
//   						Comparator: jsii.String("comparator"),
//   						Source: &EvaluationFormItemEnablementSourceProperty{
//   							RefId: jsii.String("refId"),
//   							Type: jsii.String("type"),
//   						},
//   						Values: []interface{}{
//   							&EvaluationFormItemEnablementSourceValueProperty{
//   								RefId: jsii.String("refId"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			Operator: jsii.String("operator"),
//   		},
//   		DefaultAction: jsii.String("defaultAction"),
//   	},
//   	Instructions: jsii.String("instructions"),
//   	NotApplicableEnabled: jsii.Boolean(false),
//   	QuestionType: jsii.String("questionType"),
//   	QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   		Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   			Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   				AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   					SourceType: jsii.String("sourceType"),
//   				},
//   				PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   					Label: jsii.String("label"),
//   				},
//   			},
//   			MaxValue: jsii.Number(123),
//   			MinValue: jsii.Number(123),
//   			Options: []interface{}{
//   				&EvaluationFormNumericQuestionOptionProperty{
//   					AutomaticFail: jsii.Boolean(false),
//   					AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   						TargetSection: jsii.String("targetSection"),
//   					},
//   					MaxValue: jsii.Number(123),
//   					MinValue: jsii.Number(123),
//   					Score: jsii.Number(123),
//   				},
//   			},
//   		},
//   		SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   			Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   				AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   					SourceType: jsii.String("sourceType"),
//   				},
//   				DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   				Options: []interface{}{
//   					&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   						RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   							Category: jsii.String("category"),
//   							Condition: jsii.String("condition"),
//   							OptionRefId: jsii.String("optionRefId"),
//   						},
//   					},
//   				},
//   			},
//   			DisplayAs: jsii.String("displayAs"),
//   			Options: []interface{}{
//   				&EvaluationFormSingleSelectQuestionOptionProperty{
//   					AutomaticFail: jsii.Boolean(false),
//   					AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   						TargetSection: jsii.String("targetSection"),
//   					},
//   					RefId: jsii.String("refId"),
//   					Score: jsii.Number(123),
//   					Text: jsii.String("text"),
//   				},
//   			},
//   		},
//   		Text: &EvaluationFormTextQuestionPropertiesProperty{
//   			Automation: &EvaluationFormTextQuestionAutomationProperty{
//   				AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   					SourceType: jsii.String("sourceType"),
//   				},
//   			},
//   		},
//   	},
//   	RefId: jsii.String("refId"),
//   	Title: jsii.String("title"),
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormQuestionProperty struct {
	// A question conditional enablement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html#cfn-connect-evaluationform-evaluationformquestion-enablement
	//
	Enablement interface{} `field:"optional" json:"enablement" yaml:"enablement"`
	// The instructions of the section.
	//
	// *Length Constraints* : Minimum length of 0. Maximum length of 1024.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html#cfn-connect-evaluationform-evaluationformquestion-instructions
	//
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// The flag to enable not applicable answers to the question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html#cfn-connect-evaluationform-evaluationformquestion-notapplicableenabled
	//
	NotApplicableEnabled interface{} `field:"optional" json:"notApplicableEnabled" yaml:"notApplicableEnabled"`
	// The type of the question.
	//
	// *Allowed values* : `NUMERIC` | `SINGLESELECT` | `TEXT`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html#cfn-connect-evaluationform-evaluationformquestion-questiontype
	//
	QuestionType *string `field:"optional" json:"questionType" yaml:"questionType"`
	// The properties of the type of question.
	//
	// Text questions do not have to define question type properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html#cfn-connect-evaluationform-evaluationformquestion-questiontypeproperties
	//
	QuestionTypeProperties interface{} `field:"optional" json:"questionTypeProperties" yaml:"questionTypeProperties"`
	// The identifier of the question. An identifier must be unique within the evaluation form.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 40.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html#cfn-connect-evaluationform-evaluationformquestion-refid
	//
	RefId *string `field:"optional" json:"refId" yaml:"refId"`
	// The title of the question.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 350.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html#cfn-connect-evaluationform-evaluationformquestion-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The scoring weight of the section.
	//
	// *Minimum* : 0
	//
	// *Maximum* : 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html#cfn-connect-evaluationform-evaluationformquestion-weight
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

