package awsconnect


// The automation options of the single select question.
//
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
	// The automation option based on a rule category for the single select question.
	RuleCategory interface{} `field:"required" json:"ruleCategory" yaml:"ruleCategory"`
}

