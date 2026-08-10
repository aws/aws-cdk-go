package awsconnect


// Scoring configuration for a question in an evaluation form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormQuestionScoringConfigurationProperty := &EvaluationFormQuestionScoringConfigurationProperty{
//   	IsExcludedFromScoring: jsii.Boolean(false),
//   	PointsConfiguration: &QuestionPointsConfigurationProperty{
//   		IsBonus: jsii.Boolean(false),
//   		MaxPointValue: jsii.Number(123),
//   		MinPointValue: jsii.Number(123),
//   	},
//   	ScoreThresholds: []interface{}{
//   		&EvaluationFormScoreThresholdProperty{
//   			PerformanceCategory: jsii.String("performanceCategory"),
//
//   			// the properties below are optional
//   			MaxScorePercentage: jsii.Number(123),
//   			MinScorePercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestionscoringconfiguration.html
//
type CfnEvaluationForm_EvaluationFormQuestionScoringConfigurationProperty struct {
	// The flag to exclude the question from scoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestionscoringconfiguration.html#cfn-connect-evaluationform-evaluationformquestionscoringconfiguration-isexcludedfromscoring
	//
	IsExcludedFromScoring interface{} `field:"optional" json:"isExcludedFromScoring" yaml:"isExcludedFromScoring"`
	// Information about the points configuration for a question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestionscoringconfiguration.html#cfn-connect-evaluationform-evaluationformquestionscoringconfiguration-pointsconfiguration
	//
	PointsConfiguration interface{} `field:"optional" json:"pointsConfiguration" yaml:"pointsConfiguration"`
	// The score thresholds for performance categories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestionscoringconfiguration.html#cfn-connect-evaluationform-evaluationformquestionscoringconfiguration-scorethresholds
	//
	ScoreThresholds interface{} `field:"optional" json:"scoreThresholds" yaml:"scoreThresholds"`
}

