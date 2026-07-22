package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   questionPointsConfigurationProperty := &QuestionPointsConfigurationProperty{
//   	IsBonus: jsii.Boolean(false),
//   	MaxPointValue: jsii.Number(123),
//   	MinPointValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-questionpointsconfiguration.html
//
type CfnEvaluationFormPropsMixin_QuestionPointsConfigurationProperty struct {
	// Whether the question is a bonus question.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-questionpointsconfiguration.html#cfn-connect-evaluationform-questionpointsconfiguration-isbonus
	//
	IsBonus interface{} `field:"optional" json:"isBonus" yaml:"isBonus"`
	// The point value for scoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-questionpointsconfiguration.html#cfn-connect-evaluationform-questionpointsconfiguration-maxpointvalue
	//
	MaxPointValue *float64 `field:"optional" json:"maxPointValue" yaml:"maxPointValue"`
	// The point value for scoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-questionpointsconfiguration.html#cfn-connect-evaluationform-questionpointsconfiguration-minpointvalue
	//
	MinPointValue *float64 `field:"optional" json:"minPointValue" yaml:"minPointValue"`
}

