package previewawsconnectmixins


// Items that are part of the evaluation form.
//
// The total number of sections and questions must not exceed 100 each. Questions must be contained in a section.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var evaluationFormItemProperty_ EvaluationFormItemProperty
//
//   evaluationFormItemProperty := &EvaluationFormItemProperty{
//   	Question: &EvaluationFormQuestionProperty{
//   		Enablement: &EvaluationFormItemEnablementConfigurationProperty{
//   			Action: jsii.String("action"),
//   			Condition: &EvaluationFormItemEnablementConditionProperty{
//   				Operands: []interface{}{
//   					&EvaluationFormItemEnablementConditionOperandProperty{
//   						Expression: &EvaluationFormItemEnablementExpressionProperty{
//   							Comparator: jsii.String("comparator"),
//   							Source: &EvaluationFormItemEnablementSourceProperty{
//   								RefId: jsii.String("refId"),
//   								Type: jsii.String("type"),
//   							},
//   							Values: []interface{}{
//   								&EvaluationFormItemEnablementSourceValueProperty{
//   									RefId: jsii.String("refId"),
//   									Type: jsii.String("type"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				Operator: jsii.String("operator"),
//   			},
//   			DefaultAction: jsii.String("defaultAction"),
//   		},
//   		Instructions: jsii.String("instructions"),
//   		NotApplicableEnabled: jsii.Boolean(false),
//   		QuestionType: jsii.String("questionType"),
//   		QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   			MultiSelect: &EvaluationFormMultiSelectQuestionPropertiesProperty{
//   				Automation: &EvaluationFormMultiSelectQuestionAutomationProperty{
//   					AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   						SourceType: jsii.String("sourceType"),
//   					},
//   					DefaultOptionRefIds: []*string{
//   						jsii.String("defaultOptionRefIds"),
//   					},
//   					Options: []interface{}{
//   						&EvaluationFormMultiSelectQuestionAutomationOptionProperty{
//   							RuleCategory: &MultiSelectQuestionRuleCategoryAutomationProperty{
//   								Category: jsii.String("category"),
//   								Condition: jsii.String("condition"),
//   								OptionRefIds: []*string{
//   									jsii.String("optionRefIds"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				DisplayAs: jsii.String("displayAs"),
//   				Options: []interface{}{
//   					&EvaluationFormMultiSelectQuestionOptionProperty{
//   						RefId: jsii.String("refId"),
//   						Text: jsii.String("text"),
//   					},
//   				},
//   			},
//   			Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   				Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   					AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   						SourceType: jsii.String("sourceType"),
//   					},
//   					PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   						Label: jsii.String("label"),
//   					},
//   				},
//   				MaxValue: jsii.Number(123),
//   				MinValue: jsii.Number(123),
//   				Options: []interface{}{
//   					&EvaluationFormNumericQuestionOptionProperty{
//   						AutomaticFail: jsii.Boolean(false),
//   						AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   							TargetSection: jsii.String("targetSection"),
//   						},
//   						MaxValue: jsii.Number(123),
//   						MinValue: jsii.Number(123),
//   						Score: jsii.Number(123),
//   					},
//   				},
//   			},
//   			SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   				Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   					AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   						SourceType: jsii.String("sourceType"),
//   					},
//   					DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   					Options: []interface{}{
//   						&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   							RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   								Category: jsii.String("category"),
//   								Condition: jsii.String("condition"),
//   								OptionRefId: jsii.String("optionRefId"),
//   							},
//   						},
//   					},
//   				},
//   				DisplayAs: jsii.String("displayAs"),
//   				Options: []interface{}{
//   					&EvaluationFormSingleSelectQuestionOptionProperty{
//   						AutomaticFail: jsii.Boolean(false),
//   						AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   							TargetSection: jsii.String("targetSection"),
//   						},
//   						RefId: jsii.String("refId"),
//   						Score: jsii.Number(123),
//   						Text: jsii.String("text"),
//   					},
//   				},
//   			},
//   			Text: &EvaluationFormTextQuestionPropertiesProperty{
//   				Automation: &EvaluationFormTextQuestionAutomationProperty{
//   					AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   						SourceType: jsii.String("sourceType"),
//   					},
//   				},
//   			},
//   		},
//   		RefId: jsii.String("refId"),
//   		Title: jsii.String("title"),
//   		Weight: jsii.Number(123),
//   	},
//   	Section: &EvaluationFormSectionProperty{
//   		Instructions: jsii.String("instructions"),
//   		Items: []interface{}{
//   			evaluationFormItemProperty_,
//   		},
//   		RefId: jsii.String("refId"),
//   		Title: jsii.String("title"),
//   		Weight: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitem.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormItemProperty struct {
	// The information of the question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitem.html#cfn-connect-evaluationform-evaluationformitem-question
	//
	Question interface{} `field:"optional" json:"question" yaml:"question"`
	// The information of the section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitem.html#cfn-connect-evaluationform-evaluationformitem-section
	//
	Section interface{} `field:"optional" json:"section" yaml:"section"`
}

