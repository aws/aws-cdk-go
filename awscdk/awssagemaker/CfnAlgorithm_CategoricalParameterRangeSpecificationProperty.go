package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   categoricalParameterRangeSpecificationProperty := &CategoricalParameterRangeSpecificationProperty{
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-categoricalparameterrangespecification.html
//
type CfnAlgorithm_CategoricalParameterRangeSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-categoricalparameterrangespecification.html#cfn-sagemaker-algorithm-categoricalparameterrangespecification-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

