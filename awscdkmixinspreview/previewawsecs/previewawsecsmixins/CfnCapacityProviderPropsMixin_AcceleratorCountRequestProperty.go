package previewawsecsmixins


// The minimum and maximum number of accelerators (such as GPUs) for instance type selection.
//
// This is used for workloads that require specific numbers of accelerators.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   acceleratorCountRequestProperty := &AcceleratorCountRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-acceleratorcountrequest.html
//
type CfnCapacityProviderPropsMixin_AcceleratorCountRequestProperty struct {
	// The maximum number of accelerators.
	//
	// Instance types with more accelerators are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-acceleratorcountrequest.html#cfn-ecs-capacityprovider-acceleratorcountrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum number of accelerators.
	//
	// Instance types with fewer accelerators are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-acceleratorcountrequest.html#cfn-ecs-capacityprovider-acceleratorcountrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

