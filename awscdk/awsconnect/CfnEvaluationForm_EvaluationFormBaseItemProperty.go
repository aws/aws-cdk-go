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
//   var evaluationFormSectionProperty_ evaluationFormSectionProperty
//
//   evaluationFormBaseItemProperty := &EvaluationFormBaseItemProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformbaseitem.html
//
type CfnEvaluationForm_EvaluationFormBaseItemProperty struct {
	// A subsection or inner section of an item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformbaseitem.html#cfn-connect-evaluationform-evaluationformbaseitem-section
	//
	Section interface{} `field:"required" json:"section" yaml:"section"`
}

