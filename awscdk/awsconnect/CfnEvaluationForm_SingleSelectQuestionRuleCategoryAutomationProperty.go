package awsconnect


// Information about the automation option based on a rule category for a single select question.
//
// *Length Constraints* : Minimum length of 1. Maximum length of 50.
//
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
	// The category name, as defined in Rules.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 50.
	Category *string `field:"required" json:"category" yaml:"category"`
	// The condition to apply for the automation option.
	//
	// If the condition is PRESENT, then the option is applied when the contact data includes the category. Similarly, if the condition is NOT_PRESENT, then the option is applied when the contact data does not include the category.
	//
	// *Allowed values* : `PRESENT` | `NOT_PRESENT`
	//
	// *Maximum* : 50.
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// The identifier of the answer option. An identifier must be unique within the question.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 40.
	OptionRefId *string `field:"required" json:"optionRefId" yaml:"optionRefId"`
}

