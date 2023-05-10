package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormSingleSelectQuestionAutomationOptionProperty := &EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   	RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   		Category: jsii.String("category"),
//   		Condition: jsii.String("condition"),
//   		OptionRefId: jsii.String("optionRefId"),
//   	},
//   }
//
type CfnEvaluationForm_EvaluationFormSingleSelectQuestionAutomationOptionProperty struct {
	// `CfnEvaluationForm.EvaluationFormSingleSelectQuestionAutomationOptionProperty.RuleCategory`.
	RuleCategory interface{} `field:"required" json:"ruleCategory" yaml:"ruleCategory"`
}

