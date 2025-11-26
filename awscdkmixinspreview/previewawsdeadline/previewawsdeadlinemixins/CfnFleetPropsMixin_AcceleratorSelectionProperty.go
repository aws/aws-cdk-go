package previewawsdeadlinemixins


// Describes a specific GPU accelerator required for an Amazon Elastic Compute Cloud worker host.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   acceleratorSelectionProperty := &AcceleratorSelectionProperty{
//   	Name: jsii.String("name"),
//   	Runtime: jsii.String("runtime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorselection.html
//
type CfnFleetPropsMixin_AcceleratorSelectionProperty struct {
	// The name of the chip used by the GPU accelerator.
	//
	// If you specify `l4` as the name of the accelerator, you must specify `latest` or `grid:r570` as the runtime.
	//
	// The available GPU accelerators are:
	//
	// - `t4` - NVIDIA T4 Tensor Core GPU
	// - `a10g` - NVIDIA A10G Tensor Core GPU
	// - `l4` - NVIDIA L4 Tensor Core GPU
	// - `l40s` - NVIDIA L40S Tensor Core GPU.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorselection.html#cfn-deadline-fleet-acceleratorselection-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the runtime driver to use for the GPU accelerator. You must use the same runtime for all GPUs.
	//
	// You can choose from the following runtimes:
	//
	// - `latest` - Use the latest runtime available for the chip. If you specify `latest` and a new version of the runtime is released, the new version of the runtime is used.
	// - `grid:r570` - [NVIDIA vGPU software 18](https://docs.aws.amazon.com/https://docs.nvidia.com/vgpu/18.0/index.html)
	// - `grid:r535` - [NVIDIA vGPU software 16](https://docs.aws.amazon.com/https://docs.nvidia.com/vgpu/16.0/index.html)
	//
	// If you don't specify a runtime, Deadline Cloud uses `latest` as the default. However, if you have multiple accelerators and specify `latest` for some and leave others blank, Deadline Cloud raises an exception.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorselection.html#cfn-deadline-fleet-acceleratorselection-runtime
	//
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
}

