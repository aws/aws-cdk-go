package awsautoscaling


// `MemoryGiBPerVCpuRequest` is a property of the `InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property type that describes the minimum and maximum amount of memory per vCPU for an instance type, in GiB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memoryGiBPerVCpuRequestProperty := &memoryGiBPerVCpuRequestProperty{
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   }
//
type CfnAutoScalingGroup_MemoryGiBPerVCpuRequestProperty struct {
	// The memory maximum in GiB.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The memory minimum in GiB.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

