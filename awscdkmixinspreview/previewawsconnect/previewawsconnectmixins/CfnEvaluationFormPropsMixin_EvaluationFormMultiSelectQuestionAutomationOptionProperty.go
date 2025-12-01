package previewawsconnectmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnEvaluationFormPropsMixin_EvaluationFormMultiSelectQuestionAutomationOptionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionautomationoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionautomationoption-rulecategory
	//
	RuleCategory interface{} `field:"optional" json:"ruleCategory" yaml:"ruleCategory"`
}

