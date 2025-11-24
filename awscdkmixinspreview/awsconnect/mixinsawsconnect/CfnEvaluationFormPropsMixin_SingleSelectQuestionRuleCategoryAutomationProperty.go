package mixinsawsconnect


// Information about the automation option based on a rule category for a single select question.
//
// *Length Constraints* : Minimum length of 1. Maximum length of 50.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   singleSelectQuestionRuleCategoryAutomationProperty := &SingleSelectQuestionRuleCategoryAutomationProperty{
//   	Category: jsii.String("category"),
//   	Condition: jsii.String("condition"),
//   	OptionRefId: jsii.String("optionRefId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-singleselectquestionrulecategoryautomation.html
//
type CfnEvaluationFormPropsMixin_SingleSelectQuestionRuleCategoryAutomationProperty struct {
	// The category name, as defined in Rules.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 50.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-singleselectquestionrulecategoryautomation.html#cfn-connect-evaluationform-singleselectquestionrulecategoryautomation-category
	//
	Category *string `field:"optional" json:"category" yaml:"category"`
	// The condition to apply for the automation option.
	//
	// If the condition is PRESENT, then the option is applied when the contact data includes the category. Similarly, if the condition is NOT_PRESENT, then the option is applied when the contact data does not include the category.
	//
	// *Allowed values* : `PRESENT` | `NOT_PRESENT`
	//
	// *Maximum* : 50.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-singleselectquestionrulecategoryautomation.html#cfn-connect-evaluationform-singleselectquestionrulecategoryautomation-condition
	//
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The identifier of the answer option. An identifier must be unique within the question.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 40.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-singleselectquestionrulecategoryautomation.html#cfn-connect-evaluationform-singleselectquestionrulecategoryautomation-optionrefid
	//
	OptionRefId *string `field:"optional" json:"optionRefId" yaml:"optionRefId"`
}

