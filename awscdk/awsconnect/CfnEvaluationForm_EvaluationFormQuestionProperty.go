package awsconnect


// Information about a question from an evaluation form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormQuestionProperty := &EvaluationFormQuestionProperty{
//   	QuestionType: jsii.String("questionType"),
//   	RefId: jsii.String("refId"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	Instructions: jsii.String("instructions"),
//   	NotApplicableEnabled: jsii.Boolean(false),
//   	QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   		Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   			MaxValue: jsii.Number(123),
//   			MinValue: jsii.Number(123),
//
//   			// the properties below are optional
//   			Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   				PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   					Label: jsii.String("label"),
//   				},
//   			},
//   			Options: []interface{}{
//   				&EvaluationFormNumericQuestionOptionProperty{
//   					MaxValue: jsii.Number(123),
//   					MinValue: jsii.Number(123),
//
//   					// the properties below are optional
//   					AutomaticFail: jsii.Boolean(false),
//   					Score: jsii.Number(123),
//   				},
//   			},
//   		},
//   		SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   			Options: []interface{}{
//   				&EvaluationFormSingleSelectQuestionOptionProperty{
//   					RefId: jsii.String("refId"),
//   					Text: jsii.String("text"),
//
//   					// the properties below are optional
//   					AutomaticFail: jsii.Boolean(false),
//   					Score: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   				Options: []interface{}{
//   					&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   						RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   							Category: jsii.String("category"),
//   							Condition: jsii.String("condition"),
//   							OptionRefId: jsii.String("optionRefId"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   			},
//   			DisplayAs: jsii.String("displayAs"),
//   		},
//   	},
//   	Weight: jsii.Number(123),
//   }
//
type CfnEvaluationForm_EvaluationFormQuestionProperty struct {
	// The type of the question.
	//
	// *Allowed values* : `NUMERIC` | `SINGLESELECT` | `TEXT`.
	QuestionType *string `field:"required" json:"questionType" yaml:"questionType"`
	// The identifier of the question. An identifier must be unique within the evaluation form.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 40.
	RefId *string `field:"required" json:"refId" yaml:"refId"`
	// The title of the question.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 350.
	Title *string `field:"required" json:"title" yaml:"title"`
	// The instructions of the section.
	//
	// *Length Constraints* : Minimum length of 0. Maximum length of 1024.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// The flag to enable not applicable answers to the question.
	NotApplicableEnabled interface{} `field:"optional" json:"notApplicableEnabled" yaml:"notApplicableEnabled"`
	// The properties of the type of question.
	//
	// Text questions do not have to define question type properties.
	QuestionTypeProperties interface{} `field:"optional" json:"questionTypeProperties" yaml:"questionTypeProperties"`
	// The scoring weight of the section.
	//
	// *Minimum* : 0
	//
	// *Maximum* : 100.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

