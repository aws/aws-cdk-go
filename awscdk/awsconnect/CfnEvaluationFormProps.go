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
//   	AutoEvaluationConfiguration: &AutoEvaluationConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html
//
type CfnEvaluationFormProps struct {
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-autoevaluationconfiguration
	//
	AutoEvaluationConfiguration interface{} `field:"optional" json:"autoEvaluationConfiguration" yaml:"autoEvaluationConfiguration"`
	// The description of the evaluation form.
	//
	// *Length Constraints* : Minimum length of 0. Maximum length of 1024.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-evaluationform.html#cfn-connect-evaluationform-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
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
}

