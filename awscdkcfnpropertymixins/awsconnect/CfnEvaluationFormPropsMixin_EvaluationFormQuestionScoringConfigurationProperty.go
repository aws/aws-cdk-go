package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
//   			MaxScorePercentage: jsii.Number(123),
//   			MinScorePercentage: jsii.Number(123),
//   			PerformanceCategory: jsii.String("performanceCategory"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestionscoringconfiguration.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormQuestionScoringConfigurationProperty struct {
	// Whether the question is excluded from scoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestionscoringconfiguration.html#cfn-connect-evaluationform-evaluationformquestionscoringconfiguration-isexcludedfromscoring
	//
	IsExcludedFromScoring interface{} `field:"optional" json:"isExcludedFromScoring" yaml:"isExcludedFromScoring"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestionscoringconfiguration.html#cfn-connect-evaluationform-evaluationformquestionscoringconfiguration-pointsconfiguration
	//
	PointsConfiguration interface{} `field:"optional" json:"pointsConfiguration" yaml:"pointsConfiguration"`
	// The score thresholds for the question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestionscoringconfiguration.html#cfn-connect-evaluationform-evaluationformquestionscoringconfiguration-scorethresholds
	//
	ScoreThresholds interface{} `field:"optional" json:"scoreThresholds" yaml:"scoreThresholds"`
}

