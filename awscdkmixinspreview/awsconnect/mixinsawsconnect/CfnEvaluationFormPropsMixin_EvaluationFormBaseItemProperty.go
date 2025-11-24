package mixinsawsconnect


// An item at the root level.
//
// All items must be sections.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var evaluationFormSectionProperty_ EvaluationFormSectionProperty
//
//   evaluationFormBaseItemProperty := &EvaluationFormBaseItemProperty{
//   	Section: &EvaluationFormSectionProperty{
//   		Instructions: jsii.String("instructions"),
//   		Items: []interface{}{
//   			&EvaluationFormItemProperty{
//   				Question: &EvaluationFormQuestionProperty{
//   					Enablement: &EvaluationFormItemEnablementConfigurationProperty{
//   						Action: jsii.String("action"),
//   						Condition: &EvaluationFormItemEnablementConditionProperty{
//   							Operands: []interface{}{
//   								&EvaluationFormItemEnablementConditionOperandProperty{
//   									Expression: &EvaluationFormItemEnablementExpressionProperty{
//   										Comparator: jsii.String("comparator"),
//   										Source: &EvaluationFormItemEnablementSourceProperty{
//   											RefId: jsii.String("refId"),
//   											Type: jsii.String("type"),
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
//   							Operator: jsii.String("operator"),
//   						},
//   						DefaultAction: jsii.String("defaultAction"),
//   					},
//   					Instructions: jsii.String("instructions"),
//   					NotApplicableEnabled: jsii.Boolean(false),
//   					QuestionType: jsii.String("questionType"),
//   					QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   						Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   							Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   								AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   									SourceType: jsii.String("sourceType"),
//   								},
//   								PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   									Label: jsii.String("label"),
//   								},
//   							},
//   							MaxValue: jsii.Number(123),
//   							MinValue: jsii.Number(123),
//   							Options: []interface{}{
//   								&EvaluationFormNumericQuestionOptionProperty{
//   									AutomaticFail: jsii.Boolean(false),
//   									AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   										TargetSection: jsii.String("targetSection"),
//   									},
//   									MaxValue: jsii.Number(123),
//   									MinValue: jsii.Number(123),
//   									Score: jsii.Number(123),
//   								},
//   							},
//   						},
//   						SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   							Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   								AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   									SourceType: jsii.String("sourceType"),
//   								},
//   								DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   								Options: []interface{}{
//   									&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   										RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   											Category: jsii.String("category"),
//   											Condition: jsii.String("condition"),
//   											OptionRefId: jsii.String("optionRefId"),
//   										},
//   									},
//   								},
//   							},
//   							DisplayAs: jsii.String("displayAs"),
//   							Options: []interface{}{
//   								&EvaluationFormSingleSelectQuestionOptionProperty{
//   									AutomaticFail: jsii.Boolean(false),
//   									AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   										TargetSection: jsii.String("targetSection"),
//   									},
//   									RefId: jsii.String("refId"),
//   									Score: jsii.Number(123),
//   									Text: jsii.String("text"),
//   								},
//   							},
//   						},
//   						Text: &EvaluationFormTextQuestionPropertiesProperty{
//   							Automation: &EvaluationFormTextQuestionAutomationProperty{
//   								AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   									SourceType: jsii.String("sourceType"),
//   								},
//   							},
//   						},
//   					},
//   					RefId: jsii.String("refId"),
//   					Title: jsii.String("title"),
//   					Weight: jsii.Number(123),
//   				},
//   				Section: evaluationFormSectionProperty_,
//   			},
//   		},
//   		RefId: jsii.String("refId"),
//   		Title: jsii.String("title"),
//   		Weight: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformbaseitem.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormBaseItemProperty struct {
	// A subsection or inner section of an item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformbaseitem.html#cfn-connect-evaluationform-evaluationformbaseitem-section
	//
	Section interface{} `field:"optional" json:"section" yaml:"section"`
}

