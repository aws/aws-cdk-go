package awsautoscaling


// `VCpuCountRequest` is a property of the `InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property type that describes the minimum and maximum number of vCPUs for an instance type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vCpuCountRequestProperty := &vCpuCountRequestProperty{
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   }
//
type CfnAutoScalingGroup_VCpuCountRequestProperty struct {
	// The maximum number of vCPUs.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum number of vCPUs.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

