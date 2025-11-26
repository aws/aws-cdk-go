package previewawsecsmixins


// The minimum and maximum amount of memory per vCPU in gibibytes (GiB).
//
// This helps ensure that instance types have the appropriate memory-to-CPU ratio for your workloads.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   memoryGiBPerVCpuRequestProperty := &MemoryGiBPerVCpuRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-memorygibpervcpurequest.html
//
type CfnCapacityProviderPropsMixin_MemoryGiBPerVCpuRequestProperty struct {
	// The maximum amount of memory per vCPU in GiB.
	//
	// Instance types with a higher memory-to-vCPU ratio are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-memorygibpervcpurequest.html#cfn-ecs-capacityprovider-memorygibpervcpurequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum amount of memory per vCPU in GiB.
	//
	// Instance types with a lower memory-to-vCPU ratio are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-memorygibpervcpurequest.html#cfn-ecs-capacityprovider-memorygibpervcpurequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

