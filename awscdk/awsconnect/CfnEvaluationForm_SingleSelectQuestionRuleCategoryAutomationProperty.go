package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   singleSelectQuestionRuleCategoryAutomationProperty := &SingleSelectQuestionRuleCategoryAutomationProperty{
//   	Category: jsii.String("category"),
//   	Condition: jsii.String("condition"),
//   	OptionRefId: jsii.String("optionRefId"),
//   }
//
type CfnEvaluationForm_SingleSelectQuestionRuleCategoryAutomationProperty struct {
	// `CfnEvaluationForm.SingleSelectQuestionRuleCategoryAutomationProperty.Category`.
	Category *string `field:"required" json:"category" yaml:"category"`
	// `CfnEvaluationForm.SingleSelectQuestionRuleCategoryAutomationProperty.Condition`.
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// `CfnEvaluationForm.SingleSelectQuestionRuleCategoryAutomationProperty.OptionRefId`.
	OptionRefId *string `field:"required" json:"optionRefId" yaml:"optionRefId"`
}

