package awsconnect


// An option for a multi-select question in an evaluation form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormMultiSelectQuestionOptionProperty := &EvaluationFormMultiSelectQuestionOptionProperty{
//   	RefId: jsii.String("refId"),
//   	Text: jsii.String("text"),
//
//   	// the properties below are optional
//   	AutomaticFail: jsii.Boolean(false),
//   	AutomaticFailConfiguration: &AutomaticFailConfigurationProperty{
//   		TargetSection: jsii.String("targetSection"),
//   	},
//   	PointsConfiguration: &QuestionOptionPointsConfigurationProperty{
//   		PointValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		IsBonus: jsii.Boolean(false),
//   	},
//   	Score: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html
//
type CfnEvaluationForm_EvaluationFormMultiSelectQuestionOptionProperty struct {
	// Reference identifier for this option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionoption-refid
	//
	RefId *string `field:"required" json:"refId" yaml:"refId"`
	// Display text for this option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionoption-text
	//
	Text *string `field:"required" json:"text" yaml:"text"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionoption-automaticfail
	//
	AutomaticFail interface{} `field:"optional" json:"automaticFail" yaml:"automaticFail"`
	// Information about automatic fail configuration for an evaluation form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionoption-automaticfailconfiguration
	//
	AutomaticFailConfiguration interface{} `field:"optional" json:"automaticFailConfiguration" yaml:"automaticFailConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionoption-pointsconfiguration
	//
	PointsConfiguration interface{} `field:"optional" json:"pointsConfiguration" yaml:"pointsConfiguration"`
	// The score of an answer option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionoption-score
	//
	Score *float64 `field:"optional" json:"score" yaml:"score"`
}

