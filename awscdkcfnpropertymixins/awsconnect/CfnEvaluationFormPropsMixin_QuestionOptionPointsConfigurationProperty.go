package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   questionOptionPointsConfigurationProperty := &QuestionOptionPointsConfigurationProperty{
//   	IsBonus: jsii.Boolean(false),
//   	PointValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-questionoptionpointsconfiguration.html
//
type CfnEvaluationFormPropsMixin_QuestionOptionPointsConfigurationProperty struct {
	// Whether this option is a bonus.
	//
	// Note: Bonus options are not supported for multi-select questions. This property should only be set to true for single-select and numeric question options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-questionoptionpointsconfiguration.html#cfn-connect-evaluationform-questionoptionpointsconfiguration-isbonus
	//
	IsBonus interface{} `field:"optional" json:"isBonus" yaml:"isBonus"`
	// The point value for scoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-questionoptionpointsconfiguration.html#cfn-connect-evaluationform-questionoptionpointsconfiguration-pointvalue
	//
	PointValue *float64 `field:"optional" json:"pointValue" yaml:"pointValue"`
}

