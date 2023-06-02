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
//   var evaluationFormSectionProperty_ evaluationFormSectionProperty
//
//   evaluationFormItemProperty := &EvaluationFormItemProperty{
//   	Question: &EvaluationFormQuestionProperty{
//   		QuestionType: jsii.String("questionType"),
//   		RefId: jsii.String("refId"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		Instructions: jsii.String("instructions"),
//   		NotApplicableEnabled: jsii.Boolean(false),
//   		QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   			Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   				MaxValue: jsii.Number(123),
//   				MinValue: jsii.Number(123),
//
//   				// the properties below are optional
//   				Automation: &EvaluationFormNumericQuestionAutomationProperty{
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
//   					DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   				},
//   				DisplayAs: jsii.String("displayAs"),
//   			},
//   		},
//   		Weight: jsii.Number(123),
//   	},
//   	Section: &evaluationFormSectionProperty{
//   		RefId: jsii.String("refId"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		Instructions: jsii.String("instructions"),
//   		Items: []interface{}{
//   			&EvaluationFormItemProperty{
//   				Question: &EvaluationFormQuestionProperty{
//   					QuestionType: jsii.String("questionType"),
//   					RefId: jsii.String("refId"),
//   					Title: jsii.String("title"),
//
//   					// the properties below are optional
//   					Instructions: jsii.String("instructions"),
//   					NotApplicableEnabled: jsii.Boolean(false),
//   					QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   						Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   							MaxValue: jsii.Number(123),
//   							MinValue: jsii.Number(123),
//
//   							// the properties below are optional
//   							Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   								PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   									Label: jsii.String("label"),
//   								},
//   							},
//   							Options: []interface{}{
//   								&EvaluationFormNumericQuestionOptionProperty{
//   									MaxValue: jsii.Number(123),
//   									MinValue: jsii.Number(123),
//
//   									// the properties below are optional
//   									AutomaticFail: jsii.Boolean(false),
//   									Score: jsii.Number(123),
//   								},
//   							},
//   						},
//   						SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   							Options: []interface{}{
//   								&EvaluationFormSingleSelectQuestionOptionProperty{
//   									RefId: jsii.String("refId"),
//   									Text: jsii.String("text"),
//
//   									// the properties below are optional
//   									AutomaticFail: jsii.Boolean(false),
//   									Score: jsii.Number(123),
//   								},
//   							},
//
//   							// the properties below are optional
//   							Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   								Options: []interface{}{
//   									&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   										RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   											Category: jsii.String("category"),
//   											Condition: jsii.String("condition"),
//   											OptionRefId: jsii.String("optionRefId"),
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   							},
//   							DisplayAs: jsii.String("displayAs"),
//   						},
//   					},
//   					Weight: jsii.Number(123),
//   				},
//   				Section: evaluationFormSectionProperty_,
//   			},
//   		},
//   		Weight: jsii.Number(123),
//   	},
//   }
//
type CfnEvaluationForm_EvaluationFormItemProperty struct {
	// The information of the question.
	Question interface{} `field:"optional" json:"question" yaml:"question"`
	// The information of the section.
	Section interface{} `field:"optional" json:"section" yaml:"section"`
}

