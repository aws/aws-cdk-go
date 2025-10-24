package awsconnect


// Items that are part of the evaluation form.
//
// The total number of sections and questions must not exceed 100 each. Questions must be contained in a section.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var evaluationFormItemProperty_ EvaluationFormItemProperty
//
//   evaluationFormItemProperty := &EvaluationFormItemProperty{
//   	Question: &EvaluationFormQuestionProperty{
//   		QuestionType: jsii.String("questionType"),
//   		RefId: jsii.String("refId"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		Enablement: &EvaluationFormItemEnablementConfigurationProperty{
//   			Action: jsii.String("action"),
//   			Condition: &EvaluationFormItemEnablementConditionProperty{
//   				Operands: []interface{}{
//   					&EvaluationFormItemEnablementConditionOperandProperty{
//   						Expression: &EvaluationFormItemEnablementExpressionProperty{
//   							Comparator: jsii.String("comparator"),
//   							Source: &EvaluationFormItemEnablementSourceProperty{
//   								Type: jsii.String("type"),
//
//   								// the properties below are optional
//   								RefId: jsii.String("refId"),
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
//
//   				// the properties below are optional
//   				Operator: jsii.String("operator"),
//   			},
//
//   			// the properties below are optional
//   			DefaultAction: jsii.String("defaultAction"),
//   		},
//   		Instructions: jsii.String("instructions"),
//   		NotApplicableEnabled: jsii.Boolean(false),
//   		QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   			Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   				MaxValue: jsii.Number(123),
//   				MinValue: jsii.Number(123),
//
//   				// the properties below are optional
//   				Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   					AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   						SourceType: jsii.String("sourceType"),
//   					},
//   					PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   						Label: jsii.String("label"),
//   					},
//   				},
//   				Options: []interface{}{
//   					&EvaluationFormNumericQuestionOptionProperty{
//   						MaxValue: jsii.Number(123),
//   						MinValue: jsii.Number(123),
//
//   						// the properties below are optional
//   						AutomaticFail: jsii.Boolean(false),
//   						AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   							TargetSection: jsii.String("targetSection"),
//   						},
//   						Score: jsii.Number(123),
//   					},
//   				},
//   			},
//   			SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   				Options: []interface{}{
//   					&EvaluationFormSingleSelectQuestionOptionProperty{
//   						RefId: jsii.String("refId"),
//   						Text: jsii.String("text"),
//
//   						// the properties below are optional
//   						AutomaticFail: jsii.Boolean(false),
//   						AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   							TargetSection: jsii.String("targetSection"),
//   						},
//   						Score: jsii.Number(123),
//   					},
//   				},
//
//   				// the properties below are optional
//   				Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   					Options: []interface{}{
//   						&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   							RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   								Category: jsii.String("category"),
//   								Condition: jsii.String("condition"),
//   								OptionRefId: jsii.String("optionRefId"),
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   						SourceType: jsii.String("sourceType"),
//   					},
//   					DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   				},
//   				DisplayAs: jsii.String("displayAs"),
//   			},
//   			Text: &EvaluationFormTextQuestionPropertiesProperty{
//   				Automation: &EvaluationFormTextQuestionAutomationProperty{
//   					AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   						SourceType: jsii.String("sourceType"),
//   					},
//   				},
//   			},
//   		},
//   		Weight: jsii.Number(123),
//   	},
//   	Section: &EvaluationFormSectionProperty{
//   		RefId: jsii.String("refId"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		Instructions: jsii.String("instructions"),
//   		Items: []interface{}{
//   			evaluationFormItemProperty_,
//   		},
//   		Weight: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitem.html
//
type CfnEvaluationForm_EvaluationFormItemProperty struct {
	// The information of the question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitem.html#cfn-connect-evaluationform-evaluationformitem-question
	//
	Question interface{} `field:"optional" json:"question" yaml:"question"`
	// The information of the section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitem.html#cfn-connect-evaluationform-evaluationformitem-section
	//
	Section interface{} `field:"optional" json:"section" yaml:"section"`
}

