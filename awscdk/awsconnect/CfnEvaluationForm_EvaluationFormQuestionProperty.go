package awsconnect


// Information about a question from an evaluation form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormQuestionProperty := &EvaluationFormQuestionProperty{
//   	QuestionType: jsii.String("questionType"),
//   	RefId: jsii.String("refId"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	Enablement: &EvaluationFormItemEnablementConfigurationProperty{
//   		Action: jsii.String("action"),
//   		Condition: &EvaluationFormItemEnablementConditionProperty{
//   			Operands: []interface{}{
//   				&EvaluationFormItemEnablementConditionOperandProperty{
//   					Expression: &EvaluationFormItemEnablementExpressionProperty{
//   						Comparator: jsii.String("comparator"),
//   						Source: &EvaluationFormItemEnablementSourceProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							RefId: jsii.String("refId"),
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
//
//   			// the properties below are optional
//   			Operator: jsii.String("operator"),
//   		},
//
//   		// the properties below are optional
//   		DefaultAction: jsii.String("defaultAction"),
//   	},
//   	Instructions: jsii.String("instructions"),
//   	NotApplicableEnabled: jsii.Boolean(false),
//   	QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   		Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   			MaxValue: jsii.Number(123),
//   			MinValue: jsii.Number(123),
//
//   			// the properties below are optional
//   			Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   				AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   					SourceType: jsii.String("sourceType"),
//   				},
//   				PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   					Label: jsii.String("label"),
//   				},
//   			},
//   			Options: []interface{}{
//   				&EvaluationFormNumericQuestionOptionProperty{
//   					MaxValue: jsii.Number(123),
//   					MinValue: jsii.Number(123),
//
//   					// the properties below are optional
//   					AutomaticFail: jsii.Boolean(false),
//   					AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   						TargetSection: jsii.String("targetSection"),
//   					},
//   					Score: jsii.Number(123),
//   				},
//   			},
//   		},
//   		SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   			Options: []interface{}{
//   				&EvaluationFormSingleSelectQuestionOptionProperty{
//   					RefId: jsii.String("refId"),
//   					Text: jsii.String("text"),
//
//   					// the properties below are optional
//   					AutomaticFail: jsii.Boolean(false),
//   					AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   						TargetSection: jsii.String("targetSection"),
//   					},
//   					Score: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   				Options: []interface{}{
//   					&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   						RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   							Category: jsii.String("category"),
//   							Condition: jsii.String("condition"),
//   							OptionRefId: jsii.String("optionRefId"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   					SourceType: jsii.String("sourceType"),
//   				},
//   				DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   			},
//   			DisplayAs: jsii.String("displayAs"),
//   		},
//   		Text: &EvaluationFormTextQuestionPropertiesProperty{
//   			Automation: &EvaluationFormTextQuestionAutomationProperty{
//   				AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   					SourceType: jsii.String("sourceType"),
//   				},
//   			},
//   		},
//   	},
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html
//
type CfnEvaluationForm_EvaluationFormQuestionProperty struct {
	// The type of the question.
	//
	// *Allowed values* : `NUMERIC` | `SINGLESELECT` | `TEXT`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html#cfn-connect-evaluationform-evaluationformquestion-questiontype
	//
	QuestionType *string `field:"required" json:"questionType" yaml:"questionType"`
	// The identifier of the question. An identifier must be unique within the evaluation form.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 40.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html#cfn-connect-evaluationform-evaluationformquestion-refid
	//
	RefId *string `field:"required" json:"refId" yaml:"refId"`
	// The title of the question.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 350.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html#cfn-connect-evaluationform-evaluationformquestion-title
	//
	Title *string `field:"required" json:"title" yaml:"title"`
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
	// The properties of the type of question.
	//
	// Text questions do not have to define question type properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html#cfn-connect-evaluationform-evaluationformquestion-questiontypeproperties
	//
	QuestionTypeProperties interface{} `field:"optional" json:"questionTypeProperties" yaml:"questionTypeProperties"`
	// The scoring weight of the section.
	//
	// *Minimum* : 0
	//
	// *Maximum* : 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestion.html#cfn-connect-evaluationform-evaluationformquestion-weight
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

