package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormQuestionTypePropertiesProperty := &EvaluationFormQuestionTypePropertiesProperty{
//   	Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   		MaxValue: jsii.Number(123),
//   		MinValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   			PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   				Label: jsii.String("label"),
//   			},
//   		},
//   		Options: []interface{}{
//   			&EvaluationFormNumericQuestionOptionProperty{
//   				MaxValue: jsii.Number(123),
//   				MinValue: jsii.Number(123),
//
//   				// the properties below are optional
//   				AutomaticFail: jsii.Boolean(false),
//   				Score: jsii.Number(123),
//   			},
//   		},
//   	},
//   	SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   		Options: []interface{}{
//   			&EvaluationFormSingleSelectQuestionOptionProperty{
//   				RefId: jsii.String("refId"),
//   				Text: jsii.String("text"),
//
//   				// the properties below are optional
//   				AutomaticFail: jsii.Boolean(false),
//   				Score: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   			Options: []interface{}{
//   				&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   					RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   						Category: jsii.String("category"),
//   						Condition: jsii.String("condition"),
//   						OptionRefId: jsii.String("optionRefId"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   		},
//   		DisplayAs: jsii.String("displayAs"),
//   	},
//   }
//
type CfnEvaluationForm_EvaluationFormQuestionTypePropertiesProperty struct {
	// `CfnEvaluationForm.EvaluationFormQuestionTypePropertiesProperty.Numeric`.
	Numeric interface{} `field:"optional" json:"numeric" yaml:"numeric"`
	// `CfnEvaluationForm.EvaluationFormQuestionTypePropertiesProperty.SingleSelect`.
	SingleSelect interface{} `field:"optional" json:"singleSelect" yaml:"singleSelect"`
}

