package awspersonalize


// Provides the name and range of a continuous hyperparameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   continuousHyperParameterRangeProperty := &ContinuousHyperParameterRangeProperty{
//   	MaxValue: jsii.Number(123),
//   	MinValue: jsii.Number(123),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-continuoushyperparameterrange.html
//
type CfnSolution_ContinuousHyperParameterRangeProperty struct {
	// The maximum allowable value for the hyperparameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-continuoushyperparameterrange.html#cfn-personalize-solution-continuoushyperparameterrange-maxvalue
	//
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// The minimum allowable value for the hyperparameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-continuoushyperparameterrange.html#cfn-personalize-solution-continuoushyperparameterrange-minvalue
	//
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
	// The name of the hyperparameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-continuoushyperparameterrange.html#cfn-personalize-solution-continuoushyperparameterrange-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

