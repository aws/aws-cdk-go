package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   integerParameterRangeSpecificationProperty := &IntegerParameterRangeSpecificationProperty{
//   	MaxValue: jsii.String("maxValue"),
//   	MinValue: jsii.String("minValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-integerparameterrangespecification.html
//
type CfnAlgorithmPropsMixin_IntegerParameterRangeSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-integerparameterrangespecification.html#cfn-sagemaker-algorithm-integerparameterrangespecification-maxvalue
	//
	MaxValue *string `field:"optional" json:"maxValue" yaml:"maxValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-integerparameterrangespecification.html#cfn-sagemaker-algorithm-integerparameterrangespecification-minvalue
	//
	MinValue *string `field:"optional" json:"minValue" yaml:"minValue"`
}

