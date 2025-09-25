package awspersonalize


// Provides the name and range of an integer-valued hyperparameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerHyperParameterRangeProperty := &IntegerHyperParameterRangeProperty{
//   	MaxValue: jsii.Number(123),
//   	MinValue: jsii.Number(123),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-integerhyperparameterrange.html
//
type CfnSolution_IntegerHyperParameterRangeProperty struct {
	// The maximum allowable value for the hyperparameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-integerhyperparameterrange.html#cfn-personalize-solution-integerhyperparameterrange-maxvalue
	//
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// The minimum allowable value for the hyperparameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-integerhyperparameterrange.html#cfn-personalize-solution-integerhyperparameterrange-minvalue
	//
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
	// The name of the hyperparameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-integerhyperparameterrange.html#cfn-personalize-solution-integerhyperparameterrange-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

