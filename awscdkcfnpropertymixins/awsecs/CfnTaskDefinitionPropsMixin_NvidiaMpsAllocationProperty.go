package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   nvidiaMpsAllocationProperty := &NvidiaMpsAllocationProperty{
//   	MaxComputePercent: jsii.Number(123),
//   	Memory: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-nvidiampsallocation.html
//
type CfnTaskDefinitionPropsMixin_NvidiaMpsAllocationProperty struct {
	// The maximum percentage of the GPU's compute capacity that the container can use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-nvidiampsallocation.html#cfn-ecs-taskdefinition-nvidiampsallocation-maxcomputepercent
	//
	MaxComputePercent *float64 `field:"optional" json:"maxComputePercent" yaml:"maxComputePercent"`
	// The amount of GPU memory, in MiB, allocated to the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-nvidiampsallocation.html#cfn-ecs-taskdefinition-nvidiampsallocation-memory
	//
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
}

