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
//   	MultiSelect: &EvaluationFormMultiSelectQuestionPropertiesProperty{
//   		Options: []interface{}{
//   			&EvaluationFormMultiSelectQuestionOptionProperty{
//   				RefId: jsii.String("refId"),
//   				Text: jsii.String("text"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Automation: &EvaluationFormMultiSelectQuestionAutomationProperty{
//   			Options: []interface{}{
//   				&EvaluationFormMultiSelectQuestionAutomationOptionProperty{
//   					RuleCategory: &MultiSelectQuestionRuleCategoryAutomationProperty{
//   						Category: jsii.String("category"),
//   						Condition: jsii.String("condition"),
//   						OptionRefIds: []*string{
//   							jsii.String("optionRefIds"),
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   				SourceType: jsii.String("sourceType"),
//   			},
//   			DefaultOptionRefIds: []*string{
//   				jsii.String("defaultOptionRefIds"),
//   			},
//   		},
//   		DisplayAs: jsii.String("displayAs"),
//   	},
//   	Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   		MaxValue: jsii.Number(123),
//   		MinValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   			AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   				SourceType: jsii.String("sourceType"),
//   			},
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
//   				AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   					TargetSection: jsii.String("targetSection"),
//   				},
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
//   				AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   					TargetSection: jsii.String("targetSection"),
//   				},
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
//   			AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   				SourceType: jsii.String("sourceType"),
//   			},
//   			DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   		},
//   		DisplayAs: jsii.String("displayAs"),
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
type CfnEvaluationForm_EvaluationFormQuestionTypePropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestiontypeproperties.html#cfn-connect-evaluationform-evaluationformquestiontypeproperties-multiselect
	//
	MultiSelect interface{} `field:"optional" json:"multiSelect" yaml:"multiSelect"`
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

