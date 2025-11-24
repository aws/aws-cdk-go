package mixinsawsconnect


// Information about a section from an evaluation form.
//
// A section can contain sections and/or questions. Evaluation forms can only contain sections and subsections (two level nesting).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var evaluationFormSectionProperty_ EvaluationFormSectionProperty
//
//   evaluationFormSectionProperty := &EvaluationFormSectionProperty{
//   	Instructions: jsii.String("instructions"),
//   	Items: []interface{}{
//   		&EvaluationFormItemProperty{
//   			Question: &EvaluationFormQuestionProperty{
//   				Enablement: &EvaluationFormItemEnablementConfigurationProperty{
//   					Action: jsii.String("action"),
//   					Condition: &EvaluationFormItemEnablementConditionProperty{
//   						Operands: []interface{}{
//   							&EvaluationFormItemEnablementConditionOperandProperty{
//   								Expression: &EvaluationFormItemEnablementExpressionProperty{
//   									Comparator: jsii.String("comparator"),
//   									Source: &EvaluationFormItemEnablementSourceProperty{
//   										RefId: jsii.String("refId"),
//   										Type: jsii.String("type"),
//   									},
//   									Values: []interface{}{
//   										&EvaluationFormItemEnablementSourceValueProperty{
//   											RefId: jsii.String("refId"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   						Operator: jsii.String("operator"),
//   					},
//   					DefaultAction: jsii.String("defaultAction"),
//   				},
//   				Instructions: jsii.String("instructions"),
//   				NotApplicableEnabled: jsii.Boolean(false),
//   				QuestionType: jsii.String("questionType"),
//   				QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   					Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   						Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   							AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   								SourceType: jsii.String("sourceType"),
//   							},
//   							PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   								Label: jsii.String("label"),
//   							},
//   						},
//   						MaxValue: jsii.Number(123),
//   						MinValue: jsii.Number(123),
//   						Options: []interface{}{
//   							&EvaluationFormNumericQuestionOptionProperty{
//   								AutomaticFail: jsii.Boolean(false),
//   								AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   									TargetSection: jsii.String("targetSection"),
//   								},
//   								MaxValue: jsii.Number(123),
//   								MinValue: jsii.Number(123),
//   								Score: jsii.Number(123),
//   							},
//   						},
//   					},
//   					SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   						Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   							AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   								SourceType: jsii.String("sourceType"),
//   							},
//   							DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   							Options: []interface{}{
//   								&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   									RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   										Category: jsii.String("category"),
//   										Condition: jsii.String("condition"),
//   										OptionRefId: jsii.String("optionRefId"),
//   									},
//   								},
//   							},
//   						},
//   						DisplayAs: jsii.String("displayAs"),
//   						Options: []interface{}{
//   							&EvaluationFormSingleSelectQuestionOptionProperty{
//   								AutomaticFail: jsii.Boolean(false),
//   								AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   									TargetSection: jsii.String("targetSection"),
//   								},
//   								RefId: jsii.String("refId"),
//   								Score: jsii.Number(123),
//   								Text: jsii.String("text"),
//   							},
//   						},
//   					},
//   					Text: &EvaluationFormTextQuestionPropertiesProperty{
//   						Automation: &EvaluationFormTextQuestionAutomationProperty{
//   							AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   								SourceType: jsii.String("sourceType"),
//   							},
//   						},
//   					},
//   				},
//   				RefId: jsii.String("refId"),
//   				Title: jsii.String("title"),
//   				Weight: jsii.Number(123),
//   			},
//   			Section: evaluationFormSectionProperty_,
//   		},
//   	},
//   	RefId: jsii.String("refId"),
//   	Title: jsii.String("title"),
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsection.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormSectionProperty struct {
	// The instructions of the section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsection.html#cfn-connect-evaluationform-evaluationformsection-instructions
	//
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// The items of the section.
	//
	// *Minimum* : 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsection.html#cfn-connect-evaluationform-evaluationformsection-items
	//
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// The identifier of the section. An identifier must be unique within the evaluation form.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 40.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsection.html#cfn-connect-evaluationform-evaluationformsection-refid
	//
	RefId *string `field:"optional" json:"refId" yaml:"refId"`
	// The title of the section.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 128.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsection.html#cfn-connect-evaluationform-evaluationformsection-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The scoring weight of the section.
	//
	// *Minimum* : 0
	//
	// *Maximum* : 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformsection.html#cfn-connect-evaluationform-evaluationformsection-weight
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

