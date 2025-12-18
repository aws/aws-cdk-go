package awsconnect


// Automation rule for multi-select questions based on rule categories.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiSelectQuestionRuleCategoryAutomationProperty := &MultiSelectQuestionRuleCategoryAutomationProperty{
//   	Category: jsii.String("category"),
//   	Condition: jsii.String("condition"),
//   	OptionRefIds: []*string{
//   		jsii.String("optionRefIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-multiselectquestionrulecategoryautomation.html
//
type CfnEvaluationForm_MultiSelectQuestionRuleCategoryAutomationProperty struct {
	// The category name for this automation rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-multiselectquestionrulecategoryautomation.html#cfn-connect-evaluationform-multiselectquestionrulecategoryautomation-category
	//
	Category *string `field:"required" json:"category" yaml:"category"`
	// The condition for this automation rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-multiselectquestionrulecategoryautomation.html#cfn-connect-evaluationform-multiselectquestionrulecategoryautomation-condition
	//
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// Reference IDs of options for this automation rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-multiselectquestionrulecategoryautomation.html#cfn-connect-evaluationform-multiselectquestionrulecategoryautomation-optionrefids
	//
	OptionRefIds *[]*string `field:"required" json:"optionRefIds" yaml:"optionRefIds"`
}

