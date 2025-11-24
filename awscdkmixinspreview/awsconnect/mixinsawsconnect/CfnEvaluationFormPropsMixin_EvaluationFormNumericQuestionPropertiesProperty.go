package mixinsawsconnect


// Information about properties for a numeric question in an evaluation form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evaluationFormNumericQuestionPropertiesProperty := &EvaluationFormNumericQuestionPropertiesProperty{
//   	Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   		AnswerSource: &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   			SourceType: jsii.String("sourceType"),
//   		},
//   		PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   			Label: jsii.String("label"),
//   		},
//   	},
//   	MaxValue: jsii.Number(123),
//   	MinValue: jsii.Number(123),
//   	Options: []interface{}{
//   		&EvaluationFormNumericQuestionOptionProperty{
//   			AutomaticFail: jsii.Boolean(false),
//   			AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   				TargetSection: jsii.String("targetSection"),
//   			},
//   			MaxValue: jsii.Number(123),
//   			MinValue: jsii.Number(123),
//   			Score: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformnumericquestionproperties.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormNumericQuestionPropertiesProperty struct {
	// The automation properties of the numeric question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformnumericquestionproperties.html#cfn-connect-evaluationform-evaluationformnumericquestionproperties-automation
	//
	Automation interface{} `field:"optional" json:"automation" yaml:"automation"`
	// The maximum answer value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformnumericquestionproperties.html#cfn-connect-evaluationform-evaluationformnumericquestionproperties-maxvalue
	//
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// The minimum answer value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformnumericquestionproperties.html#cfn-connect-evaluationform-evaluationformnumericquestionproperties-minvalue
	//
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
	// The scoring options of the numeric question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformnumericquestionproperties.html#cfn-connect-evaluationform-evaluationformnumericquestionproperties-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
}

