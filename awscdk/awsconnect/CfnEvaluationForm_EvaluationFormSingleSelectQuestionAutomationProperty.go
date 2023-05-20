package awsconnect


// Information about the automation configuration in single select questions.
//
// Automation options are evaluated in order, and the first matched option is applied. If no automation option matches, and there is a default option, then the default option is applied.
//
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
	// The automation options of the single select question.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 20.
	Options interface{} `field:"required" json:"options" yaml:"options"`
	// The identifier of the default answer option, when none of the automation options match the criteria.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 40.
	DefaultOptionRefId *string `field:"optional" json:"defaultOptionRefId" yaml:"defaultOptionRefId"`
}

