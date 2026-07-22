package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEvaluationForm`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var evaluationFormSectionProperty_ EvaluationFormSectionProperty
//
//   cfnEvaluationFormProps := &CfnEvaluationFormProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	Items: []interface{}{
//   		&EvaluationFormBaseItemProperty{
//   			Section: &EvaluationFormSectionProperty{
//   				RefId: jsii.String("refId"),
//   				Title: jsii.String("title"),
//
//   				// the properties below are optional
//   				Instructions: jsii.String("instructions"),
//   				IsExcludedFromScoring: jsii.Boolean(false),
//   				Items: []interface{}{
//   					&EvaluationFormItemProperty{
//   						Question: &EvaluationFormQuestionProperty{
//   							QuestionType: jsii.String("questionType"),
//   							RefId: jsii.String("refId"),
//   							Title: jsii.String("title"),
//
//   							// the properties below are optional
//   							Enablement: &EvaluationFormItemEnablementConfigurationProperty{
//   								Action: jsii.String("action"),
//   								Condition: &EvaluationFormItemEnablementConditionProperty{
//   									Operands: []interface{}{
//   										&EvaluationFormItemEnablementConditionOperandProperty{
//   											Expression: &EvaluationFormItemEnablementExpressionProperty{
//   												Comparator: jsii.String("comparator"),
//   												Source: &EvaluationFormItemEnablementSourceProperty{
//   													Type: jsii.String("type"),
//
//   													// the properties below are optional
//   													RefId: jsii.String("refId"),
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
//
//   									// the properties below are optional
//   									Operator: jsii.String("operator"),
//   								},
//
//   								// the properties below are optional
//   								DefaultAction: jsii.String("defaultAction"),
//   							},
//   							Instructions: jsii.String("instructions"),
//   							NotApplicableEnabled: jsii.Boolean(false),
//   							QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   								MultiSelect: &EvaluationFormMultiSelectQuestionPropertiesProperty{
//   									Options: []interface{}{
//   										&EvaluationFormMultiSelectQuestionOptionProperty{
//   											RefId: jsii.String("refId"),
//   											Text: jsii.String("text"),
//
//   											// the properties below are optional
//   											AutomaticFail: jsii.Boolean(false),
//   											AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   												TargetSection: jsii.String("targetSection"),
//   											},
//   											PointsConfiguration: &QuestionOptionPointsConfigurationProperty{
//   												PointValue: jsii.Number(123),
//
//   												// the properties below are optional
//   												IsBonus: jsii.Boolean(false),
//   											},
//   											Score: jsii.Number(123),
//   										},
//   									},
//
//   									// the properties below are optional
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
//   								},
//   								Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   									MaxValue: jsii.Number(123),
//   									MinValue: jsii.Number(123),
//
//   									// the properties below are optional
//   									Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   										AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   											SourceType: jsii.String("sourceType"),
//   										},
//   										PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   											Label: jsii.String("label"),
//   										},
//   									},
//   									Options: []interface{}{
//   										&EvaluationFormNumericQuestionOptionProperty{
//   											MaxValue: jsii.Number(123),
//   											MinValue: jsii.Number(123),
//
//   											// the properties below are optional
//   											AutomaticFail: jsii.Boolean(false),
//   											AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   												TargetSection: jsii.String("targetSection"),
//   											},
//   											PointsConfiguration: &QuestionOptionPointsConfigurationProperty{
//   												PointValue: jsii.Number(123),
//
//   												// the properties below are optional
//   												IsBonus: jsii.Boolean(false),
//   											},
//   											Score: jsii.Number(123),
//   										},
//   									},
//   								},
//   								SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   									Options: []interface{}{
//   										&EvaluationFormSingleSelectQuestionOptionProperty{
//   											RefId: jsii.String("refId"),
//   											Text: jsii.String("text"),
//
//   											// the properties below are optional
//   											AutomaticFail: jsii.Boolean(false),
//   											AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   												TargetSection: jsii.String("targetSection"),
//   											},
//   											PointsConfiguration: &QuestionOptionPointsConfigurationProperty{
//   												PointValue: jsii.Number(123),
//
//   												// the properties below are optional
//   												IsBonus: jsii.Boolean(false),
//   											},
//   											Score: jsii.Number(123),
//   										},
//   									},
//
//   									// the properties below are optional
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
//   								},
//   								Text: &EvaluationFormTextQuestionPropertiesProperty{
//   									Automation: &EvaluationFormTextQuestionAutomationProperty{
//   										AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   											SourceType: jsii.String("sourceType"),
//   										},
//   									},
//   								},
//   							},
//   							ScoringConfiguration: &EvaluationFormQuestionScoringConfigurationProperty{
//   								IsExcludedFromScoring: jsii.Boolean(false),
//   								PointsConfiguration: &QuestionPointsConfigurationProperty{
//   									IsBonus: jsii.Boolean(false),
//   									MaxPointValue: jsii.Number(123),
//   									MinPointValue: jsii.Number(123),
//   								},
//   								ScoreThresholds: []interface{}{
//   									&EvaluationFormScoreThresholdProperty{
//   										PerformanceCategory: jsii.String("performanceCategory"),
//
//   										// the properties below are optional
//   										MaxScorePercentage: jsii.Number(123),
//   										MinScorePercentage: jsii.Number(123),
//   									},
//   								},
//   							},
//   							Weight: jsii.Number(123),
//   						},
//   						Section: evaluationFormSectionProperty_,
//   					},
//   				},
//   				ScoreThresholds: []interface{}{
//   					&EvaluationFormScoreThresholdProperty{
//   						PerformanceCategory: jsii.String("performanceCategory"),
//
//   						// the properties below are optional
//   						MaxScorePercentage: jsii.Number(123),
//   						MinScorePercentage: jsii.Number(123),
//   					},
//   				},
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Status: jsii.String("status"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	AutoEvaluationConfiguration: &AutoEvaluationConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Description: jsii.String("description"),
//   	LanguageConfiguration: &EvaluationFormLanguageConfigurationProperty{
//   		FormLanguage: jsii.String("formLanguage"),
//   	},
//   	ReviewConfiguration: &EvaluationReviewConfigurationProperty{
//   		ReviewNotificationRecipients: []interface{}{
//   			&EvaluationReviewNotificationRecipientProperty{
//   				Type: jsii.String("type"),
//   				Value: &EvaluationReviewNotificationRecipientValueProperty{
//   					UserId: jsii.String("userId"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		EligibilityDays: jsii.Number(123),
//   	},
//   	ScoringStrategy: &ScoringStrategyProperty{
//   		Mode: jsii.String("mode"),
//   		Status: jsii.String("status"),
//
//   		// the properties below are optional
//   		ScoreThresholds: []interface{}{
//   			&EvaluationFormScoreThresholdProperty{
//   				PerformanceCategory: jsii.String("performanceCategory"),
//
//   				// the properties below are optional
//   				MaxScorePercentage: jsii.Number(123),
//   				MinScorePercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetConfiguration: &EvaluationFormTargetConfigurationProperty{
//   		ContactInteractionType: jsii.String("contactInteractionType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html
//
type CfnEvaluationFormProps struct {
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-instancearn
	//
	InstanceArn interface{} `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Items that are part of the evaluation form.
	//
	// The total number of sections and questions must not exceed 100 each. Questions must be contained in a section.
	//
	// *Minimum size* : 1
	//
	// *Maximum size* : 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-items
	//
	Items interface{} `field:"required" json:"items" yaml:"items"`
	// The status of the evaluation form.
	//
	// *Allowed values* : `DRAFT` | `ACTIVE`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-status
	//
	// Default: - "DRAFT".
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// A title of the evaluation form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-title
	//
	Title *string `field:"required" json:"title" yaml:"title"`
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
	// Configuration for language settings of this evaluation form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-languageconfiguration
	//
	LanguageConfiguration interface{} `field:"optional" json:"languageConfiguration" yaml:"languageConfiguration"`
	// Configuration settings for evaluation reviews.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-reviewconfiguration
	//
	ReviewConfiguration interface{} `field:"optional" json:"reviewConfiguration" yaml:"reviewConfiguration"`
	// A scoring strategy of the evaluation form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-scoringstrategy
	//
	ScoringStrategy interface{} `field:"optional" json:"scoringStrategy" yaml:"scoringStrategy"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Configuration that specifies the target for this evaluation form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-targetconfiguration
	//
	TargetConfiguration interface{} `field:"optional" json:"targetConfiguration" yaml:"targetConfiguration"`
}

