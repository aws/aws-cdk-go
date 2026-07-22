package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerParameterRangeSpecificationProperty := &IntegerParameterRangeSpecificationProperty{
//   	MaxValue: jsii.String("maxValue"),
//   	MinValue: jsii.String("minValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-integerparameterrangespecification.html
//
type CfnAlgorithm_IntegerParameterRangeSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-integerparameterrangespecification.html#cfn-sagemaker-algorithm-integerparameterrangespecification-maxvalue
	//
	MaxValue *string `field:"required" json:"maxValue" yaml:"maxValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-integerparameterrangespecification.html#cfn-sagemaker-algorithm-integerparameterrangespecification-minvalue
	//
	MinValue *string `field:"required" json:"minValue" yaml:"minValue"`
}

