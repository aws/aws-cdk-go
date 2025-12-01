package awsconnect


// An item at the root level.
//
// All items must be sections.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var evaluationFormSectionProperty_ EvaluationFormSectionProperty
//
//   evaluationFormBaseItemProperty := &EvaluationFormBaseItemProperty{
//   	Section: &EvaluationFormSectionProperty{
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
//   					Enablement: &EvaluationFormItemEnablementConfigurationProperty{
//   						Action: jsii.String("action"),
//   						Condition: &EvaluationFormItemEnablementConditionProperty{
//   							Operands: []interface{}{
//   								&EvaluationFormItemEnablementConditionOperandProperty{
//   									Expression: &EvaluationFormItemEnablementExpressionProperty{
//   										Comparator: jsii.String("comparator"),
//   										Source: &EvaluationFormItemEnablementSourceProperty{
//   											Type: jsii.String("type"),
//
//   											// the properties below are optional
//   											RefId: jsii.String("refId"),
//   										},
//   										Values: []interface{}{
//   											&EvaluationFormItemEnablementSourceValueProperty{
//   												RefId: jsii.String("refId"),
//   												Type: jsii.String("type"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							Operator: jsii.String("operator"),
//   						},
//
//   						// the properties below are optional
//   						DefaultAction: jsii.String("defaultAction"),
//   					},
//   					Instructions: jsii.String("instructions"),
//   					NotApplicableEnabled: jsii.Boolean(false),
//   					QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   						MultiSelect: &EvaluationFormMultiSelectQuestionPropertiesProperty{
//   							Options: []interface{}{
//   								&EvaluationFormMultiSelectQuestionOptionProperty{
//   									RefId: jsii.String("refId"),
//   									Text: jsii.String("text"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							Automation: &EvaluationFormMultiSelectQuestionAutomationProperty{
//   								Options: []interface{}{
//   									&EvaluationFormMultiSelectQuestionAutomationOptionProperty{
//   										RuleCategory: &MultiSelectQuestionRuleCategoryAutomationProperty{
//   											Category: jsii.String("category"),
//   											Condition: jsii.String("condition"),
//   											OptionRefIds: []*string{
//   												jsii.String("optionRefIds"),
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   									SourceType: jsii.String("sourceType"),
//   								},
//   								DefaultOptionRefIds: []*string{
//   									jsii.String("defaultOptionRefIds"),
//   								},
//   							},
//   							DisplayAs: jsii.String("displayAs"),
//   						},
//   						Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   							MaxValue: jsii.Number(123),
//   							MinValue: jsii.Number(123),
//
//   							// the properties below are optional
//   							Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   								AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   									SourceType: jsii.String("sourceType"),
//   								},
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
//   									AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   										TargetSection: jsii.String("targetSection"),
//   									},
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
//   									AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   										TargetSection: jsii.String("targetSection"),
//   									},
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
//   								AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   									SourceType: jsii.String("sourceType"),
//   								},
//   								DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   							},
//   							DisplayAs: jsii.String("displayAs"),
//   						},
//   						Text: &EvaluationFormTextQuestionPropertiesProperty{
//   							Automation: &EvaluationFormTextQuestionAutomationProperty{
//   								AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   									SourceType: jsii.String("sourceType"),
//   								},
//   							},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformbaseitem.html
//
type CfnEvaluationForm_EvaluationFormBaseItemProperty struct {
	// A subsection or inner section of an item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformbaseitem.html#cfn-connect-evaluationform-evaluationformbaseitem-section
	//
	Section interface{} `field:"required" json:"section" yaml:"section"`
}

