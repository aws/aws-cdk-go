package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   continuousParameterRangeSpecificationProperty := &ContinuousParameterRangeSpecificationProperty{
//   	MaxValue: jsii.String("maxValue"),
//   	MinValue: jsii.String("minValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-continuousparameterrangespecification.html
//
type CfnAlgorithm_ContinuousParameterRangeSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-continuousparameterrangespecification.html#cfn-sagemaker-algorithm-continuousparameterrangespecification-maxvalue
	//
	MaxValue *string `field:"required" json:"maxValue" yaml:"maxValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-continuousparameterrangespecification.html#cfn-sagemaker-algorithm-continuousparameterrangespecification-minvalue
	//
	MinValue *string `field:"required" json:"minValue" yaml:"minValue"`
}

