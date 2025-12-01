package previewawsconnectmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnEvaluationFormPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var evaluationFormSectionProperty_ EvaluationFormSectionProperty
//
//   cfnEvaluationFormMixinProps := &CfnEvaluationFormMixinProps{
//   	AutoEvaluationConfiguration: &AutoEvaluationConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Description: jsii.String("description"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Items: []interface{}{
//   		&EvaluationFormBaseItemProperty{
//   			Section: &EvaluationFormSectionProperty{
//   				Instructions: jsii.String("instructions"),
//   				Items: []interface{}{
//   					&EvaluationFormItemProperty{
//   						Question: &EvaluationFormQuestionProperty{
//   							Enablement: &EvaluationFormItemEnablementConfigurationProperty{
//   								Action: jsii.String("action"),
//   								Condition: &EvaluationFormItemEnablementConditionProperty{
//   									Operands: []interface{}{
//   										&EvaluationFormItemEnablementConditionOperandProperty{
//   											Expression: &EvaluationFormItemEnablementExpressionProperty{
//   												Comparator: jsii.String("comparator"),
//   												Source: &EvaluationFormItemEnablementSourceProperty{
//   													RefId: jsii.String("refId"),
//   													Type: jsii.String("type"),
//   												},
//   												Values: []interface{}{
//   													&EvaluationFormItemEnablementSourceValueProperty{
//   														RefId: jsii.String("refId"),
//   														Type: jsii.String("type"),
//   													},
//   												},
//   											},
//   										},
//   									},
//   									Operator: jsii.String("operator"),
//   								},
//   								DefaultAction: jsii.String("defaultAction"),
//   							},
//   							Instructions: jsii.String("instructions"),
//   							NotApplicableEnabled: jsii.Boolean(false),
//   							QuestionType: jsii.String("questionType"),
//   							QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   								MultiSelect: &EvaluationFormMultiSelectQuestionPropertiesProperty{
//   									Automation: &EvaluationFormMultiSelectQuestionAutomationProperty{
//   										AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   											SourceType: jsii.String("sourceType"),
//   										},
//   										DefaultOptionRefIds: []*string{
//   											jsii.String("defaultOptionRefIds"),
//   										},
//   										Options: []interface{}{
//   											&EvaluationFormMultiSelectQuestionAutomationOptionProperty{
//   												RuleCategory: &MultiSelectQuestionRuleCategoryAutomationProperty{
//   													Category: jsii.String("category"),
//   													Condition: jsii.String("condition"),
//   													OptionRefIds: []*string{
//   														jsii.String("optionRefIds"),
//   													},
//   												},
//   											},
//   										},
//   									},
//   									DisplayAs: jsii.String("displayAs"),
//   									Options: []interface{}{
//   										&EvaluationFormMultiSelectQuestionOptionProperty{
//   											RefId: jsii.String("refId"),
//   											Text: jsii.String("text"),
//   										},
//   									},
//   								},
//   								Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   									Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   										AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   											SourceType: jsii.String("sourceType"),
//   										},
//   										PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   											Label: jsii.String("label"),
//   										},
//   									},
//   									MaxValue: jsii.Number(123),
//   									MinValue: jsii.Number(123),
//   									Options: []interface{}{
//   										&EvaluationFormNumericQuestionOptionProperty{
//   											AutomaticFail: jsii.Boolean(false),
//   											AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   												TargetSection: jsii.String("targetSection"),
//   											},
//   											MaxValue: jsii.Number(123),
//   											MinValue: jsii.Number(123),
//   											Score: jsii.Number(123),
//   										},
//   									},
//   								},
//   								SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   									Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   										AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   											SourceType: jsii.String("sourceType"),
//   										},
//   										DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   										Options: []interface{}{
//   											&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   												RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   													Category: jsii.String("category"),
//   													Condition: jsii.String("condition"),
//   													OptionRefId: jsii.String("optionRefId"),
//   												},
//   											},
//   										},
//   									},
//   									DisplayAs: jsii.String("displayAs"),
//   									Options: []interface{}{
//   										&EvaluationFormSingleSelectQuestionOptionProperty{
//   											AutomaticFail: jsii.Boolean(false),
//   											AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   												TargetSection: jsii.String("targetSection"),
//   											},
//   											RefId: jsii.String("refId"),
//   											Score: jsii.Number(123),
//   											Text: jsii.String("text"),
//   										},
//   									},
//   								},
//   								Text: &EvaluationFormTextQuestionPropertiesProperty{
//   									Automation: &EvaluationFormTextQuestionAutomationProperty{
//   										AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   											SourceType: jsii.String("sourceType"),
//   										},
//   									},
//   								},
//   							},
//   							RefId: jsii.String("refId"),
//   							Title: jsii.String("title"),
//   							Weight: jsii.Number(123),
//   						},
//   						Section: evaluationFormSectionProperty_,
//   					},
//   				},
//   				RefId: jsii.String("refId"),
//   				Title: jsii.String("title"),
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   	},
//   	LanguageConfiguration: &EvaluationFormLanguageConfigurationProperty{
//   		FormLanguage: jsii.String("formLanguage"),
//   	},
//   	ScoringStrategy: &ScoringStrategyProperty{
//   		Mode: jsii.String("mode"),
//   		Status: jsii.String("status"),
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetConfiguration: &EvaluationFormTargetConfigurationProperty{
//   		ContactInteractionType: jsii.String("contactInteractionType"),
//   	},
//   	Title: jsii.String("title"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html
//
type CfnEvaluationFormMixinProps struct {
	// The automatic evaluation configuration of an evaluation form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-autoevaluationconfiguration
	//
	AutoEvaluationConfiguration interface{} `field:"optional" json:"autoEvaluationConfiguration" yaml:"autoEvaluationConfiguration"`
	// The description of the evaluation form.
	//
	// *Length Constraints* : Minimum length of 0. Maximum length of 1024.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// Items that are part of the evaluation form.
	//
	// The total number of sections and questions must not exceed 100 each. Questions must be contained in a section.
	//
	// *Minimum size* : 1
	//
	// *Maximum size* : 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-items
	//
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-languageconfiguration
	//
	LanguageConfiguration interface{} `field:"optional" json:"languageConfiguration" yaml:"languageConfiguration"`
	// A scoring strategy of the evaluation form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-scoringstrategy
	//
	ScoringStrategy interface{} `field:"optional" json:"scoringStrategy" yaml:"scoringStrategy"`
	// The status of the evaluation form.
	//
	// *Allowed values* : `DRAFT` | `ACTIVE`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-status
	//
	// Default: - "DRAFT".
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-targetconfiguration
	//
	TargetConfiguration interface{} `field:"optional" json:"targetConfiguration" yaml:"targetConfiguration"`
	// A title of the evaluation form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
}

