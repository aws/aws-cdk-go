package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterRangeProperty := &ParameterRangeProperty{
//   	CategoricalParameterRangeSpecification: &CategoricalParameterRangeSpecificationProperty{
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	ContinuousParameterRangeSpecification: &ContinuousParameterRangeSpecificationProperty{
//   		MaxValue: jsii.String("maxValue"),
//   		MinValue: jsii.String("minValue"),
//   	},
//   	IntegerParameterRangeSpecification: &IntegerParameterRangeSpecificationProperty{
//   		MaxValue: jsii.String("maxValue"),
//   		MinValue: jsii.String("minValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-parameterrange.html
//
type CfnAlgorithm_ParameterRangeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-parameterrange.html#cfn-sagemaker-algorithm-parameterrange-categoricalparameterrangespecification
	//
	CategoricalParameterRangeSpecification interface{} `field:"optional" json:"categoricalParameterRangeSpecification" yaml:"categoricalParameterRangeSpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-parameterrange.html#cfn-sagemaker-algorithm-parameterrange-continuousparameterrangespecification
	//
	ContinuousParameterRangeSpecification interface{} `field:"optional" json:"continuousParameterRangeSpecification" yaml:"continuousParameterRangeSpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-parameterrange.html#cfn-sagemaker-algorithm-parameterrange-integerparameterrangespecification
	//
	IntegerParameterRangeSpecification interface{} `field:"optional" json:"integerParameterRangeSpecification" yaml:"integerParameterRangeSpecification"`
}

