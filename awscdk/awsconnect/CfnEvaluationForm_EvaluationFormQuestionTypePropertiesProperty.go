package awsconnect


// Information about properties for a question in an evaluation form.
//
// The question type properties must be either for a numeric question or a single select question.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormQuestionTypePropertiesProperty := &EvaluationFormQuestionTypePropertiesProperty{
//   	Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   		MaxValue: jsii.Number(123),
//   		MinValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   			PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   				Label: jsii.String("label"),
//   			},
//   		},
//   		Options: []interface{}{
//   			&EvaluationFormNumericQuestionOptionProperty{
//   				MaxValue: jsii.Number(123),
//   				MinValue: jsii.Number(123),
//
//   				// the properties below are optional
//   				AutomaticFail: jsii.Boolean(false),
//   				Score: jsii.Number(123),
//   			},
//   		},
//   	},
//   	SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   		Options: []interface{}{
//   			&EvaluationFormSingleSelectQuestionOptionProperty{
//   				RefId: jsii.String("refId"),
//   				Text: jsii.String("text"),
//
//   				// the properties below are optional
//   				AutomaticFail: jsii.Boolean(false),
//   				Score: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   			Options: []interface{}{
//   				&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   					RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   						Category: jsii.String("category"),
//   						Condition: jsii.String("condition"),
//   						OptionRefId: jsii.String("optionRefId"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   		},
//   		DisplayAs: jsii.String("displayAs"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestiontypeproperties.html
//
type CfnEvaluationForm_EvaluationFormQuestionTypePropertiesProperty struct {
	// The properties of the numeric question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestiontypeproperties.html#cfn-connect-evaluationform-evaluationformquestiontypeproperties-numeric
	//
	Numeric interface{} `field:"optional" json:"numeric" yaml:"numeric"`
	// The properties of the numeric question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestiontypeproperties.html#cfn-connect-evaluationform-evaluationformquestiontypeproperties-singleselect
	//
	SingleSelect interface{} `field:"optional" json:"singleSelect" yaml:"singleSelect"`
}

