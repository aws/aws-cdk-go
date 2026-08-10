package awsconnect


// Information about the points configuration for an answer option.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   questionOptionPointsConfigurationProperty := &QuestionOptionPointsConfigurationProperty{
//   	PointValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	IsBonus: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-questionoptionpointsconfiguration.html
//
type CfnEvaluationForm_QuestionOptionPointsConfigurationProperty struct {
	// The point value for scoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-questionoptionpointsconfiguration.html#cfn-connect-evaluationform-questionoptionpointsconfiguration-pointvalue
	//
	PointValue *float64 `field:"required" json:"pointValue" yaml:"pointValue"`
	// The flag to mark the option as a bonus option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-questionoptionpointsconfiguration.html#cfn-connect-evaluationform-questionoptionpointsconfiguration-isbonus
	//
	IsBonus interface{} `field:"optional" json:"isBonus" yaml:"isBonus"`
}

