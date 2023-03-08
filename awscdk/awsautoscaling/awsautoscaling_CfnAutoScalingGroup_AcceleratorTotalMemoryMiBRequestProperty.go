package awsautoscaling


// `AcceleratorTotalMemoryMiBRequest` is a property of the `InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property type that describes the minimum and maximum total memory size for the accelerators for an instance type, in MiB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorTotalMemoryMiBRequestProperty := &acceleratorTotalMemoryMiBRequestProperty{
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   }
//
type CfnAutoScalingGroup_AcceleratorTotalMemoryMiBRequestProperty struct {
	// The memory maximum in MiB.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The memory minimum in MiB.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

