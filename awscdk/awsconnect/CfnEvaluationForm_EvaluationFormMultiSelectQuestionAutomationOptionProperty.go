package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormMultiSelectQuestionAutomationOptionProperty := &EvaluationFormMultiSelectQuestionAutomationOptionProperty{
//   	RuleCategory: &MultiSelectQuestionRuleCategoryAutomationProperty{
//   		Category: jsii.String("category"),
//   		Condition: jsii.String("condition"),
//   		OptionRefIds: []*string{
//   			jsii.String("optionRefIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionautomationoption.html
//
type CfnEvaluationForm_EvaluationFormMultiSelectQuestionAutomationOptionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionautomationoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionautomationoption-rulecategory
	//
	RuleCategory interface{} `field:"required" json:"ruleCategory" yaml:"ruleCategory"`
}

