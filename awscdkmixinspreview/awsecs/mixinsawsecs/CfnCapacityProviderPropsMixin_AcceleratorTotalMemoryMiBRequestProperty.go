package mixinsawsecs


// The minimum and maximum total accelerator memory in mebibytes (MiB) for instance type selection.
//
// This is important for GPU workloads that require specific amounts of video memory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   acceleratorTotalMemoryMiBRequestProperty := &AcceleratorTotalMemoryMiBRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-acceleratortotalmemorymibrequest.html
//
type CfnCapacityProviderPropsMixin_AcceleratorTotalMemoryMiBRequestProperty struct {
	// The maximum total accelerator memory in MiB.
	//
	// Instance types with more accelerator memory are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-acceleratortotalmemorymibrequest.html#cfn-ecs-capacityprovider-acceleratortotalmemorymibrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum total accelerator memory in MiB.
	//
	// Instance types with less accelerator memory are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-acceleratortotalmemorymibrequest.html#cfn-ecs-capacityprovider-acceleratortotalmemorymibrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

