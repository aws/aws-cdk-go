package mixinsawsconnect


// Information about properties for a question in an evaluation form.
//
// The question type properties must be either for a numeric question or a single select question.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evaluationFormQuestionTypePropertiesProperty := &EvaluationFormQuestionTypePropertiesProperty{
//   	Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   		Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   			AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   				SourceType: jsii.String("sourceType"),
//   			},
//   			PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   				Label: jsii.String("label"),
//   			},
//   		},
//   		MaxValue: jsii.Number(123),
//   		MinValue: jsii.Number(123),
//   		Options: []interface{}{
//   			&EvaluationFormNumericQuestionOptionProperty{
//   				AutomaticFail: jsii.Boolean(false),
//   				AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   					TargetSection: jsii.String("targetSection"),
//   				},
//   				MaxValue: jsii.Number(123),
//   				MinValue: jsii.Number(123),
//   				Score: jsii.Number(123),
//   			},
//   		},
//   	},
//   	SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   		Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   			AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   				SourceType: jsii.String("sourceType"),
//   			},
//   			DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   			Options: []interface{}{
//   				&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   					RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   						Category: jsii.String("category"),
//   						Condition: jsii.String("condition"),
//   						OptionRefId: jsii.String("optionRefId"),
//   					},
//   				},
//   			},
//   		},
//   		DisplayAs: jsii.String("displayAs"),
//   		Options: []interface{}{
//   			&EvaluationFormSingleSelectQuestionOptionProperty{
//   				AutomaticFail: jsii.Boolean(false),
//   				AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   					TargetSection: jsii.String("targetSection"),
//   				},
//   				RefId: jsii.String("refId"),
//   				Score: jsii.Number(123),
//   				Text: jsii.String("text"),
//   			},
//   		},
//   	},
//   	Text: &EvaluationFormTextQuestionPropertiesProperty{
//   		Automation: &EvaluationFormTextQuestionAutomationProperty{
//   			AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   				SourceType: jsii.String("sourceType"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestiontypeproperties.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormQuestionTypePropertiesProperty struct {
	// The properties of the numeric question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestiontypeproperties.html#cfn-connect-evaluationform-evaluationformquestiontypeproperties-numeric
	//
	Numeric interface{} `field:"optional" json:"numeric" yaml:"numeric"`
	// The properties of the numeric question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestiontypeproperties.html#cfn-connect-evaluationform-evaluationformquestiontypeproperties-singleselect
	//
	SingleSelect interface{} `field:"optional" json:"singleSelect" yaml:"singleSelect"`
	// The properties of the text question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestiontypeproperties.html#cfn-connect-evaluationform-evaluationformquestiontypeproperties-text
	//
	Text interface{} `field:"optional" json:"text" yaml:"text"`
}

