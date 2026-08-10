package awsconnect


// An option for a multi-select question in an evaluation form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   evaluationFormMultiSelectQuestionOptionProperty := &EvaluationFormMultiSelectQuestionOptionProperty{
//   	AutomaticFail: jsii.Boolean(false),
//   	AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   		TargetSection: jsii.String("targetSection"),
//   	},
//   	PointsConfiguration: &QuestionOptionPointsConfigurationProperty{
//   		IsBonus: jsii.Boolean(false),
//   		PointValue: jsii.Number(123),
//   	},
//   	RefId: jsii.String("refId"),
//   	Score: jsii.Number(123),
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormMultiSelectQuestionOptionProperty struct {
	// The flag to mark the option as automatic fail.
	//
	// If an automatic fail answer is provided, the overall evaluation gets a score of 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionoption-automaticfail
	//
	AutomaticFail interface{} `field:"optional" json:"automaticFail" yaml:"automaticFail"`
	// Information about automatic fail configuration for an evaluation form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionoption-automaticfailconfiguration
	//
	AutomaticFailConfiguration interface{} `field:"optional" json:"automaticFailConfiguration" yaml:"automaticFailConfiguration"`
	// Information about the points configuration for an answer option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionoption-pointsconfiguration
	//
	PointsConfiguration interface{} `field:"optional" json:"pointsConfiguration" yaml:"pointsConfiguration"`
	// Reference identifier for this option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionoption-refid
	//
	RefId *string `field:"optional" json:"refId" yaml:"refId"`
	// The score of an answer option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionoption-score
	//
	Score *float64 `field:"optional" json:"score" yaml:"score"`
	// Display text for this option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionoption-text
	//
	Text *string `field:"optional" json:"text" yaml:"text"`
}

