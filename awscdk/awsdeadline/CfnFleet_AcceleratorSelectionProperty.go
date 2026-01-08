package awsdeadline


// Describes a specific GPU accelerator required for an Amazon Elastic Compute Cloud worker host.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorSelectionProperty := &AcceleratorSelectionProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Runtime: jsii.String("runtime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorselection.html
//
type CfnFleet_AcceleratorSelectionProperty struct {
	// The name of the chip used by the GPU accelerator.
	//
	// The available GPU accelerators are:
	//
	// - `t4` - NVIDIA T4 Tensor Core GPU (16 GiB memory)
	// - `a10g` - NVIDIA A10G Tensor Core GPU (24 GiB memory)
	// - `l4` - NVIDIA L4 Tensor Core GPU (24 GiB memory)
	// - `l40s` - NVIDIA L40S Tensor Core GPU (48 GiB memory).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorselection.html#cfn-deadline-fleet-acceleratorselection-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the runtime driver to use for the GPU accelerator.
	//
	// You must use the same runtime for all GPUs in a fleet.
	//
	// You can choose from the following runtimes:
	//
	// - `latest` - Use the latest runtime available for the chip. If you specify `latest` and a new version of the runtime is released, the new version of the runtime is used.
	// - `grid:r570` - [NVIDIA vGPU software 18](https://docs.aws.amazon.com/https://docs.nvidia.com/vgpu/18.0/index.html)
	// - `grid:r535` - [NVIDIA vGPU software 16](https://docs.aws.amazon.com/https://docs.nvidia.com/vgpu/16.0/index.html)
	//
	// If you don't specify a runtime, AWS Deadline Cloud uses `latest` as the default. However, if you have multiple accelerators and specify `latest` for some and leave others blank, AWS Deadline Cloud raises an exception.
	//
	// > Not all runtimes are compatible with all accelerator types:
	// >
	// > - `t4` and `a10g` : Support all runtimes ( `grid:r570` , `grid:r535` )
	// > - `l4` and `l40s` : Only support `grid:r570` and newer
	// >
	// > All accelerators in a fleet must use the same runtime version. You cannot mix different runtime versions within a single fleet. > When you specify `latest` , it resolves to `grid:r570` for all currently supported accelerators.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorselection.html#cfn-deadline-fleet-acceleratorselection-runtime
	//
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
}

