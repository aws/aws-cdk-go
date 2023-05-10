package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormSingleSelectQuestionAutomationProperty := &EvaluationFormSingleSelectQuestionAutomationProperty{
//   	Options: []interface{}{
//   		&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   			RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   				Category: jsii.String("category"),
//   				Condition: jsii.String("condition"),
//   				OptionRefId: jsii.String("optionRefId"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   }
//
type CfnEvaluationForm_EvaluationFormSingleSelectQuestionAutomationProperty struct {
	// `CfnEvaluationForm.EvaluationFormSingleSelectQuestionAutomationProperty.Options`.
	Options interface{} `field:"required" json:"options" yaml:"options"`
	// `CfnEvaluationForm.EvaluationFormSingleSelectQuestionAutomationProperty.DefaultOptionRefId`.
	DefaultOptionRefId *string `field:"optional" json:"defaultOptionRefId" yaml:"defaultOptionRefId"`
}

