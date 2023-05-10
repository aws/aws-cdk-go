package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnEvaluationForm`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var evaluationFormSectionProperty_ evaluationFormSectionProperty
//
//   cfnEvaluationFormProps := &CfnEvaluationFormProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	Items: []interface{}{
//   		&EvaluationFormBaseItemProperty{
//   			Section: &evaluationFormSectionProperty{
//   				RefId: jsii.String("refId"),
//   				Title: jsii.String("title"),
//
//   				// the properties below are optional
//   				Instructions: jsii.String("instructions"),
//   				Items: []interface{}{
//   					&EvaluationFormItemProperty{
//   						Question: &EvaluationFormQuestionProperty{
//   							QuestionType: jsii.String("questionType"),
//   							RefId: jsii.String("refId"),
//   							Title: jsii.String("title"),
//
//   							// the properties below are optional
//   							Instructions: jsii.String("instructions"),
//   							NotApplicableEnabled: jsii.Boolean(false),
//   							QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   								Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   									MaxValue: jsii.Number(123),
//   									MinValue: jsii.Number(123),
//
//   									// the properties below are optional
//   									Automation: &EvaluationFormNumericQuestionAutomationProperty{
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
//   											Score: jsii.Number(123),
//   										},
//   									},
//
//   									// the properties below are optional
//   									Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   										Options: []interface{}{
//   											&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   												RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   													Category: jsii.String("category"),
//   													Condition: jsii.String("condition"),
//   													OptionRefId: jsii.String("optionRefId"),
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   									},
//   									DisplayAs: jsii.String("displayAs"),
//   								},
//   							},
//   							Weight: jsii.Number(123),
//   						},
//   						Section: evaluationFormSectionProperty_,
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
//   	Description: jsii.String("description"),
//   	ScoringStrategy: &ScoringStrategyProperty{
//   		Mode: jsii.String("mode"),
//   		Status: jsii.String("status"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEvaluationFormProps struct {
	// `AWS::Connect::EvaluationForm.InstanceArn`.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// `AWS::Connect::EvaluationForm.Items`.
	Items interface{} `field:"required" json:"items" yaml:"items"`
	// `AWS::Connect::EvaluationForm.Status`.
	Status *string `field:"required" json:"status" yaml:"status"`
	// `AWS::Connect::EvaluationForm.Title`.
	Title *string `field:"required" json:"title" yaml:"title"`
	// `AWS::Connect::EvaluationForm.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Connect::EvaluationForm.ScoringStrategy`.
	ScoringStrategy interface{} `field:"optional" json:"scoringStrategy" yaml:"scoringStrategy"`
	// `AWS::Connect::EvaluationForm.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

