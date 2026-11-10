package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nvidiaMpsAllocationProperty := &NvidiaMpsAllocationProperty{
//   	Memory: jsii.Number(123),
//
//   	// the properties below are optional
//   	MaxComputePercent: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-nvidiampsallocation.html
//
type CfnTaskDefinition_NvidiaMpsAllocationProperty struct {
	// The amount of GPU memory, in MiB, allocated to the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-nvidiampsallocation.html#cfn-ecs-taskdefinition-nvidiampsallocation-memory
	//
	Memory *float64 `field:"required" json:"memory" yaml:"memory"`
	// The maximum percentage of the GPU's compute capacity that the container can use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-nvidiampsallocation.html#cfn-ecs-taskdefinition-nvidiampsallocation-maxcomputepercent
	//
	MaxComputePercent *float64 `field:"optional" json:"maxComputePercent" yaml:"maxComputePercent"`
}

